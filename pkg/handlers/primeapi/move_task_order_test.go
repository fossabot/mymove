package primeapi

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/mock"
	mtoshipmentops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/mto_shipment"
	"github.com/transcom/mymove/pkg/handlers/primeapi/internal/payloads"
	"github.com/transcom/mymove/pkg/services/mocks"
	mtoshipment "github.com/transcom/mymove/pkg/services/mto_shipment"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/transcom/mymove/pkg/models"

	movetaskorderops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/move_task_order"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestListMoveTaskOrdersHandler() {
	moveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
		MoveTaskOrder: models.MoveTaskOrder{
			IsAvailableToPrime: true,
		},
	})

	testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
		PaymentRequest: models.PaymentRequest{
			MoveTaskOrderID: moveTaskOrder.ID,
		},
	})

	testdatagen.MakeMTOServiceItem(suite.DB(), testdatagen.Assertions{
		MTOServiceItem: models.MTOServiceItem{
			MoveTaskOrderID: moveTaskOrder.ID,
		},
	})

	// unavailable MTO
	testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})

	request := httptest.NewRequest("GET", "/move-task-orders", nil)

	params := movetaskorderops.FetchMTOUpdatesParams{HTTPRequest: request}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

	// make the request
	handler := FetchMTOUpdatesHandler{HandlerContext: context}
	response := handler.Handle(params)

	suite.IsNotErrResponse(response)
	moveTaskOrdersResponse := response.(*movetaskorderops.FetchMTOUpdatesOK)
	moveTaskOrdersPayload := moveTaskOrdersResponse.Payload

	suite.Equal(1, len(moveTaskOrdersPayload))
	suite.Equal(moveTaskOrder.ID.String(), moveTaskOrdersPayload[0].ID.String())
	suite.Equal(1, len(moveTaskOrdersPayload[0].PaymentRequests))
	suite.Equal(1, len(moveTaskOrdersPayload[0].MtoServiceItems))
}

func (suite *HandlerSuite) TestListMoveTaskOrdersHandlerReturnsUpdated() {
	now := time.Now()
	lastFetch := now.Add(-time.Second)

	moveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
		MoveTaskOrder: models.MoveTaskOrder{
			IsAvailableToPrime: true,
		},
	})

	// this MTO should not be returned
	olderMoveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
		MoveTaskOrder: models.MoveTaskOrder{
			IsAvailableToPrime: true,
		},
	})

	// Pop will overwrite UpdatedAt when saving a model, so use SQL to set it in the past
	suite.NoError(suite.DB().RawQuery("UPDATE move_task_orders SET updated_at=? WHERE id=?",
		now.Add(-2*time.Second), olderMoveTaskOrder.ID).Exec())

	since := lastFetch.Unix()
	request := httptest.NewRequest("GET", fmt.Sprintf("/move-task-orders?since=%d", lastFetch.Unix()), nil)

	params := movetaskorderops.FetchMTOUpdatesParams{HTTPRequest: request, Since: &since}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

	// make the request
	handler := FetchMTOUpdatesHandler{HandlerContext: context}
	response := handler.Handle(params)

	suite.IsNotErrResponse(response)
	moveTaskOrdersResponse := response.(*movetaskorderops.FetchMTOUpdatesOK)
	moveTaskOrdersPayload := moveTaskOrdersResponse.Payload

	suite.Equal(1, len(moveTaskOrdersPayload))
	suite.Equal(moveTaskOrder.ID.String(), moveTaskOrdersPayload[0].ID.String())
}

func (suite *HandlerSuite) TestUpdateMTOPostCounselingInformationHandler() {
	ppm := testdatagen.MakeDefaultPPM(suite.DB())
	mto := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{
		MoveTaskOrder: models.MoveTaskOrder{
			PersonallyProcuredMoveID: ppm.ID,
		},
	})

	req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/post-counseling-info", mto.ID.String()), nil)

	params := movetaskorderops.UpdateMTOPostCounselingInformationParams{
		HTTPRequest:       req,
		MoveTaskOrderID:   mto.ID.String(),
		Body: movetaskorderops.UpdateMTOPostCounselingInformationBody{
			PpmEstimatedWeight: 2000,
			PpmIsIncluded:      true,
			PpmType:            "FULL",
		},
		IfUnmodifiedSince: strfmt.DateTime(ppm.UpdatedAt),
	}

	suite.T().Run("Successful PATCH - Integration Test", func(t *testing.T) {
		updater := mto.NewMoveTaskOrderUpdater(suite.DB())
		handler := UpdateMTOPostCounselingInfoHandler{
			handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			updater,
		}

		response := handler.Handle(params)
		suite.IsType(&movetaskorderops.UpdateMTOPostCounselingInformationOK{}, response)

		okResponse := response.(*movetaskorderops.UpdateMTOPostCounselingInformationOK)
		suite.Equal(ppm.ID.String(), okResponse.Payload.ID.String())
	})
}