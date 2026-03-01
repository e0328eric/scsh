package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	var mode, out_path string
	var backend Backend
	var err error

	if len(os.Args) < 2 {
		fmt.Println("invalid argument")
		goto USAGE
	}

	mode = os.Args[1]
	backend = GetBackend()

	out_path, err = outPath()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	switch mode {
	case "output":
		monitor, err := backend.FocusedOutput()
		if err != nil {
			fmt.Println("fallback full screenshot")
			exec.Command("sh", "-c", "grim - | wl-copy").Run()
			return
		}
		exec.Command("grim", "-o", monitor, out_path).Run()

	case "window":
		win_geometry, err := backend.ActiveWindowGeometry()
		if err != nil {
			fmt.Println("fallback region selection")
			exec.Command("sh", "-c", "grim -g $(slurp) - | wl-copy").Run()
			return
		}
		exec.Command("grim", "-g", win_geometry, out_path).Run()
	case "region":
	default:
		fmt.Printf("invalid mode %s was given\n", mode)
		goto USAGE
	}

	os.Exit(0)

USAGE:
	fmt.Println("Usage: scsh <output|window|region>")
	os.Exit(1)
}

func outPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	pic := filepath.Join(home, "Pictures")

	if err := os.MkdirAll(pic, 0o755); err != nil {
		return "", err
	}

	ts := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("screenshot-%s.png", ts)

	return filepath.Join(pic, filename), nil
}
