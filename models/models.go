package models

import (
	"gorm.io/gorm"
)

type Professor struct {
	gorm.Model
	Nome  string
	Email string
	CPF   string
}

type Turma struct {
	gorm.Model
	Nome        string
	Semestre    int
	Ano         int
	ProfessorID uint
	Professor   Professor
}

type Aluno struct {
	gorm.Model
	Nome      string
	Matricula string
	Turmas    []Turma `gorm:"many2many:aluno_turmas"`
}

type Atividade struct {
	gorm.Model
	TurmaID uint
	Valor   int
	Data    string
}

type Nota struct {
	gorm.Model
	AlunoID     uint
	AtividadeID uint
	Nota        int
}
