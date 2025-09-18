package data

import (
	"time"
)

type Movie struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year_at"`
	Runtime   Runtime   `json:"runtime,string"`
	Genre     []string  `json:"genres"`
	Version   int32     `json:"version,omitempty"`
}
