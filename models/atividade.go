package models

import "time"

type Atividade struct {
	ID      int       `json:"id"`
	TurmaID int       `json:"turma_id"`
	Valor   float64   `json:"valor"`
	Data    time.Time `json:"data"`
}
