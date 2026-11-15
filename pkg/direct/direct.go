package direct

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gookit/ini/v2"
)

func ResovleCommand(originalExe string, targetCommand string, cfg *ini.Ini, originalArgs []string) (string, []string, []string, error) {
	originalExeDir := filepath.Dir(originalExe)
	targetBinary := strings.TrimSpace(cfg.String("binary", targetCommand))
	baseDir := strings.TrimSpace(cfg.String("baseDir", originalExeDir))
	extraPath := strings.TrimSpace(cfg.String("extraPath", ""))
	resolvedEnv := os.Environ()
	if extraPath != "" {
		newPath := os.Getenv("PATH") + string(os.PathListSeparator) + extraPath
		resolvedEnv = append(resolvedEnv, "PATH="+newPath)
	}
	if !filepath.IsAbs(baseDir) {
		baseDir = originalExe
	}
	if !filepath.IsAbs(targetBinary) {
		targetBinary = filepath.Join(baseDir, targetBinary)
	}
	resolvedBinary, err := exec.LookPath(targetBinary)
	if err != nil {
		return "", nil, nil, err
	}
	if resolvedBinary == originalExe {
		return "", nil, nil, fmt.Errorf("possible recursion detected calling: %s", resolvedBinary)
	}
	if strings.TrimSuffix(filepath.Base(resolvedBinary), filepath.Ext(resolvedBinary)) != targetCommand {
		return "", nil, nil, fmt.Errorf("invalid binary: %s", resolvedBinary)
	}
	return resolvedBinary, originalArgs, resolvedEnv, nil
}
