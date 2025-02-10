package initialize

import "go.uber.org/zap"

func InitiallizeLogger() {
	logger,_ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}