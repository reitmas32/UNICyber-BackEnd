package schemas

type ComputerCreateSchema struct {
	Name   string `json:"name"`
	IdRoom string `json:"id_room"`
	Type   string `json:"type"`
}

type ComputerFindSchema struct {
	IdComputer string `json:"id_computer"`
}

type ComputerUpdateSchema struct {
	// Metadata
	//IdComputer string `gorm:"primaryKey"`
	//IdRoom     string

	// UI
	Pos_x float32 `json:"pos_x"`
	Pos_y float32 `json:"pos_y"`

	// Data
	Name    string `json:"name"`
	State   string `json:"state"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
