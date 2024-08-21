package models

type Professor struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

type Turma struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Ano         int    `json:"ano"`
	ProfessorID int    `json:"professor_id"`
}

type Aluno struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Matricula string `json:"matricula"`
	TurmaID   int    `json:"turma_id"`
}

type Atividade struct {
	ID      int     `json:"id"`
	TurmaID int     `json:"turma_id"`
	Valor   float64 `json:"valor"`
	Data    string  `json:"data"`
}

type Nota struct {
	ID          int     `json:"id"`
	AlunoID     int     `json:"aluno_id"`
	AtividadeID int     `json:"atividade_id"`
	Valor       float64 `json:"nota"`
}
