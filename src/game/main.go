package main

import (
	"../engine"
)

func main() {
	matrix := engine.NewMatrix(20, 30)

	matrix.PlaceNewPieceAt(0, 0, engine.NewSquarePiece())
	matrix.PlaceNewPieceAt(10, 10, engine.NewLinePiece())
	matrix.PlaceNewPieceAt(9, 2, engine.NewTPiece())
	matrix.PlaceNewPieceAt(0, 4, engine.NewLLPiece())
	matrix.PlaceNewPieceAt(4, 4, engine.NewRLPiece())

	renderer := &engine.SimpleMatrixRenderer{matrix}
	renderer.Render()
}
