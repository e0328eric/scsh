//go:build niri

package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type niriBackend struct{}

func GetBackend() Backend {
	return niriBackend{}
}

type niriFocused struct {
	Name string `json:"name"`
}

func (niriBackend) FocusedOutput() (string, error) {
	out, err := exec.Command("niri", "msg", "--json", "focused-output").Output()
	if err != nil {
		return "", err
	}
	var f niriFocused
	if err := json.Unmarshal(out, &f); err != nil {
		return "", err
	}
	return f.Name, nil
}

// niri does not reliably expose window geometry
// so we intentionally return error and let caller fallback to slurp.
func (niriBackend) ActiveWindowGeometry() (string, error) {
	return "", fmt.Errorf("window geometry unsupported in niri backend")
}
