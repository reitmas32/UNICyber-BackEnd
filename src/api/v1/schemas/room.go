package schemas

type RoomCreateSchema struct {
	Name          string `json:"name" binding:"required"`
	IdComputerLab uint   `json:"id_computer_lab" binding:"required"`
}

type RoomUpdateSchema struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}
