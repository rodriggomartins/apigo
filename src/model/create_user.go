package model

import (
	"api/src/configuration/logger"
	"api/src/configuration/rest_err"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("create user model", zap.String("journey", "createUser"))
	ud.EncriptyPassword()
	return nil
}
