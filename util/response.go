package util

import "school/models"

func GenerateResponse(ok bool, message string, result any) models.JsonResponse {
	return models.JsonResponse{
		OK:      ok,
		Message: message,
		Result:  result,
	}
}
