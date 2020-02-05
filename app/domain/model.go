package domain

import "time"

type Placar struct {
	GolsTimeDaCasa    int `json:"golsTimeDaCasa"`
	GolsTimeVisitante int `json:"golsTimeVisitante"`
}

type Jogo struct {
	ID         int       `json:"id"`
	Data       time.Time `json:"data"`
	Local      string    `json:"local"`
	TimeDaCasa string    `json:"timeDaCasa"`
	Visitante  string    `json:"timeVisitante"`
	Placar     Placar    `json:"placar"`
}

type Aposta struct {
	ID     int    `json:"id"`
	IDJogo int    `json:"idJogo"`
	Placar Placar `json:"placar"`
}
