package input

import "github.com/rof20004/healthy-api/application/domains/consulta/entities"

type CreateConsultaInputPort interface {
	CreateConsulta(consulta entities.Consulta) error
}
