package engine

import (
	"github.com/fatih/color"
)

type SimpleMatrixRenderer struct {
	// The matrix to render.
	Matrix *Matrix
}

var black = color.New(color.FgBlack)

// Render lets the SimpleMatrixRenderer render its matrix exactly once.
func (renderer *SimpleMatrixRenderer) Render() {
	renderer.renderCharacter("\n")

	for row := 0; row < renderer.Matrix.Height; row++ {
		for col := 0; col < renderer.Matrix.Width; col++ {
			if !renderer.Matrix.IsFieldOccupied(row, col) {
				renderer.renderUnoccupiedField()
			} else {
				renderer.renderOccupiedFieldAtCurrentCursorPos(row, col)
			}
		}

		renderer.renderCharacter("\n")
	}

	renderer.renderCharacter("\n")
}

// renderOccupiedFieldAtCurrentCursorPos renders a certain field of the matrix.
// It's important to only call this method from within Render since it renders
// the field at the current cursor position on the screen.
func (renderer *SimpleMatrixRenderer) renderOccupiedFieldAtCurrentCursorPos(row int, col int) {
	piece := renderer.Matrix.GetPieceOccupying(row, col)

	r, g, b := (*piece).GetColorAsRGB()
	total := int(r) + int(g) + int(b)

	character := "▧"
	characterColor := black

	if total > 230*3 {
		characterColor = color.New(color.FgBlue)
	} else if total > 180*3 {
		characterColor = color.New(color.FgYellow)
	} else if total > 120*3 {
		characterColor = color.New(color.FgRed)
	} else if total > 50*3 {
		characterColor = color.New(color.FgGreen)
	} else {
		characterColor = color.New(color.FgMagenta)
	}

	renderer.renderCharacterWithColor(character, characterColor)
}

// renderUnoccupiedField is used to render all fields that are not occupied
// by any Piece of the matrix.
func (renderer *SimpleMatrixRenderer) renderUnoccupiedField() {
	renderer.renderCharacterWithColor("▣", color.New(color.FgWhite))
}

// renderCharacter renders a normal character or control character on
// the current terminal at the current cursor position.
func (renderer *SimpleMatrixRenderer) renderCharacter(character string) {
	renderer.renderCharacterWithColor(character, black)
}

func (renderer *SimpleMatrixRenderer) renderCharacterWithColor(character string, c *color.Color) {
	c.Add(color.BgHiWhite).Printf("%s ", character)
}
