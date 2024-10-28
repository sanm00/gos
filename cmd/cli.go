package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "创建新的 RESTful API 项目",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "项目名称",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "输出目录 (默认为当前目录)",
				Value:   ".",
			},
		},
		Action: func(c *cli.Context) error {
			projectName := c.String("name")
			outputDir := c.String("output")
			return generateProject(projectName, outputDir)
		},
	}
}

func generateProject(name, outputDir string) error {
	// 创建项目目录
	projectPath := filepath.Join(outputDir, name)
	err := os.MkdirAll(projectPath, 0755)
	if err != nil {
		return fmt.Errorf("创建项目目录失败: %w", err)
	}

	// 创建子目录
	dirs := []string{"cmd", "internal", "pkg", "internal/handlers", "internal/models", "internal/middleware", "internal/router", "internal/configs", "pkg/logger", "pkg/response", "pkg/errors", "scripts"}
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(projectPath, dir), 0755)
		if err != nil {
			return fmt.Errorf("创建目录 %s 失败: %w", dir, err)
		}
	}

	// 生成文件
	files := map[string]string{
		"main.go":                         "main.tmpl",
		"go.mod":                          "go.mod.tmpl",
		"config.yml.example":              "config.yml.example.tmpl",
		"README.md":                       "README.tmpl",
		"internal/configs/config.go":      "internal_config.tmpl",
		"internal/router/router.go":       "internal_router.tmpl",
		"internal/handlers/handler.go":    "internal_handler.tmpl",
		"internal/models/model.go":        "internal_model.tmpl",
		"internal/middleware/base.go":     "middleware_base.tmpl",
		"internal/middleware/logger.go":   "middleware_logger.tmpl",
		"internal/middleware/recovery.go": "middleware_recovery.tmpl",
		"internal/middleware/error.go":    "middleware_error.tmpl",
		"pkg/logger/logger.go":            "pkg_logger.tmpl",
		"pkg/response/response.go":        "pkg_response.tmpl",
		"pkg/errors/errors.go":            "pkg_errors.tmpl",
	}

	for file, tmpl := range files {
		err := createFileFromTemplate(filepath.Join(projectPath, file), tmpl, name)
		if err != nil {
			return fmt.Errorf("创建文件 %s 失败: %w", file, err)
		}
	}

	fmt.Printf("项目 %s 已成功创建在 %s 目录下!\n", name, projectPath)
	return nil
}

func createFileFromTemplate(path, tmplName, projectName string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	tmplPath := filepath.Join("templates", tmplName)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	data := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}

	return tmpl.Execute(f, data)
}
