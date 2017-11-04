package engine

// SquarePiece - See interface Piece.
type SquarePiece struct {
	BasePiece
}

func NewSquarePiece() *SquarePiece {
	piece := new(SquarePiece)
	piece.self = piece
	return piece
}

// GetColorAsRGB - See interface Piece.
func (p *SquarePiece) GetColorAsRGB() (r uint8, g uint8, b uint8) {
	return 200, 200, 200
}

// GetFieldOccupation - See interface Piece.
func (p *SquarePiece) GetFieldOccupation() PieceFieldOccupationFlags {
	return PieceRow1Column1 | PieceRow1Column2 | PieceRow2Column1 | PieceRow2Column2
}
