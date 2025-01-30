package wsl

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/gookit/ini/v2"
)

func ResovleCommand(binary string, cfg *ini.Ini, originalArgs []string) (string, []string, error) {
	distro := strings.TrimSpace(cfg.String("distro", ""))
	user := strings.TrimSpace(cfg.String("user", ""))
	resolvedBinary := strings.TrimSpace(cfg.String("binary", binary))
	if len(distro) == 0 {
		return "", nil, errors.New("Distro is not set")
	}
	if resolvedBinary != binary {
		if !path.IsAbs(resolvedBinary) || path.Base(resolvedBinary) != binary {
			return "", nil, fmt.Errorf("Invalid binary: %s", resolvedBinary)
		}
	}
	args := []string{"-d", distro}
	if len(user) > 0 {
		args = append(args, "-u", user)
	}
	args = append(args, resolvedBinary)
	return "wsl", append(args, originalArgs...), nil
}
