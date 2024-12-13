package domain

type user struct {
	ID   uint32 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
