package engine

import (
	"math/rand"
	"time"
)

// A PieceSpawner is an object that spawns Pieces on a matrix.
type PieceSpawner struct {
	// The matrix on which Pieces should be spawned.
	matrix *Matrix

	// A function that generates a new Piece.
	generateNewPiece func() Piece

	// The piece that will be spawned in the next call to method `Spawn`.
	upcomingPiece *Piece
	// The piece that was spawned in the most recent call to method `Spawn`.
	latestPiece *Piece
}

// NewPieceSpawner creates a new instance of PieceSpawner.
func NewPieceSpawner(matrix *Matrix, generateNewPiece func() Piece) *PieceSpawner {
	spawner := new(PieceSpawner)

	spawner.matrix = matrix
	spawner.generateNewPiece = generateNewPiece

	// Call `prepareUpcomingPiece` once so the first call to
	// `Spawn` will work as expected.
	spawner.prepareUpcomingPiece()

	return spawner
}

// prepareNextPiece generates a new piece and assigns it to
// the `upcomingPiece` property.
func (spawner *PieceSpawner) prepareUpcomingPiece() {
	piece := spawner.generateNewPiece()
	spawner.upcomingPiece = &piece
}

// randIntBetween generates a random integer between to inclusive values.
func (spawner *PieceSpawner) randIntBetween(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Spawn spawns a single Piece onto the matrix and prepares
// another Piece (the "upcoming Piece"), which will be spawned
// to the Matrix the next time the `Span` method is called.
func (spawner *PieceSpawner) Spawn() {
	// Swap the previous "upcoming Piece" so we can actually spawn it,
	// then generate a new upcoming Piece.
	spawner.latestPiece = spawner.upcomingPiece
	spawner.prepareUpcomingPiece()

	// Now, find a position to spawn. Since we always spawn in the
	// top row of the matrix, we only need to generate a random
	// column to spawn in.
	column := spawner.randIntBetween(0, spawner.matrix.Width-1)
	spawner.matrix.PlaceNewPieceAt(0, column, *spawner.latestPiece)
}
