package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	testcases := []struct {
		w      int
		h      int
		result bool
	}{
		{4, 5, true},
		{1, 1, true},
		{0, 1, false},
		{4, 0, false},
		{4, -1, false},
		{-3, 5, false},
	}

	for _, testcase := range testcases {
		board, err := NewBoard(testcase.w, testcase.h)

		if testcase.result {
			assert.NotNil(t, board)
			assert.Nil(t, err)

			assert.Equal(t, board.Width, testcase.w)
			assert.Equal(t, board.Height, testcase.h)
		} else {
			assert.Nil(t, board)
			assert.NotNil(t, err)
		}
	}
}

func TestGetPieceAt(t *testing.T) {
	board, _ := NewBoard(4, 5)
	piece, _ := NewPiece(10, 2, 2)
	board.AddPiece(piece, 1, 1)

	testcases := []struct {
		x      int
		y      int
		err    bool
		result bool
	}{
		{x: 1, y: 1, err: false, result: true},
		{x: 0, y: 1, err: false, result: false},
		{x: 1, y: 0, err: false, result: false},
		{x: -1, y: 1, err: true, result: false},
		{x: 1, y: -6, err: true, result: false},
		{x: 6, y: 1, err: true, result: false},
		{x: 1, y: 19, err: true, result: false},
		{x: 2, y: 1, err: false, result: true},
		{x: 2, y: 2, err: false, result: true},
	}

	for _, testcase := range testcases {
		foundPiece, err := board.GetPieceAt(testcase.x, testcase.y)
		if testcase.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		if testcase.result {
			assert.Equal(t, piece, foundPiece)
		} else {
			assert.Nil(t, foundPiece)
		}
	}
}

func TestAddPiece(t *testing.T) {
	testcases := []struct {
		pID    int
		pW     int
		pH     int
		x      int
		y      int
		result bool
	}{
		{pID: 10, pW: 1, pH: 1, x: 0, y: 0, result: true},   // valid addition
		{pID: 20, pW: 1, pH: 1, x: 4, y: 0, result: false},  // out of bounds
		{pID: 30, pW: 1, pH: 1, x: 1, y: 6, result: false},  // out of bounds
		{pID: 40, pW: 1, pH: 1, x: -1, y: 0, result: false}, // out of bounds
		{pID: 50, pW: 1, pH: 1, x: 0, y: -1, result: false}, // out of bounds
		{pID: 60, pW: 5, pH: 1, x: 0, y: 0, result: false},  // piece that's too large
		{pID: 70, pW: 1, pH: 6, x: 0, y: 0, result: false},  // piece that's too large
	}

	for _, testcase := range testcases {
		board, _ := NewBoard(4, 5)
		piece, _ := NewPiece(testcase.pID, testcase.pW, testcase.pH)
		err := board.AddPiece(piece, testcase.x, testcase.y)

		if testcase.result {
			assert.Nil(t, err)

			for x := 0; x < board.Width; x++ {
				for y := 0; y < board.Height; y++ {
					var expectPiece *Piece
					if x >= testcase.x && x < testcase.x+testcase.pW && y >= testcase.y && y < testcase.y+testcase.pH {
						expectPiece = piece
					} else {
						expectPiece = nil
					}
					actualPiece, _ := board.GetPieceAt(x, y)
					assert.Equal(t, expectPiece, actualPiece, "Unexpected piece at %d, %d", x, y)
				}
			}

		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestAddSecondPiece(t *testing.T) {
	testcases := []struct {
		pID    int
		pW     int
		pH     int
		x      int
		y      int
		result bool
	}{
		{pID: 20, pW: 1, pH: 1, x: 2, y: 0, result: true},  // adding besides first piece
		{pID: 30, pW: 1, pH: 1, x: 0, y: 2, result: true},  // adding below first piece
		{pID: 40, pW: 1, pH: 1, x: 1, y: 1, result: false}, // overlapping first piece
	}

	for _, testcase := range testcases {
		board, _ := NewBoard(4, 5)
		firstPiece, _ := NewPiece(10, 2, 2)
		firstPieceX := 0
		firstPieceY := 0
		board.AddPiece(firstPiece, firstPieceX, firstPieceY)

		piece, _ := NewPiece(testcase.pID, testcase.pW, testcase.pH)
		err := board.AddPiece(piece, testcase.x, testcase.y)

		if testcase.result {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)

			for x := 0; x < board.Width; x++ {
				for y := 0; y < board.Height; y++ {
					var expectPiece *Piece
					if x >= firstPieceX && x < firstPieceX+firstPiece.Width && y >= firstPieceY && y < firstPieceY+firstPiece.Height {
						expectPiece = firstPiece
					} else {
						expectPiece = nil
					}
					actualPiece, _ := board.GetPieceAt(x, y)
					assert.Equal(t, expectPiece, actualPiece)
				}
			}
		}
	}
}

func TestRemovePiece(t *testing.T) {
	testcases := []struct {
		id int
		w  int
		h  int
		x  int
		y  int
	}{
		{id: 10, w: 1, h: 1, x: 0, y: 0},   // valid piece
		{id: 20, w: 1, h: 1, x: -1, y: -1}, // piece that can't be added
	}

	for _, testcase := range testcases {
		board, _ := NewBoard(4, 5)
		piece, _ := NewPiece(testcase.id, testcase.w, testcase.h)

		board.AddPiece(piece, testcase.x, testcase.y)

		board.RemovePiece(piece)

		for x := 0; x < board.Width; x++ {
			for y := 0; y < board.Height; y++ {
				p, _ := board.GetPieceAt(x, y)
				assert.Nil(t, p)
			}
		}
	}
}

func TestMoveOnlyPiece(t *testing.T) {
	board, _ := NewBoard(4, 5)
	piece, _ := NewPiece(10, 2, 2)
	board.AddPiece(piece, 2, 0) // piece is nestled in top right corner

	testcases := []struct {
		orient Orientation
		err    bool
	}{
		{orient: Up, err: true},
		{orient: Right, err: true},
		{orient: Down, err: false},
		{orient: Left, err: false},
	}

	for _, testcase := range testcases {
		err := board.MovePiece(piece, testcase.orient)
		if testcase.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.NotEqual(t, Location{x: 2, y: 0}, board.pieces[piece])
		}
	}
}

func TestMovePiece(t *testing.T) {
	board, _ := NewBoard(4, 5)
	firstPiece, _ := NewPiece(10, 2, 2)
	board.AddPiece(firstPiece, 2, 0) // piece is nestled in top right corner
	piece, _ := NewPiece(20, 1, 1)
	board.AddPiece(piece, 1, 1)

	testcases := []struct {
		orient Orientation
		err    bool
	}{
		{orient: Up, err: false},
		{orient: Right, err: true},
		{orient: Down, err: false},
		{orient: Left, err: false},
	}

	for _, testcase := range testcases {
		err := board.MovePiece(piece, testcase.orient)
		if testcase.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.NotEqual(t, Location{x: 2, y: 0}, board.pieces[piece])
		}
	}
}

func TestSetGoal(t *testing.T) {
	board, _ := NewBoard(4, 5)
	firstPiece, _ := NewPiece(10, 2, 2)
	board.AddPiece(firstPiece, 2, 0) // piece is nestled in top right corner
	secondPiece, _ := NewPiece(20, 1, 1)
	board.AddPiece(secondPiece, 1, 1)
	notInBoard, _ := NewPiece(30, 3, 3)

	testcases := []struct {
		piece *Piece
		x, y  int
		err   bool
	}{
		{notInBoard, 1, 1, true},
		{firstPiece, 0, 0, false}, // piece is in top left corner
		{firstPiece, 2, 3, false}, // piece is in bottom right corner
		{firstPiece, 1, 2, false},
		{firstPiece, -1, 2, true},
		{firstPiece, 2, -4, true},
		{firstPiece, -2, -4, true},
		{firstPiece, -2, -4, true},
		{firstPiece, 3, 3, true}, // piece sticks out of board on right
		{firstPiece, 2, 4, true}, // piece sticks out of board on bottom
	}

	for _, testcase := range testcases {
		err := board.SetGoal(testcase.piece, testcase.x, testcase.y)
		if testcase.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestIsSolved(t *testing.T) {
	board, _ := NewBoard(4, 5)
	piece, _ := NewPiece(10, 2, 2)
	board.AddPiece(piece, 2, 0)
	_ = board.SetGoal(piece, 1, 3) // goal is bottom middle

	testcases := []struct {
		x, y     int
		isSovled bool
	}{
		{1, 1, false},
		{1, 3, true},
	}

	for _, testcase := range testcases {
		board.RemovePiece(piece)
		board.AddPiece(piece, testcase.x, testcase.y)
		assert.Equal(t, testcase.isSovled, board.IsSolved(), "Isn't in right solved state")
	}
}
