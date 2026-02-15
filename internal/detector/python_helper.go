package detector

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/frsfahd/dockergen/internal/catalog"
)

func pythonRequirementsFile(root string) string {
	if exists(filepath.Join(root, "requirements.txt")) {
		return "requirements.txt"
	}
	if exists(filepath.Join(root, "pyproject.toml")) {
		return "pyproject.toml"
	}
	if exists(filepath.Join(root, "Pipfile")) {
		return "Pipfile"
	}
	return "requirements.txt"
}

func detectPythonFramework(root, requirements string) string {
	paths := []string{requirements, "pyproject.toml", "Pipfile"}
	content := readAny(root, paths)
	content = strings.ToLower(content)

	if strings.Contains(content, "django") {
		return string(catalog.Django)
	}
	if strings.Contains(content, "fastapi") {
		return string(catalog.FastAPI)
	}
	if strings.Contains(content, "flask") {
		return string(catalog.Flask)
	}
	if strings.Contains(content, "uvicorn") {
		return string(catalog.FastAPI)
	}
	return string(catalog.PythonDefault)
}

func pythonDefaultPort(framework string) int {
	switch framework {
	case string(catalog.Flask):
		return 5000
	default:
		return 8000
	}
}

func pythonStartCommand(framework string, port string) string {
	switch framework {
	case string(catalog.Django):
		return fmt.Sprintf("python manage.py runserver 0.0.0.0:%s", port)
	case string(catalog.Flask):
		return fmt.Sprintf("flask --app app run --host 0.0.0.0 --port %s", port)
	case string(catalog.FastAPI):
		return fmt.Sprintf("uvicorn app:app --host 0.0.0.0 --port %s", port)
	default:
		return "python app.py"
	}
}
