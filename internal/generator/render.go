package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/frsfahd/dockergen/internal/catalog"
	placeholder "github.com/frsfahd/dockergen/internal/templates"
)

// GenerateDockerfile renders a Dockerfile based on the given config.
func GenerateDockerfile(cfg Config) (string, error) {
	tpl, err := templateFor(cfg)
	if err != nil {
		return "", err
	}

	data := templateData(cfg)
	parsed, err := template.New("dockerfile").Funcs(template.FuncMap{
		"shell": shellCommand,
	}).Parse(tpl)
	if err != nil {
		return "", fmt.Errorf("parse template: %w", err)
	}

	var out bytes.Buffer
	if err := parsed.Execute(&out, data); err != nil {
		return "", fmt.Errorf("render template: %w", err)
	}

	return strings.TrimSpace(out.String()) + "\n", nil
}

type templatePayload struct {
	Config
	BaseImageBuilder string
	BaseImageRuntime string
	InstallCommand   string
	CopyDepsCommand  string
	StartJSON        string
}

func templateData(cfg Config) templatePayload {
	payload := templatePayload{Config: cfg}

	switch cfg.Language {
	case string(catalog.Node):
		payload.BaseImageBuilder = nodeImage(cfg.NodeVersion, cfg.UseAlpineImages)
		payload.BaseImageRuntime = nodeImage(cfg.NodeVersion, cfg.UseAlpineImages)
		payload.InstallCommand = nodeInstallCommand(cfg.PackageManager)
		payload.CopyDepsCommand = nodeCopyDepsCommand(cfg.PackageManager)
		payload.StartJSON = jsonCommand(cfg.StartCommand)
	case string(catalog.Python):
		payload.BaseImageBuilder = pythonImage(cfg.PythonVersion)
		payload.BaseImageRuntime = pythonImage(cfg.PythonVersion)
		payload.StartJSON = jsonCommand(cfg.StartCommand)
	}

	return payload
}

func templateFor(cfg Config) (string, error) {
	switch cfg.Language {
	case string(catalog.Node):
		return placeholder.NodeTemplate, nil
	case string(catalog.Python):
		return placeholder.PythonTemplate, nil
	default:
		return "", fmt.Errorf("unsupported language: %s", cfg.Language)
	}
}
