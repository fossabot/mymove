package models_test

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/models"
	. "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ModelSuite) TestCreateNewMoveShow() {
	orders := testdatagen.MakeDefaultOrder(suite.DB())

	selectedMoveType := SelectedMoveTypeHHG

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	_, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")

	moves, moveErrs := GetMoveQueueItems(suite.DB(), "all")
	suite.Nil(moveErrs)
	suite.Len(moves, 1)
}

func (suite *ModelSuite) TestCreateNewMoveShowFalse() {
	orders := testdatagen.MakeDefaultOrder(suite.DB())

	selectedMoveType := SelectedMoveTypeHHG

	moveOptions := MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(false),
	}
	_, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "failed to validate move")

	moves, moveErrs := GetMoveQueueItems(suite.DB(), "all")
	suite.Nil(moveErrs)
	suite.Empty(moves)
}

func (suite *ModelSuite) TestShowPPMQueue() {
	all := map[string]bool{
		string(models.PPMStatusAPPROVED):         true,
		string(models.PPMStatusPAYMENTREQUESTED): true,
		string(models.PPMStatusCOMPLETED):        true,
		string(models.PPMStatusSUBMITTED):        true,
		string(models.PPMStatusDRAFT):            true,
	}

	new := map[string]bool{
		string(models.PPMStatusSUBMITTED): true,
		string(models.PPMStatusDRAFT):     true,
	}

	tests := []struct {
		input      string
		movesCount int
		want       map[string]bool
	}{
		{input: "new", movesCount: 2, want: new},
		{input: "ppm_payment_requested", movesCount: 1, want: map[string]bool{string(models.PPMStatusPAYMENTREQUESTED): true}},
		{input: "ppm_completed", movesCount: 1, want: map[string]bool{string(models.PPMStatusCOMPLETED): true}},
		{input: "ppm_approved", movesCount: 1, want: map[string]bool{string(models.PPMStatusAPPROVED): true}},
		{input: "all", movesCount: 5, want: all},
	}

	// Make PPMs with different statuses
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusAPPROVED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusPAYMENTREQUESTED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusCOMPLETED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			Status: models.MoveStatusSUBMITTED,
		},
	})
	testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			Status: models.MoveStatusAPPROVED,
		},
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			Status: models.PPMStatusSUBMITTED,
		},
	})

	for _, tc := range tests {
		moves, err := GetMoveQueueItems(suite.DB(), tc.input)

		suite.NoError(err)
		suite.Len(moves, tc.movesCount)
		for _, move := range moves {
			fmt.Printf("%+v", *move.PpmStatus)
			fmt.Println(tc.want[*move.PpmStatus])
			suite.True(tc.want[*move.PpmStatus])
		}
	}

}

func (suite *ModelSuite) TestQueueNotFound() {
	moves, moveErrs := GetMoveQueueItems(suite.DB(), "queue_not_found")
	suite.Equal(ErrFetchNotFound, moveErrs, "Expected not to find move queue items")
	suite.Empty(moves)
}
