package models

type Header struct {
	CreatedAt string
	Name      string
	Email     string
	Password  []byte
}
