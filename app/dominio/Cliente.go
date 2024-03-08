package dominio

type Cliente struct {
	CPF    string
	ID     string
	Nome   string
	Email  string
	Status string
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func NewCliente(cpf, id, nome, email, status string) (*Cliente, error) {

	return &Cliente{
		CPF:    cpf,
		ID:     id,
		Nome:   nome,
		Email:  email,
		Status: status,
	}, nil
}
