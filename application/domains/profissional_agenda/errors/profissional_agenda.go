package errors

import appErrors "github.com/rof20004/healthy-api/application/errors"

var (
	ErrIdInvalido                        = appErrors.BuildCustomError(nil, 400, "Agenda com id inválido")
	ErrProfissionalIdObrigatorio         = appErrors.BuildCustomError(nil, 400, "Id do profissional é obrigatório")
	ErrDataInvalida                      = appErrors.BuildCustomError(nil, 400, "Data da agenda é inválida")
	ErrProfissionalAgendasNaoEncontradas = appErrors.BuildCustomError(nil, 400, "Não existe agendas cadastradas para o profissional")
)
