package models

type Aluno struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Matricula string `json:"matricula"`
	TurmaID   int    `json:"turma_id"`
}
