package models

import (
	"fmt"
	"time"
)

// CustomDate é um tipo para lidar com datas no formato desejado
type CustomDate struct {
	time.Time
}

// MarshalJSON converte CustomDate para o formato desejado
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	// Formato desejado para o JSON
	dateStr := cd.Time.Format("02-01-2006")
	return []byte(fmt.Sprintf(`"%s"`, dateStr)), nil
}

// UnmarshalJSON converte JSON para CustomDate
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// Remover aspas ao redor da data se houver
	strData := string(data)
	if strData == "null" {
		cd.Time = time.Time{}
		return nil
	}
	strData = strData[1 : len(strData)-1] // Remove aspas
	parsedTime, err := time.Parse("02-01-2006", strData)
	if err != nil {
		return err
	}
	cd.Time = parsedTime
	return nil
}

// Atividade representa uma atividade com seus atributos
type Atividade struct {
	ID          int        `json:"id"`
	TurmaID     int        `json:"turma_id"`
	Valor       float64    `json:"valor"`
	DataEntrega CustomDate `json:"data_entrega"`
}
