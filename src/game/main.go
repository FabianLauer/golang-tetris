package main

import (
	"math/rand"
	"time"

	"../engine"
)

func main() {
	matrix := engine.NewMatrix(20, 30)
	renderer := &engine.SimpleMatrixRenderer{matrix}

	spawner := engine.NewPieceSpawner(matrix, func() engine.Piece {
		rand.Seed(time.Now().Unix())
		index := rand.Intn(4)

		switch index {
		default:
		case 0:
			return engine.NewSquarePiece()
		case 1:
			return engine.NewLinePiece()
		case 2:
			return engine.NewLLPiece()
		case 3:
			return engine.NewRLPiece()
		case 4:
			return engine.NewTPiece()
		}

		return nil
	})

	spawner.Spawn()

	renderer.Render()
}
