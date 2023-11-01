package dominio

import (
	"regexp"
)

type Cliente struct {
	CPF   string
	ID    string
	Nome  string
	Email string
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func NewCliente(cpf, id, nome, email string) (*Cliente, error) {
	if err := validarCPF(cpf); err != nil {
		return nil, err
	}

	if err := validarID(id); err != nil {
		return nil, err
	}

	if err := validarNome(nome); err != nil {
		return nil, err
	}

	if err := validarEmail(email); err != nil {
		return nil, err
	}

	return &Cliente{
		CPF:   cpf,
		ID:    id,
		Nome:  nome,
		Email: email,
	}, nil
}

func validarCPF(cpf string) error {
	if !cpfEhValido(cpf) {
		return &ValidationError{"CPF", "CPF inválido"}
	}
	return nil
}

func validarID(id string) error {
	if len(id) == 0 {
		return &ValidationError{"ID", "ID é obrigatório"}
	}
	return nil
}

func validarNome(nome string) error {
	if len(nome) == 0 {
		return &ValidationError{"Nome", "Nome é obrigatório"}
	}
	return nil
}

func validarEmail(email string) error {
	if !emailEhValido(email) {
		return &ValidationError{"Email", "Email inválido"}
	}
	return nil
}

func cpfEhValido(cpf string) bool {

	return len(cpf) == 11
}

func emailEhValido(email string) bool {

	emailRegex := regexp.MustCompile(`^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)*(\.[a-z]{2,})$`)
	return emailRegex.MatchString(email)
}
