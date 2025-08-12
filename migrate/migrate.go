package migrate

import (
	"github.com/hire-life/service-bootstrap-libraries/logger"
	"go.uber.org/zap"
	"os"
	"os/exec"
)

func Run(conn string) {
	log := logger.Get()

	source := "migrations"

	cmd := exec.Command("migrate", "-database", conn, "-path", source, "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Info("Running migration",
		zap.String("source", source), zap.String("path", source), zap.String("url", conn))

	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to run migration", zap.Error(err))
	}

	log.Info("Migration completed successfully")
}
