package analyzer

import (
	"os"
	"path/filepath"
)

type GoDetector struct{}

func NewGoDetector() *GoDetector {
	return &GoDetector{}
}

func (d *GoDetector) Detect(directory string) (info *ProjectInfo, detected bool) {
	goModPath := filepath.Join(directory, "go.mod")

	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		return nil, false
	}

	info = &ProjectInfo{
		Language:       "Go",
		InstallCommand: "go mod download",
		BuildCommand:   "go build ./...",
		TestCommand:    "go test ./...",
		DockerImage:    "golang:latest",
	}

	return info, true
}
