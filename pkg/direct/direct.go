package direct

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gookit/ini/v2"
)

func ResovleCommand(originalExe string, targetCommand string, cfg *ini.Ini, originalArgs []string) (string, []string, error) {
	originalExeDir := filepath.Dir(originalExe)
	targetBinary := strings.TrimSpace(cfg.String("binary", targetCommand))
	if !filepath.IsAbs(targetBinary) {
		targetBinary = filepath.Join(originalExeDir, targetBinary)
	}
	resolvedBinary, err := exec.LookPath(targetBinary)
	if err != nil {
		return "", nil, err
	}
	if resolvedBinary == originalExe {
		return "", nil, fmt.Errorf("Possible recursion detected calling: %s", resolvedBinary)
	}
	if strings.TrimSuffix(filepath.Base(resolvedBinary), filepath.Ext(resolvedBinary)) != targetCommand {
		return "", nil, fmt.Errorf("Invalid binary: %s", resolvedBinary)
	}
	return resolvedBinary, originalArgs, nil
}
