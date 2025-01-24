package service

import (
	"api/src/configuration/rest_err"
	"api/src/model"
)

func (*userDomainService) UpdateUser(
	userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}
