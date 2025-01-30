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

	var app string
	var args []string
	switch strings.TrimSpace(cfg.String("mode", "")) {
	case "wsl", "":
		app, args, err = wsl.ResovleCommand(targetBinary, cfg, os.Args[1:])
	case "direct":
		app, args, err = direct.ResovleCommand(binary, targetBinary, cfg, os.Args[1:])
	default:
		err = errors.New("Unsupported mode")
	}
	if err != nil {
		os.Exit(-1)
	}
	cmd := exec.Command(app, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	if e != nil {
		os.Exit(-1)
	}
	os.Exit(cmd.ProcessState.ExitCode())
}
