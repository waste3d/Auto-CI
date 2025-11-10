package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/waste3d/Auto-CI/internal/analyzer"
	"github.com/waste3d/Auto-CI/internal/gitter"
)

var (
	repoURL string
	output  string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Генерирует CI/CD файл на основе анализа Git-репозитория",
	Long: `Эта команда клонирует указанный Git-репозиторий, анализирует его
	технологический стек и создает готовый к использованию CI/CD файл.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- Auto-CI ---")
		fmt.Printf("Получен URL репозитория: %s\n", repoURL)

		tempDir, err := gitter.CloneToTemp(repoURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при клонировании репозитория: %s\n", err)
			os.Exit(1)
		}
		defer os.RemoveAll(tempDir)

		fmt.Printf("Результат будет сохранен в: %s\n", output)

		projectInfo, err := analyzer.Analyze(tempDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при анализе проекта: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("\n--- Результаты анализа ---")
		fmt.Printf("Язык: %s\n", projectInfo.Language)
		fmt.Printf("Команда для установки зависимостей: %s\n", projectInfo.InstallCommand)
		fmt.Printf("Команда для сборки: %s\n", projectInfo.BuildCommand)
		fmt.Printf("Команда для запуска тестов: %s\n", projectInfo.TestCommand)
		fmt.Printf("Используемый Docker образ: %s\n", projectInfo.DockerImage)
		fmt.Println("\n(На следующем шаге эти данные будут использованы для генерации YAML-файла)")
	},
}

var rootCmd = &cobra.Command{
	Use:   "auto-ci",
	Short: "Auto-CI - это инструмент для автоматической генерации CI/CD пайплайнов.",
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&repoURL, "repo-url", "r", "", "URL-адрес Git-репозитория (обязательный)")
	generateCmd.MarkFlagRequired("repo-url")
	generateCmd.Flags().StringVarP(&output, "output", "o", ".gitlab-ci.yml", "Путь к выходному файлу CI/CD конфигурации")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при выполнении команды: '%s'", err)
		os.Exit(1)
	}
}
