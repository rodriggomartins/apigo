package service

import (
	"api/src/configuration/rest_err"
	"api/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface,
	*rest_err.RestErr) {
	return nil, nil
}
