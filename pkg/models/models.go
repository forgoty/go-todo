package models

import (
	"github.com/forgoty/go-todo/pkg/infrastructure/logger"
	"github.com/forgoty/go-todo/pkg/web"
)

type ReqContext struct {
	web.Context
	logger *logger.Logger
}

func (c *ReqContext) Logger() logger.Logger {
	l := c.logger
	return *l
}
