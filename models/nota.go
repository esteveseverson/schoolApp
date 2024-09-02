package models

type Nota struct {
	ID          int     `json:"id"`
	AlunoID     int     `json:"aluno_id"`
	AtividadeID int     `json:"atividade_id"`
	TurmaID     int     `json:"turma_id"`
	ValorTotal  float64 `json:"valor_total"`
	ValorObtido float64 `json:"valor_obtido"`
}
