package logger

import (
	"go.uber.org/zap"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Get() *zap.Logger {
	once.Do(func() {
		var err error

		logger, err = zap.NewDevelopment()
		//if production {
		//	logger, err = zap.NewProduction()
		//} else {
		//	logger, err = zap.NewDevelopment()
		//}

		if err != nil {
			panic(err)
		}
	})

	return logger
}
