package models

type Aluno struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Matricula string `json:"matricula"`
	TurmaIDs  []int  `json:"turma_ids"`
}
