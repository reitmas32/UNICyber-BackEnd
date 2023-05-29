package schemas

type ComputerCreateSchema struct {
	Name   string `json:"name"`
	IdRoom string `json:"id_room"`
	Type   string `json:"type"`
}

type ComputerFindSchema struct {
	IdComputer string `json:"id_computer"`
}

type ComputerRenameSchema struct {
	Name string `json:"name"`
}
