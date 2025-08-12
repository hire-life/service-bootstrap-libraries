package jet

import (
	"github.com/hire-life/service-bootstrap-libraries/pkg/logger"
	"go.uber.org/zap"
	"os"
	"os/exec"
)

func Generate(conn string) {
	log := logger.Get()

	cmd := exec.Command("jet", "-dsn="+conn, "-schema=public", "-path=./.gen")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Info("Running jet generator", zap.String("url", conn))

	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to run jet generator", zap.Error(err))
	}

	log.Info("Jet generator completed successfully")
}
