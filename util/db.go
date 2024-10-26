package util

import (
	"fmt"
	"school/config"
)

func ToDSN(opts *config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		opts.Host, opts.User, opts.Password, opts.Name, opts.Port,
	)
}
