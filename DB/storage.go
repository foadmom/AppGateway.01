package DB

import (
	"AppGateway/service"
)

type Config struct {
}

type DBService interface {
	Connect(config Config) error
	Save(service.ServiceInfo) error
}

type PostGreSQL struct {
}

var dbService DBService

func GetDBService(DBService) DBService {
	return dbService
}

func (p *PostGreSQL) Connect(config Config) error {
	var _err error
	return _err
}

func (p *PostGreSQL) Save(service service.ServiceInfo) error {
	var _err error
	return _err
}
