package repository

import (
	"school/models"

	"gorm.io/gorm"
)

type TokenRepository struct {
	DB *gorm.DB
}

func (repo *TokenRepository) GetRefreshToken(userID uint) (string, error) {
	token := models.RefreshToken{
		UserID:  userID,
		Revoked: false,
	}

	err := repo.DB.Model(&models.RefreshToken{}).Find(&token).Error
	if err != nil {
		return "", err
	}
	if (token != models.RefreshToken{}) {
		return token.Token, nil
	}
	return "", nil
}

func (repo *TokenRepository) SaveRefreshToken(userID uint, token string) error {
	rtoken := models.RefreshToken{
		UserID:  userID,
		Token:   token,
		Revoked: false,
	}
	return repo.DB.Create(&rtoken).Error
}
