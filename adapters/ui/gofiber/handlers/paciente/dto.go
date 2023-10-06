package paciente

import pacienteEntities "github.com/rof20004/healthy-api/application/domains/paciente/entities"

type CreatePacienteRequest struct {
	Nome   string `json:"nome"`
	Avatar string `json:"avatar" format:"image base64"`
	Idade  int64  `json:"idade"`
}

func (c CreatePacienteRequest) ToEntity() pacienteEntities.Paciente {
	return pacienteEntities.NewPaciente(c.Nome, c.Idade, c.Avatar)
}
