package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewPiece(t *testing.T) {
	testcases := []struct {
		i      int
		w      int
		h      int
		result bool
	}{
		{i: 50, w: 1, h: 1, result: true},
		{i: 100, w: 0, h: 1, result: false},
		{i: 150, w: 2, h: 0, result: false},
		{i: 200, w: -1, h: 1, result: false},
		{i: 0, w: 1, h: 2, result: true},
		{i: -100, w: 2, h: 2, result: true},
	}

	for _, testcase := range testcases {
		piece, err := NewPiece(testcase.i, testcase.w, testcase.h)

		if testcase.result {
			assert.NotNil(t, piece)
			assert.Nil(t, err)

			assert.Equal(t, testcase.i, piece.ID)
			assert.Equal(t, testcase.w, piece.Width)
			assert.Equal(t, testcase.h, piece.Height)
		} else {
			assert.Nil(t, piece)
			assert.NotNil(t, err)
		}
	}
}

func TestNewPieceWithDimensionSetToZero(t *testing.T) {

}
