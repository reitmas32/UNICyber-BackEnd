package schemas

type RoomCreateSchema struct {
	Name          string `json:"name"`
	IdComputerLab string `json:"id_computer_lab"`
}

type RoomUpdateSchema struct {
	// Metadata

	// Data
	Name  string `json:"name"`
	Index string `json:"index"`
}
