package engine

// PieceFieldOccupationFlags is a type alias for bitmask flags that are used in the
// Piece interface. A combination of the flags is used to describe the shape of a
// Piece. See more in the Piece interface's GetFieldOccupation method.
type PieceFieldOccupationFlags uint32

const (
	PieceRow1Column1 PieceFieldOccupationFlags = 1 << iota
	PieceRow1Column2
	PieceRow1Column3
	PieceRow1Column4

	PieceRow2Column1
	PieceRow2Column2
	PieceRow2Column3
	PieceRow2Column4

	PieceRow3Column1
	PieceRow3Column2
	PieceRow3Column3
	PieceRow3Column4

	PieceRow4Column1
	PieceRow4Column2
	PieceRow4Column3
	PieceRow4Column4
)

// A Piece is the common base type of all pieces possible the game.
// It contains information about the piece's layout (and hence its dimensions),
// but also about its current appearance (color) and its current state (rotation),
type Piece interface {
	// Returns which fields in the 4x4 perimeter of the Piece are currently being
	// occupied by the piece. If the piece looks like this, for example...
	//
	//     X X X o
	//     o o X o
	//     o o o o
	//
	// ...this method would return:
	//
	//     PieceRow1Column1 | PieceRow1Column2 | PieceRow1Column3 | PieceRow2Column3
	//
	// The only 2 rules that apply are that the top left field *must* be occupied at
	// all times and that there must always be a connection between the fields.
	GetFieldOccupation() PieceFieldOccupationFlags

	// Check whether a field is occupied or not.
	IsFieldOccupied(field PieceFieldOccupationFlags) bool

	// Check whether a field with a certain offset to the Piece's origin is occupied.
	IsFieldWithOffsetOccupied(rowOffset uint8, colOffset uint8) bool

	// Returns the color with which the Piece should be rendered.
	GetColorAsRGB() (r uint8, g uint8, b uint8)
}
