package main

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/gookit/ini/v2"
)

func main() {
	binary, err := os.Executable()
	if err != nil {
		os.Exit(-1)
	}
	binDir := filepath.Dir(binary)
	targetBinary := strings.TrimSuffix(filepath.Base(binary), filepath.Ext(binary))

	cfg := ini.NewWithOptions(func(opts *ini.Options) {
		opts.Readonly = true
		opts.ParseEnv = false
		opts.ParseVar = false
	})

	err = cfg.LoadExists(filepath.Join(binDir, "wsllinks.ini"), filepath.Join(binDir, targetBinary+".ini"))
	if err != nil {
		os.Exit(-1)
	}
	distro := cfg.String("distro", "")
	user := cfg.String("user", "")
	resolvedBinary := cfg.String("binary", targetBinary)
	if len(strings.TrimSpace(distro)) == 0 || len(strings.TrimSpace(user)) == 0 {
		os.Exit(-1)
	}
	if resolvedBinary != targetBinary {
		if !path.IsAbs(resolvedBinary) || path.Base(resolvedBinary) != targetBinary {
			os.Exit(-1)
		}
	}
	args := append([]string{"-d", distro, "-u", user, resolvedBinary}, os.Args[1:]...)
	cmd := exec.Command("wsl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	if e != nil {
		os.Exit(-1)
	}
	os.Exit(cmd.ProcessState.ExitCode())
}
