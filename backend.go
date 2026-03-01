package main

type Backend interface {
	FocusedOutput() (string, error)
	ActiveWindowGeometry() (string, error)
}
