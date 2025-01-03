package request_models

type AddressRequest struct {
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
	State  string `json:"state" binding:"required"`
	Zip    string `json:"zip" binding:"required"`
}
