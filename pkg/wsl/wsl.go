package wsl

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/gookit/ini/v2"
)

func ResovleCommand(targetCommand string, cfg *ini.Ini, originalArgs []string) (string, []string, error) {
	distro := strings.TrimSpace(cfg.String("distro", ""))
	user := strings.TrimSpace(cfg.String("user", ""))
	targetBinary := strings.TrimSpace(cfg.String("binary", targetCommand))
	if len(distro) == 0 {
		return "", nil, errors.New("distro is not set")
	}
	if targetBinary != targetCommand {
		if !path.IsAbs(targetBinary) || path.Base(targetBinary) != targetCommand {
			return "", nil, fmt.Errorf("invalid binary: %s", targetBinary)
		}
	}
	wslArgs := []string{"-d", distro}
	if len(user) > 0 {
		wslArgs = append(wslArgs, "-u", user)
	}
	wslArgs = append(wslArgs, "--exec", targetBinary)
	return "wsl", append(wslArgs, originalArgs...), nil
}
