package engine

// LinePiece - See interface Piece.
type LinePiece struct {
	BasePiece
}

func NewLinePiece() *LinePiece {
	piece := new(LinePiece)
	piece.self = piece
	return piece
}

// GetColorAsRGB - See interface Piece.
func (p *LinePiece) GetColorAsRGB() (r uint8, g uint8, b uint8) {
	return 100, 100, 100
}

// GetFieldOccupation - See interface Piece.
func (p *LinePiece) GetFieldOccupation() PieceFieldOccupationFlags {
	return PieceRow1Column1 | PieceRow1Column2 | PieceRow1Column3 | PieceRow1Column4
}
