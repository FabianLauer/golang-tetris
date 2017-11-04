package engine

// RLPiece - See interface Piece.
type RLPiece struct {
	BasePiece
}

func NewRLPiece() *RLPiece {
	piece := new(RLPiece)
	piece.self = piece
	return piece
}

// GetColorAsRGB - See interface Piece.
func (p *RLPiece) GetColorAsRGB() (r uint8, g uint8, b uint8) {
	return 200, 200, 200
}

// GetFieldOccupation - See interface Piece.
func (p *RLPiece) GetFieldOccupation() PieceFieldOccupationFlags {
	return PieceRow1Column1 | PieceRow1Column2 | PieceRow1Column3 | PieceRow2Column1
}
