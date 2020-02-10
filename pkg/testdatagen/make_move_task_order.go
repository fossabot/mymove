package testdatagen

import (
	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/models"
)

// MakeMoveTaskOrder creates a single MoveTaskOrder and associated set relationships
func MakeMoveTaskOrder(db *pop.Connection, assertions Assertions) models.MoveTaskOrder {
	moveOrder := assertions.MoveOrder
	ppm := assertions.PersonallyProcuredMove

	if isZeroUUID(moveOrder.ID) {
		moveOrder = MakeMoveOrder(db, assertions)
	}

	if isZeroUUID(ppm.ID) {
		ppm = MakePPM(db, assertions)
	}
	var referenceID *string

	moveTaskOrder := models.MoveTaskOrder{
		MoveOrder:          moveOrder,
		MoveOrderID:        moveOrder.ID,
		ReferenceID:        referenceID,
		PersonallyProcuredMoveID: ppm.ID,
		IsAvailableToPrime: false,
		IsCanceled:         false,
	}

	// Overwrite values with those from assertions
	mergeModels(&moveTaskOrder, assertions.MoveTaskOrder)

	mustCreate(db, &moveTaskOrder)

	return moveTaskOrder
}

// MakeDefaultMoveTaskOrder makes an MoveTaskOrder with default values
func MakeDefaultMoveTaskOrder(db *pop.Connection) models.MoveTaskOrder {
	return MakeMoveTaskOrder(db, Assertions{})
}
