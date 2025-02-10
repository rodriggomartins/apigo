package service

import (
	"api/src/configuration/logger"
	"api/src/configuration/rest_err"
	"api/src/model"
	"fmt"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("create user model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
