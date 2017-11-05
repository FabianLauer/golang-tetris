package engine

// A Matrix holds information about where each game piece is located.
// Use function NewMatrix to construct.
type Matrix struct {
	Width  int
	Height int
	pieces [][]*Piece
}

// NewMatrix creates a new Matrix instance that can be used right away.
func NewMatrix(width int, height int) *Matrix {
	matrix := new(Matrix)

	matrix.Width = width
	matrix.Height = height
	matrix.pieces = make([][]*Piece, height)

	for row := 0; row < matrix.Height; row++ {
		matrix.pieces[row] = make([]*Piece, width)
	}

	return matrix
}

// ForEachField runs a function once per field.
func (m *Matrix) ForEachField(fn func(row int, col int)) {
	for row := 0; row < m.Height; row++ {
		for col := 0; col < m.Width; col++ {
			fn(row, col)
		}
	}
}

// IsFieldOccupied checks if a certain column of a certain row of the
// matrix is occupied by a Piece.
func (m *Matrix) IsFieldOccupied(row int, col int) bool {
	return m.GetPieceOccupying(row, col) != nil
}

// GetPieceOccupying returns the Piece that currently occupies a certain
// row and column of the Matrix.
func (m *Matrix) GetPieceOccupying(row int, col int) *Piece {
	// Check if a Piece's origin is at the requested row and column and
	// return that piece if yes.
	piece := m.pieces[row][col]
	if piece != nil {
		return piece
	}

	// Even though no Piece's origin is at the row and column, there
	// could be neighbouring fields with a Piece that occupies the
	// requested row and column.
	// We loop through all neighbouring fields and check if any of the
	// fields is the origin of a Piece. If we find such a Piece, we
	// check if it also occupies the original row and column we're
	// looking for in this method.
	for rowOffset := uint8(0); rowOffset <= 3; rowOffset++ {
		if row-int(rowOffset) < 0 {
			continue
		}

		for colOffset := uint8(0); colOffset <= 3; colOffset++ {
			if col-int(colOffset) < 0 {
				continue
			}

			piece = m.pieces[row-int(rowOffset)][col-int(colOffset)]

			if piece == nil {
				continue
			}

			if (*piece).IsFieldWithOffsetOccupied(rowOffset, colOffset) {
				return piece
			}
		}
	}

	return nil
}

// PlaceNewPieceAt places a new piece at a certain row and column of the Matrix.
func (m *Matrix) PlaceNewPieceAt(row int, col int, piece Piece) {
	m.pieces[row][col] = &piece
}
