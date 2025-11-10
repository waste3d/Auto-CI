package analyzer

import (
	"fmt"
)

type ProjectInfo struct {
	Language       string
	BuildCommand   string
	TestCommand    string
	InstallCommand string
	DockerImage    string
}

type Detector interface {
	Detect(directory string) (info *ProjectInfo, detected bool)
}

func Analyze(dir string) (*ProjectInfo, error) {
	detectors := []Detector{
		NewGoDetector(),
	}

	for _, detector := range detectors {
		if info, detected := detector.Detect(dir); detected {
			fmt.Printf("✅ Стек успешно определен: %s\n", info.Language)
			return info, nil
		}
	}
	return nil, fmt.Errorf("не удалось определить технологический стек проекта. Попробуйте указать язык вручную.")
}
