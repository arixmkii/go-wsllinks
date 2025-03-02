package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/arixmkii/go-wsllinks/pkg/direct"
	"github.com/arixmkii/go-wsllinks/pkg/wsl"
	"github.com/gookit/ini/v2"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		os.Exit(-1)
	}
	exeDir := filepath.Dir(exe)
	targetCommand := strings.TrimSuffix(filepath.Base(exe), filepath.Ext(exe))

	cfg := ini.NewWithOptions(func(opts *ini.Options) {
		opts.Readonly = true
		opts.ParseEnv = false
		opts.ParseVar = false
	})

	err = cfg.LoadExists(filepath.Join(exeDir, "wsllinks.ini"), filepath.Join(exeDir, targetCommand+".ini"))
	if err != nil {
		os.Exit(-1)
	}

	var runnerBinary string
	var runnerArgs []string
	switch strings.TrimSpace(cfg.String("mode", "")) {
	case "wsl", "":
		runnerBinary, runnerArgs, err = wsl.ResovleCommand(targetCommand, cfg, os.Args[1:])
	case "direct":
		runnerBinary, runnerArgs, err = direct.ResovleCommand(exe, targetCommand, cfg, os.Args[1:])
	default:
		err = errors.New("unsupported mode")
	}
	if err != nil {
		os.Exit(-1)
	}
	cmd := exec.Command(runnerBinary, runnerArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	if e != nil {
		os.Exit(-1)
	}
	os.Exit(cmd.ProcessState.ExitCode())
}
