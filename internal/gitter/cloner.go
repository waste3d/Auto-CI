package gitter

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func CloneToTemp(url string) (string, error) {
	tempDir, err := os.MkdirTemp("", "auto-ci-clone-*")
	if err != nil {
		return "", fmt.Errorf("не удалось создать временную директорию: %w", err)
	}

	fmt.Printf("Клонирование репозитория в %s...\n", tempDir)

	_, err = git.PlainClone(tempDir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		os.RemoveAll(tempDir)
		return "", fmt.Errorf("не удалось клонировать репозиторий: %w", err)
	}

	fmt.Println("Репозиторий успешно склонирован.")
	return tempDir, nil
}
