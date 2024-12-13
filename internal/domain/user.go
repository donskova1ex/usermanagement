package domain

type user struct {
	UUID         string `json:"uuid" db:"uuid"`
	ID           uint32 `json:"id" db:"id"`
	login        string `json:"login" db:"login"`
	passwordhash string `json:"passwordhash" db:"passwordhash"`
	firstname    string `json:"firstname"`
}
