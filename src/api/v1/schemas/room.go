package schemas

type RoomCreateSchema struct {
	Name          string `json:"name"`
	IdComputerLab string `json:"id_computer_lab"`
}

func (c *RoomCreateSchema) isValid() bool {
	if c.Name == "" || c.IdComputerLab == "" {
		return false
	}
	return true
}

type RoomUpdateSchema struct {
	// Metadata

	// Data
	Name  string `json:"name"`
	Index int    `json:"index"`
}

func (c *RoomUpdateSchema) isValid() bool {
	return true
}
