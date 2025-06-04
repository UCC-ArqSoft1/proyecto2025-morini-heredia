package dto

type ActividadMinDTO struct {
	Id          uint   `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Cupo        uint   `json:"cupo"`
	Dia         string `json:"dia"`
	HoraInicio  string `json:"hora_inicio"`
	HoraFin     string `json:"hora_fin"`
	Instructor  string `json:"instructor"`
	Categoria   string `json:"categoria"`
}

type ActividadesMinDTO []ActividadMinDTO
