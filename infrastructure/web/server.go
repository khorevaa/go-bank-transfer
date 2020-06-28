package web

import (
	"errors"

	"github.com/gsabadini/go-bank-transfer/infrastructure/database"
	"github.com/gsabadini/go-bank-transfer/infrastructure/logger"
	"github.com/gsabadini/go-bank-transfer/infrastructure/validator"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid web server instance")
)

const (
	InstanceGorillaMux int = iota
	InstanceGin
)

//NewWebServer
func NewWebServer(
	instance int,
	log logger.Logger,
	dbConnSQL database.SQLHandler,
	dbConnNoSQL database.NoSQLHandler,
	validation validator.Validator,
	port Port,
) (Server, error) {
	switch instance {
	case InstanceGorillaMux:
		return NewGorillaMux(log, dbConnSQL, validation, port), nil
	case InstanceGin:
		return NewGin(log, dbConnNoSQL, port), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
