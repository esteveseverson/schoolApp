package models

type Counter struct {
	ID  string `json:"id" bson:"_id"`
	Seq int    `json:"seq" bson:"seq"`
}

type Professor struct {
	ID    int    `json:"id" bson:"id"`
	Nome  string `json:"nome" bson:"nome"`
	Email string `json:"email" bson:"email"`
	CPF   string `json:"cpf" bson:"cpf"`
}

type Turma struct {
	ID          int    `json:"id" bson:"id"`
	Nome        string `json:"nome" bson:"nome"`
	Semestre    int    `json:"semestre" bson:"semestre"`
	Ano         int    `json:"ano" bson:"ano"`
	ProfessorID int    `json:"professor_id" bson:"professor_id"`
}

type Aluno struct {
	ID        int    `json:"id" bson:"id"`
	Nome      string `json:"nome" bson:"nome"`
	Matricula string `json:"matricula" bson:"matricula"`
	Turmas    []int  `json:"turmas" bson:"turmas"`
}

type Atividade struct {
	ID      int     `json:"id" bson:"id"`
	TurmaID int     `json:"turma_id" bson:"turma_id"`
	Valor   float64 `json:"valor" bson:"valor"`
	Data    string  `json:"data" bson:"data"`
}

type Nota struct {
	ID          int     `json:"id" bson:"id"`
	AlunoID     int     `json:"aluno_id" bson:"aluno_id"`
	AtividadeID int     `json:"atividade_id" bson:"atividade_id"`
	Nota        float64 `json:"nota" bson:"nota"`
}
