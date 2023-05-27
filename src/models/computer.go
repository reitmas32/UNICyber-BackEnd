package models

type Computer struct {

	// Metadata
	idComputerUI string
	idRoom       string

	// UI
	pos_x int
	pos_y int

	// Data
	name         string
	state        string
	message      string
	typeComputer string
}
