package util

import (
	"be-golang-chapter-36-implem/helper"
	"os"

	"go.uber.org/zap"
)

func LoggerInit() (*zap.Logger, error) {

	var logger *zap.Logger
	var err error
	if helper.StringToBool(os.Getenv("DEBUG")) {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewDevelopment()

	}

	return logger, err
}
