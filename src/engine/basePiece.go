package engine

// BasePiece - See interface Piece.
type BasePiece struct {
	self Piece
}

// GetColorAsRGB - See interface Piece.
func (p *BasePiece) GetColorAsRGB() (r uint8, g uint8, b uint8) {
	return 255, 255, 255
}

// GetFieldOccupation - See interface Piece.
func (p *BasePiece) GetFieldOccupation() PieceFieldOccupationFlags {
	return PieceRow1Column1
}

// IsFieldOccupied - See interface Piece.
func (p *BasePiece) IsFieldOccupied(field PieceFieldOccupationFlags) bool {
	return (p.self.GetFieldOccupation() & field) != 0
}

// IsFieldWithOffsetOccupied - See interface Piece.
func (p *BasePiece) IsFieldWithOffsetOccupied(rowOffset uint8, colOffset uint8) bool {
	// the origin itself
	if rowOffset == 0 && colOffset == 0 {
		return true
	}

	switch rowOffset {
	case 0:
		switch colOffset {
		case 0:
			return p.self.IsFieldOccupied(PieceRow1Column1)
		case 1:
			return p.self.IsFieldOccupied(PieceRow1Column2)
		case 2:
			return p.self.IsFieldOccupied(PieceRow1Column3)
		}

	case 1:
		switch colOffset {
		case 0:
			return p.self.IsFieldOccupied(PieceRow2Column1)
		case 1:
			return p.self.IsFieldOccupied(PieceRow2Column2)
		case 2:
			return p.self.IsFieldOccupied(PieceRow2Column3)
		}

	case 2:
		switch colOffset {
		case 0:
			return p.self.IsFieldOccupied(PieceRow3Column1)
		case 1:
			return p.self.IsFieldOccupied(PieceRow3Column2)
		case 2:
			return p.self.IsFieldOccupied(PieceRow3Column3)
		}

	case 3:
		switch colOffset {
		case 0:
			return p.self.IsFieldOccupied(PieceRow4Column1)
		case 1:
			return p.self.IsFieldOccupied(PieceRow4Column2)
		case 2:
			return p.self.IsFieldOccupied(PieceRow4Column3)
		}
	}

	// We should never reach this point, but just in case:
	return false
}
