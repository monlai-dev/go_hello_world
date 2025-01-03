package response_models

type Response struct {
	ResponseCode int           `json:"response_code"`
	Message      string        `json:"message"`
	Data         []interface{} `json:"data"`
}
