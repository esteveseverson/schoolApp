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
	Semestre    int    `json:"semestre"`
	Ano         int    `json:"ano"`
	ProfessorID int    `json:"professor_id"`
}

type Aluno struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Matricula string `json:"matricula"`
	Turmas    []int  `json:"turmas"`
}

type Atividade struct {
	ID      int    `json:"id"`
	TurmaID int    `json:"turma_id"`
	Valor   int    `json:"valor"`
	Data    string `json:"data"`
}

type Nota struct {
	ID          int `json:"id"`
	AlunoID     int `json:"aluno_id"`
	AtividadeID int `json:"atividade_id"`
	Nota        int `json:"nota"`
}
