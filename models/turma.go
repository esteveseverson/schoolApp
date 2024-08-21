package models

type Turma struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Ano         int    `json:"ano"`
	Semestre    int    `json:"semestre"`
	ProfessorID int    `json:"professor_id"`
}
