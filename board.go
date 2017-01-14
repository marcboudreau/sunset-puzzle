package main

import (
	"errors"
	"fmt"
)

// Orientation is a type that indicates in which direction the piece is moved.
type Orientation int

// Location is a type that represents a 2 dimensional coordinate.
type Location struct {
	x int
	y int
}

// Direction constants.
const (
	Up Orientation = iota
	Right
	Down
	Left
)

// Board is a structure that represents the puzzle board.
type Board struct {
	// Width contains the width of the puzzle board.
	Width int

	// Height contains the height of the puzzle board.
	Height int

	slots [][]*Piece

	pieces map[*Piece]Location
}

// NewBoard constructs a new empty puzzle board.
func NewBoard(width, height int) (*Board, error) {
	if width < 1 {
		return nil, errors.New("Cannot create a board with a width that is less than 1.")
	}

	if height < 1 {
		return nil, errors.New("Cannot create a board with a height that is less than 1.")
	}

	slotsValue := make([][]*Piece, width)
	for i := range slotsValue {
		slotsValue[i] = make([]*Piece, height)
	}

	return &Board{
		Width:  width,
		Height: height,
		slots:  slotsValue,
		pieces: make(map[*Piece]Location),
	}, nil
}

// AddPiece adds a Piece instance at the specified location.  If the piece
// would overlap with an existing piece, it is not placed and an error is
// returned.
func (p *Board) AddPiece(piece *Piece, x, y int) error {
	// Validate location inputs
	if x < 0 || x > p.Width-piece.Width || y < 0 || y > p.Height-piece.Height {
		return fmt.Errorf("Piece %d cannot be added at %d, %d because it would not fit within the board.", piece.ID, x, y)
	}

	for i := 0; i < piece.Width; i++ {
		for j := 0; j < piece.Height; j++ {
			slotPiece := p.slots[x+i][y+j]
			if slotPiece != nil {
				return fmt.Errorf("Piece %d overlaps piece %d at %d, %d\n", piece.ID, slotPiece, x+i, y+j)
			}
		}
	}

	p.pieces[piece] = Location{
		x: x,
		y: y,
	}

	for i := 0; i < piece.Width; i++ {
		for j := 0; j < piece.Height; j++ {
			p.slots[x+i][y+j] = piece
		}
	}

	return nil
}

// MovePiece moves the specified piece by 1 square in the given direction.
func (p *Board) MovePiece(piece *Piece, orientation Orientation) error {
	// Check if the piece can be moved.
	l := p.pieces[piece]
	x := l.x
	y := l.y

	switch orientation {
	case Up:
		y--
		break
	case Right:
		x++
		break
	case Down:
		y++
		break
	case Left:
		x--
		break
	}

	if x < 0 || x+piece.Width > p.Width || y < 0 || y+piece.Height > p.Height {
		return fmt.Errorf("Moving piece %d (%d, %d - %d, %d) would put it outside the bounds of the puzzle board (0, 0 - %d, %d)", piece.ID, x, y, x+piece.Width, y+piece.Height, p.Width, p.Height)
	}

	for i := 0; i < piece.Width; i++ {
		for j := 0; j < piece.Height; j++ {
			slotPiece := p.slots[x+i][y+j]
			if slotPiece != nil && slotPiece != piece {
				return fmt.Errorf("Moving piece %d would cause it to overlap piece %d", piece.ID, slotPiece.ID)
			}
		}
	}

	p.RemovePiece(piece)
	p.AddPiece(piece, x, y)

	return nil
}

// RemovePiece removes a Piece instance from the puzzle board.
func (p *Board) RemovePiece(piece *Piece) {
	if l, ok := p.pieces[piece]; ok {
		for x := 0; x < piece.Width; x++ {
			for y := 0; y < piece.Height; y++ {
				p.slots[x+l.x][y+l.y] = nil
			}
		}

		delete(p.pieces, piece)
	}
}

// GetPieceAt returns the Piece instance that occupies the specified location.
func (p *Board) GetPieceAt(x, y int) (*Piece, error) {
	if x < 0 || x >= p.Width {
		return nil, fmt.Errorf("The specified x coordinate %d is invalid.", x)
	}

	if y < 0 || y >= p.Height {
		return nil, fmt.Errorf("The specified y coordinate %d is invalid.", y)
	}

	return p.slots[x][y], nil
}
