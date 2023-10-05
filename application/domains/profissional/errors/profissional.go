package errors

import appErrors "github.com/rof20004/healthy-api/application/errors"

var (
	ErrIdInvalido                  = appErrors.BuildCustomError(nil, 400, "Profissional com id inválido")
	ErrNomeObrigatorio             = appErrors.BuildCustomError(nil, 400, "Nome do profissional é obrigatório")
	ErrCPFObrigatorio              = appErrors.BuildCustomError(nil, 400, "CPF do profissional é obrigatório")
	ErrEmailInvalido               = appErrors.BuildCustomError(nil, 400, "E-mail do profissional inválido")
	ErrEmailObrigatorio            = appErrors.BuildCustomError(nil, 400, "E-mail do profissional é obrigatório")
	ErrFotoObrigatorio             = appErrors.BuildCustomError(nil, 400, "Foto do profissional é obrigatório")
	ErrCRPObrigatorio              = appErrors.BuildCustomError(nil, 400, "CRP do profissional é obrigatório")
	ErrSenhaObrigatoria            = appErrors.BuildCustomError(nil, 400, "Senha do profissional é obrigatória")
	ErrProfissionaisNaoEncontrados = appErrors.BuildCustomError(nil, 404, "Não existe profissionais cadastrados")
)
