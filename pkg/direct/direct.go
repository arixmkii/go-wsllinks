package direct

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gookit/ini/v2"
)

func ResovleCommand(originalExe string, binary string, cfg *ini.Ini, originalArgs []string) (string, []string, error) {
	targetBinary := strings.TrimSpace(cfg.String("binary", binary))
	resolvedBinary, err := exec.LookPath(targetBinary)
	if err != nil {
		return "", nil, err
	}
	if resolvedBinary == originalExe {
		return "", nil, fmt.Errorf("Possible recursion detected calling: %s", resolvedBinary)
	}
	if strings.TrimSuffix(filepath.Base(resolvedBinary), filepath.Ext(resolvedBinary)) != binary {
		return "", nil, fmt.Errorf("Invalid binary: %s", resolvedBinary)
	}
	return resolvedBinary, originalArgs, nil
}
