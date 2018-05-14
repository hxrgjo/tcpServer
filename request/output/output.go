package output

type Status struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Datetime interface{} `json:"datetime"`
}

type Output struct {
	Data   interface{} `json:"data"`
	Status `json:"status"`
}
