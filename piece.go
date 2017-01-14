package main

import "errors"

// Piece is a structure that represents a moveable piece in the puzzle.
type Piece struct {
	// ID contains an identifier for this piece.
	ID int

	// Width contains the width of this piece.
	Width int

	// Height contains the height of this piece.
	Height int
}

// NewPiece constructs a new Piece instance of the specified dimensions.
func NewPiece(id, width, height int) (*Piece, error) {
	if width < 1 {
		return nil, errors.New("Cannot create a piece with a width that is less than 1.")
	}

	if height < 1 {
		return nil, errors.New("Cannot create a piece with a height that is less than 1.")
	}

	return &Piece{
		ID:     id,
		Width:  width,
		Height: height,
	}, nil
}
