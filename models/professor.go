package models

type Professor struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}
