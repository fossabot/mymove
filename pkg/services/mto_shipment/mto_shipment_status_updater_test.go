package mtoshipment

import (
	"fmt"
	"testing"
	"time"

	mtoserviceitem "github.com/transcom/mymove/pkg/services/mto_service_item"

	"github.com/go-openapi/strfmt"

	mtoshipmentops "github.com/transcom/mymove/pkg/gen/ghcapi/ghcoperations/mto_shipment"
	"github.com/transcom/mymove/pkg/gen/ghcmessages"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/query"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *MTOShipmentServiceSuite) TestUpdateMTOShipmentStatus() {
	mto := testdatagen.MakeDefaultMoveTaskOrder(suite.DB())
	shipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
		MoveTaskOrder: mto,
		MTOShipment: models.MTOShipment{
			ShipmentType: models.MTOShipmentTypeHHGLongHaulDom,
		},
	})
	shipment.Status = models.MTOShipmentStatusSubmitted
	params := mtoshipmentops.PatchMTOShipmentStatusParams{
		ShipmentID:        strfmt.UUID(shipment.ID.String()),
		IfUnmodifiedSince: strfmt.DateTime(shipment.UpdatedAt),
		Body:              &ghcmessages.MTOShipment{Status: "APPROVED"},
	}
	//Need some values for reServices
	reServiceNames := []models.ReServiceName{
		models.DomesticLinehaul,
		models.FuelSurcharge,
		models.DomesticOriginPrice,
		models.DomesticDestinationPrice,
		models.DomesticPacking,
		models.DomesticUnpacking,
	}

	for i, serviceName := range reServiceNames {
		testdatagen.MakeReService(suite.DB(), testdatagen.Assertions{
			ReService: models.ReService{
				Code:      fmt.Sprintf("code%d", i),
				Name:      string(serviceName),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
	}

	builder := query.NewQueryBuilder(suite.DB())
	siCreator := mtoserviceitem.NewMTOServiceItemCreator(builder)
	updater := NewMTOShipmentStatusUpdater(suite.DB(), builder, siCreator)

	suite.T().Run("If we get a mto shipment pointer with a status it should update and return no error", func(t *testing.T) {

		_, err := updater.UpdateMTOShipmentStatus(params)
		serviceItems := models.MTOServiceItems{}
		_ = suite.DB().All(&serviceItems)
		shipments := models.MTOShipment{}
		suite.DB().All(&shipments)
		suite.NoError(err)
	})

	suite.T().Run("Passing in a stale identifier", func(t *testing.T) {
		params := mtoshipmentops.PatchMTOShipmentStatusParams{
			ShipmentID:        strfmt.UUID(shipment.ID.String()),
			IfUnmodifiedSince: strfmt.DateTime(time.Now()), // Stale identifier
			Body:              &ghcmessages.MTOShipment{Status: "APPROVED"},
		}

		_, err := updater.UpdateMTOShipmentStatus(params)
		suite.Error(err)
		suite.IsType(PreconditionFailedError{}, err)
	})

	suite.T().Run("Passing in an invalid status", func(t *testing.T) {
		params := mtoshipmentops.PatchMTOShipmentStatusParams{
			ShipmentID:        strfmt.UUID(shipment.ID.String()),
			IfUnmodifiedSince: strfmt.DateTime(time.Now()), // Stale identifier
			Body:              &ghcmessages.MTOShipment{Status: "invalid"},
		}

		_, err := updater.UpdateMTOShipmentStatus(params)
		suite.Error(err)
		fmt.Printf("%#v", err)
		suite.IsType(ValidationError{}, err)
	})

	suite.T().Run("Passing in a bad shipment id", func(t *testing.T) {
		params := mtoshipmentops.PatchMTOShipmentStatusParams{
			ShipmentID:        strfmt.UUID("424d930b-cf8d-4c10-8059-be8a25ba952a"),
			IfUnmodifiedSince: strfmt.DateTime(time.Now()), // Stale identifier
			Body:              &ghcmessages.MTOShipment{Status: "invalid"},
		}

		_, err := updater.UpdateMTOShipmentStatus(params)
		suite.Error(err)
		fmt.Printf("%#v", err)
		suite.IsType(NotFoundError{}, err)
	})

	suite.T().Run("Changing to APPROVED status records approved_date", func(t *testing.T) {
		shipment2 := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MoveTaskOrder: mto,
		})
		params := mtoshipmentops.PatchMTOShipmentStatusParams{
			ShipmentID:        strfmt.UUID(shipment2.ID.String()),
			IfUnmodifiedSince: strfmt.DateTime(shipment2.UpdatedAt),
			Body:              &ghcmessages.MTOShipment{Status: "APPROVED"},
		}

		suite.Nil(shipment2.ApprovedDate)
		_, err := updater.UpdateMTOShipmentStatus(params)
		suite.NoError(err)
		suite.DB().Find(&shipment2, shipment2.ID)
		suite.Equal(models.MTOShipmentStatusApproved, shipment2.Status)
		suite.NotNil(shipment2.ApprovedDate)
	})

	suite.T().Run("Changing to a non-APPROVED status does not record approved_date", func(t *testing.T) {
		shipment3 := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MoveTaskOrder: mto,
		})
		params := mtoshipmentops.PatchMTOShipmentStatusParams{
			ShipmentID:        strfmt.UUID(shipment3.ID.String()),
			IfUnmodifiedSince: strfmt.DateTime(shipment3.UpdatedAt),
			Body:              &ghcmessages.MTOShipment{Status: "REJECTED"},
		}

		suite.Nil(shipment3.ApprovedDate)
		_, err := updater.UpdateMTOShipmentStatus(params)
		suite.NoError(err)
		suite.DB().Find(&shipment3, shipment3.ID)
		suite.Equal(models.MTOShipmentStatusRejected, shipment3.Status)
		suite.Nil(shipment3.ApprovedDate)
	})
}
