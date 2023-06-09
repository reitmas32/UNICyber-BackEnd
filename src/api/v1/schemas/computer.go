package schemas

// @Summary Ejemplo de estructura A
// @Description Descripción de la estructura A
// @Tags Computer
type ComputerCreateSchema struct {
	Name   string `json:"name" example:"Equipo 2" binding:"required"`
	IdRoom uint   `json:"id_room" example:"3" binding:"required"`
	Type   string `json:"type" example:"Impresora" binding:"required"`
}

type ComputerFindSchema struct {
	IdComputer uint `json:"id_computer" binding:"required"`
}

type ComputerUpdateSchema struct {
	// Metadata
	//IdComputer string `gorm:"primaryKey"`
	//IdRoom     string

	// UI
	Pos_x float32 `json:"pos_x" example:"156.29"`
	Pos_y float32 `json:"pos_y" example:"157.30"`

	// Data
	Name    string `json:"name" example:"Equipo 3"`
	State   string `json:"state" example:"Disponible"`
	Message string `json:"message" example:"Solo falta instalar DevC++"`
	Type    string `json:"type" example:"Prestamo"`
}

type SetStateSchema struct {
	IdState uint `json:"id_state" example:"1"`
}
