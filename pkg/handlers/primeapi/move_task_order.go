package primeapi

import (
	"github.com/transcom/mymove/pkg/services"
	mtoshipmentservice "github.com/transcom/mymove/pkg/services/mto_shipment"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/gofrs/uuid"
	"github.com/transcom/mymove/pkg/handlers/primeapi/internal/payloads"
	"github.com/transcom/mymove/pkg/models"

	movetaskorderops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/move_task_order"
	"github.com/transcom/mymove/pkg/handlers"
)

// FetchMTOUpdatesHandler lists move task orders with the option to filter since a particular date
type FetchMTOUpdatesHandler struct {
	handlers.HandlerContext
}

// Handle fetches all move task orders with the option to filter since a particular date
func (h FetchMTOUpdatesHandler) Handle(params movetaskorderops.FetchMTOUpdatesParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)

	var mtos models.MoveTaskOrders

	query := h.DB().Where("is_available_to_prime = ?", true).Eager("PaymentRequests", "MTOServiceItems", "MTOServiceItems.ReService")
	if params.Since != nil {
		since := time.Unix(*params.Since, 0)
		query = query.Where("updated_at > ?", since)
	}

	err := query.All(&mtos)

	if err != nil {
		logger.Error("Unable to fetch records:", zap.Error(err))
		return movetaskorderops.NewFetchMTOUpdatesInternalServerError()
	}

	payload := payloads.MoveTaskOrders(&mtos)

	return movetaskorderops.NewFetchMTOUpdatesOK().WithPayload(payload)
}

type UpdateMTOPostCounselingInfoHandler struct {
	handlers.HandlerContext
	mtoUpdater services.MoveTaskOrderUpdater
}

// Handle handler that updates a mto shipment
func (h UpdateMTOPostCounselingInfoHandler) Handle(params movetaskorderops.UpdateMTOPostCounselingInformationParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)
	mto, err := h.mtoUpdater.UpdatePostCounselingInfo(uuid.FromStringOrNil(params.MoveTaskOrderID), params)
	if err != nil {
		logger.Error("primeapi.UpdateMTOPostCounselingInfoHandler error", zap.Error(err))
		switch err.(type) {
		case mtoshipmentservice.ErrNotFound:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationNotFound()
		case mtoshipmentservice.ErrInvalidInput:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationUnprocessableEntity()
		case mtoshipmentservice.ErrPreconditionFailed:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationPreconditionFailed()
		default:
			return movetaskorderops.NewUpdateMTOPostCounselingInformationInternalServerError()
		}
	}
	mtoPayload := payloads.MoveTaskOrder(mto)
	return movetaskorderops.NewUpdateMTOPostCounselingInformationOK().WithPayload(mtoPayload)
}

