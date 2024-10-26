package service

import (
	"school/repository"
)

type TokenService struct {
	Repo *repository.TokenRepository
}

func (service *TokenService) GetRefreshToken(userID uint) (string, error) {
	return service.Repo.GetRefreshToken(userID)
}

func (service *TokenService) SaveRefreshToken(userID uint, token string) error {
	return service.Repo.SaveRefreshToken(userID, token)
}
