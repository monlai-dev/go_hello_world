package response_models

import ()

type LoginResponse struct {
	Jwt string `json:"jwt"`
}