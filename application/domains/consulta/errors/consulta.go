package errors

import appErrors "github.com/rof20004/healthy-api/application/errors"

var (
	ErrIdInvalido                = appErrors.BuildCustomError(nil, 400, "Agenda com id inválido")
	ErrPacienteIdObrigatorio     = appErrors.BuildCustomError(nil, 400, "Id do paciente é obrigatório")
	ErrProfissionalIdObrigatorio = appErrors.BuildCustomError(nil, 400, "Id do profissional é obrigatório")
	ErrDataInvalida              = appErrors.BuildCustomError(nil, 400, "Data da agenda é inválida")
	ErrConsultasNaoEncontradas   = appErrors.BuildCustomError(nil, 400, "Não existe consultas cadastradas")
)
