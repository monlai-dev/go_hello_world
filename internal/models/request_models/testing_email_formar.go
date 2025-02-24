package request_models

type TestingEmailFormat struct {
	Email   string `json:"email"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
}
