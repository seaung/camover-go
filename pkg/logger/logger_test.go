package logger_test

import (
	"testing"

	"github.com/seaung/camover-go/pkg/logger"
)

func TestLogger(t *testing.T) {
    logger.NewLogger().Success("Success")
    logger.NewLogger().Errorw("Error")
    logger.NewLogger().Info("Info")
    logger.NewLogger().Process("Process")
    logger.NewLogger().Warning("Warning")
}
