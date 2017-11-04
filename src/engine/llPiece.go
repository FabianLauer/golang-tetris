package engine

// LLPiece - See interface Piece.
type LLPiece struct {
	BasePiece
}

func NewLLPiece() *LLPiece {
	piece := new(LLPiece)
	piece.self = piece
	return piece
}

// GetColorAsRGB - See interface Piece.
func (p *LLPiece) GetColorAsRGB() (r uint8, g uint8, b uint8) {
	return 255, 255, 255
}

// GetFieldOccupation - See interface Piece.
func (p *LLPiece) GetFieldOccupation() PieceFieldOccupationFlags {
	return PieceRow1Column1 | PieceRow1Column2 | PieceRow1Column3 | PieceRow2Column3
}
