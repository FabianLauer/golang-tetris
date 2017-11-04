package engine

// TPiece - See interface Piece.
type TPiece struct {
	BasePiece
}

func NewTPiece() *TPiece {
	piece := new(TPiece)
	piece.self = piece
	return piece
}

// GetColorAsRGB - See interface Piece.
func (p *TPiece) GetColorAsRGB() (r uint8, g uint8, b uint8) {
	return 0, 100, 30
}

// GetFieldOccupation - See interface Piece.
func (p *TPiece) GetFieldOccupation() PieceFieldOccupationFlags {
	return PieceRow1Column1 | PieceRow1Column2 | PieceRow1Column3 | PieceRow2Column2
}
