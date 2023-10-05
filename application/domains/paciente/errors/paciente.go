package errors

import appErrors "github.com/rof20004/healthy-api/application/errors"

var (
	ErrIdInvalido              = appErrors.BuildCustomError(nil, 400, "Paciente com id inválido")
	ErrNomeObrigatorio         = appErrors.BuildCustomError(nil, 400, "Nome do paciente é obrigatório")
	ErrIdadeInvalida           = appErrors.BuildCustomError(nil, 400, "Idade do paciente precisa ser maior que 0")
	ErrPacientesNaoEncontrados = appErrors.BuildCustomError(nil, 400, "Não existe pacientes cadastrados")
)
