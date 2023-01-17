package view

type Response struct {
	Code int			`json:"code"`
	Body interface{}	`json:"body"`
}

type TodoTable struct {
	Name string	`json:"name"`
	Todo string	`json:"todo"`
}