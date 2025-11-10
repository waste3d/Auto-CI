package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
		fmt.Printf("Результат будет сохранен в: %s\n", output)
		fmt.Println("\n(На этом шаге логика анализа и генерации еще не реализована)")
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
