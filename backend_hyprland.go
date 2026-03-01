//go:build hyprland

package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type hyprBackend struct{}

func GetBackend() Backend {
	return hyprBackend{}
}

type hyprMonitor struct {
	Name    string `json:"name"`
	Focused bool   `json:"focused"`
}

type hyprWindow struct {
	At   [2]int `json:"at"`
	Size [2]int `json:"size"`
}

func (hyprBackend) FocusedOutput() (string, error) {
	out, err := exec.Command("hyprctl", "monitors", "-j").Output()
	if err != nil {
		return "", err
	}
	var mons []hyprMonitor
	if err := json.Unmarshal(out, &mons); err != nil {
		return "", err
	}
	for _, m := range mons {
		if m.Focused {
			return m.Name, nil
		}
	}
	return "", fmt.Errorf("no focused monitor")
}

func (hyprBackend) ActiveWindowGeometry() (string, error) {
	out, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		return "", err
	}
	var w hyprWindow
	if err := json.Unmarshal(out, &w); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d,%d %dx%d",
		w.At[0], w.At[1], w.Size[0], w.Size[1]), nil
}
