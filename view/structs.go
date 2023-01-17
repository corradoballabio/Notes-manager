package view

type Response struct {
	Code int					`json:"code"`
	Body interface{}	`json:"body"`
}

type NoteTable struct {
	Title string	`json:"title"`
	Body string		`json:"body"`
}