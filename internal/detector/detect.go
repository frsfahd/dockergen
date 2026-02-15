package detector

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/frsfahd/dockergen/internal/catalog"
)

type packageJSON struct {
	Name         string            `json:"name"`
	Scripts      map[string]string `json:"scripts"`
	Dependencies map[string]string `json:"dependencies"`
	DevDeps      map[string]string `json:"devDependencies"`
}

// Detect inspects the current directory to determine project properties.
func Detect(root string) (ProjectInfo, error) {
	root, err := filepath.Abs(root)
	if err != nil {
		return ProjectInfo{}, fmt.Errorf("resolve root: %w", err)
	}

	if exists(filepath.Join(root, "package.json")) {
		return detectNode(root)
	}

	if exists(filepath.Join(root, "requirements.txt")) || exists(filepath.Join(root, "pyproject.toml")) || exists(filepath.Join(root, "Pipfile")) {
		return detectPython(root)
	}

	return ProjectInfo{}, errors.New("unable to detect project type (no package.json, requirements.txt, or pyproject.toml found)")
}

func detectNode(root string) (ProjectInfo, error) {
	pkgPath := filepath.Join(root, "package.json")
	data, err := os.ReadFile(pkgPath)
	if err != nil {
		return ProjectInfo{}, fmt.Errorf("read package.json: %w", err)
	}

	var pkg packageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return ProjectInfo{}, fmt.Errorf("parse package.json: %w", err)
	}

	framework := detectNodeFramework(pkg)
	manager := detectPackageManager(root)
	start := scriptOrDefault(pkg.Scripts, "start", manager, fmt.Sprintf("%s run start", manager))
	build := scriptOrDefault(pkg.Scripts, "build", manager, "")
	version := readVersionFile(root, []string{".nvmrc", ".node-version"})

	return ProjectInfo{
		Language:       string(catalog.Node),
		Framework:      framework,
		ProjectName:    packageNameOrDir(pkg.Name, root),
		PackageManager: manager,
		NodeVersion:    version,
		StartCommand:   start,
		BuildCommand:   build,
		DefaultPort:    nodeDefaultPort(framework),
	}, nil
}

func detectPython(root string) (ProjectInfo, error) {
	requirements := pythonRequirementsFile(root)
	framework := detectPythonFramework(root, requirements)
	version := readVersionFile(root, []string{".python-version"})

	port := pythonDefaultPort(framework)
	start := pythonStartCommand(framework, strconv.Itoa(port))

	return ProjectInfo{
		Language:      string(catalog.Python),
		Framework:     framework,
		ProjectName:   filepath.Base(root),
		PythonVersion: version,
		StartCommand:  start,
		DefaultPort:   port,
		Requirements:  requirements,
	}, nil
}
