package models

type JsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}
