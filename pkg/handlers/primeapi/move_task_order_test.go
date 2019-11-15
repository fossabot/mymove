package primeapi

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"

	movetaskorderops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/move_task_order"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/mocks"
	movetaskorder "github.com/transcom/mymove/pkg/services/move_task_order"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestListMoveTaskOrdersHandler() {
	moveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})

	request := httptest.NewRequest("GET", "/move-task-orders", nil)

	params := movetaskorderops.ListMoveTaskOrdersParams{HTTPRequest: request}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

	// make the request
	handler := ListMoveTaskOrdersHandler{HandlerContext: context}
	response := handler.Handle(params)

	suite.IsNotErrResponse(response)
	moveTaskOrdersResponse := response.(*movetaskorderops.ListMoveTaskOrdersOK)
	moveTaskOrdersPayload := moveTaskOrdersResponse.Payload

	suite.Equal(1, len(moveTaskOrdersPayload))
	suite.Equal(moveTaskOrder.ID.String(), moveTaskOrdersPayload[0].ID.String())
}

func (suite *HandlerSuite) TestListMoveTaskOrdersHandlerReturnsUpdated() {
	now := time.Now()
	lastFetch := now.Add(-time.Second)

	moveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})

	// this MTO should not be returned
	olderMoveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})

	// Pop will overwrite UpdatedAt when saving a model, so use SQL to set it in the past
	suite.NoError(suite.DB().RawQuery("UPDATE move_task_orders SET updated_at=? WHERE id=?",
		now.Add(-2*time.Second), olderMoveTaskOrder.ID).Exec())

	since := lastFetch.Unix()
	request := httptest.NewRequest("GET", fmt.Sprintf("/move-task-orders?since=%d", lastFetch.Unix()), nil)

	params := movetaskorderops.ListMoveTaskOrdersParams{HTTPRequest: request, Since: &since}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

	// make the request
	handler := ListMoveTaskOrdersHandler{HandlerContext: context}
	response := handler.Handle(params)

	suite.IsNotErrResponse(response)
	moveTaskOrdersResponse := response.(*movetaskorderops.ListMoveTaskOrdersOK)
	moveTaskOrdersPayload := moveTaskOrdersResponse.Payload

	suite.Equal(1, len(moveTaskOrdersPayload))
	suite.Equal(moveTaskOrder.ID.String(), moveTaskOrdersPayload[0].ID.String())
}

func (suite *HandlerSuite) TestUpdateMoveTaskOrderEstimatedWeightHandlerSuccess() {
	serviceItem := testdatagen.MakeServiceItem(suite.DB(), testdatagen.Assertions{})
	moveTaskOrder := serviceItem.MoveTaskOrder

	// set up what needs to be passed to handler
	request := httptest.NewRequest("PATCH", fmt.Sprintf("/move-task-orders/%s/prime-estimated-weight", moveTaskOrder.ID), nil)
	params := movetaskorderops.UpdateMoveTaskOrderEstimatedWeightParams{
		HTTPRequest: request,
		Body: movetaskorderops.UpdateMoveTaskOrderEstimatedWeightBody{
			PrimeEstimatedWeight: 1220,
		},
		MoveTaskOrderID: moveTaskOrder.ID.String(),
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

	// make the request
	handler := UpdateMoveTaskOrderEstimatedWeightHandler{context,
		movetaskorder.NewMoveTaskOrderEstimatedWeightUpdater(suite.DB())}
	response := handler.Handle(params)

	suite.IsNotErrResponse(response)
	updateMoveTaskOrderEstimatedWeightResponse := response.(*movetaskorderops.UpdateMoveTaskOrderEstimatedWeightOK)
	updateMoveTaskOrderEstimatedWeightPayload := updateMoveTaskOrderEstimatedWeightResponse.Payload

	suite.NotNil(updateMoveTaskOrderEstimatedWeightPayload)
	suite.NotNil(updateMoveTaskOrderEstimatedWeightPayload.PrimeEstimatedWeight)
	suite.Equal(params.Body.PrimeEstimatedWeight, *updateMoveTaskOrderEstimatedWeightPayload.PrimeEstimatedWeight)
	now := strfmt.Date(time.Now())
	suite.NotNil(updateMoveTaskOrderEstimatedWeightPayload.PrimeEstimatedWeightRecordedDate)
	suite.Equal(now.String(), updateMoveTaskOrderEstimatedWeightPayload.PrimeEstimatedWeightRecordedDate.String())
}

func (suite *HandlerSuite) TestUpdateMoveTaskOrderEstimatedWeightHandlerUnprocessableEntity() {
	serviceItem := testdatagen.MakeServiceItem(suite.DB(), testdatagen.Assertions{})
	moveTaskOrder := serviceItem.MoveTaskOrder

	// set up what needs to be passed to handler
	request := httptest.NewRequest("PATCH", fmt.Sprintf("/move-task-orders/%s/prime-estimated-weight", moveTaskOrder.ID), nil)
	params := movetaskorderops.UpdateMoveTaskOrderEstimatedWeightParams{
		HTTPRequest: request,
	}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

	// make the request
	updater := &mocks.MoveTaskOrderPrimeEstimatedWeightUpdater{}

	id, _ := uuid.NewV4()
	verrs := make(map[string][]string)
	verrs["validation xyz"] = []string{"validation failure"}
	moveTaskOrderValidationError := movetaskorder.NewErrInvalidInput(id, errors.New("error"), verrs)
	updater.On("UpdatePrimeEstimatedWeight",
		mock.AnythingOfType("uuid.UUID"),
		mock.AnythingOfType("unit.Pound"),
		mock.AnythingOfType("time.Time"),
	).Return(&models.MoveTaskOrder{}, moveTaskOrderValidationError)
	handler := UpdateMoveTaskOrderEstimatedWeightHandler{
		context,
		updater,
	}
	response := handler.Handle(params)

	suite.Assertions.IsType(&movetaskorderops.UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity{}, response)
	clientErr := response.(*movetaskorderops.UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity).Payload
	suite.Equal(clientErr.InvalidFields, map[string]string{"validation xyz": "validation failure"})
	suite.Equal(*clientErr.Title, handlers.ValidationErrMessage)
	suite.Equal(*clientErr.Detail, moveTaskOrderValidationError.Error())
}

func (suite *HandlerSuite) TestGetMoveTaskOrdersCustomerHandler() {
	moveTaskOrder := testdatagen.MakeMoveTaskOrder(suite.DB(), testdatagen.Assertions{})

	request := httptest.NewRequest("GET", "/move-task-orders/:id/customer", nil)

	params := movetaskorderops.GetMoveTaskOrderCustomerParams{HTTPRequest: request, MoveTaskOrderID: moveTaskOrder.ID.String()}
	context := handlers.NewHandlerContext(suite.DB(), suite.TestLogger())

	// make the request
	handler := GetMoveTaskOrderCustomerHandler{HandlerContext: context,
		moveTaskOrderFetcher: movetaskorder.NewMoveTaskOrderFetcher(suite.DB()),
	}
	response := handler.Handle(params)

	suite.IsNotErrResponse(response)
	customer := response.(*movetaskorderops.GetMoveTaskOrderCustomerOK)
	moveTaskOrdersCustomerPayload := customer.Payload

	suite.Equal(moveTaskOrder.CustomerID.String(), moveTaskOrdersCustomerPayload.ID.String())
}
