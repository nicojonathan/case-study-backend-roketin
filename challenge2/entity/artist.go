package entity

type Artist struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
