package detector

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/frsfahd/dockergen/internal/catalog"
)

func detectNodeFramework(pkg packageJSON) string {
	deps := mergeMaps(pkg.Dependencies, pkg.DevDeps)
	keys := map[string]string{
		"next":          string(catalog.NextJS),
		"react-scripts": "cra",
		"vite":          string(catalog.Vite),
		"@nestjs/core":  "nestjs",
		"express":       string(catalog.ExpressJS),
		"koa":           "koa",
		"hapi":          string(catalog.Hapi),
		"fastify":       string(catalog.Fastify),
	}
	for dep, name := range keys {
		if _, ok := deps[dep]; ok {
			return name
		}
	}
	return string(catalog.NodeDefault)
}

func detectPackageManager(root string) string {
	if exists(filepath.Join(root, "pnpm-lock.yaml")) {
		return "pnpm"
	}
	if exists(filepath.Join(root, "yarn.lock")) {
		return "yarn"
	}
	if exists(filepath.Join(root, "package-lock.json")) {
		return "npm"
	}
	return "npm"
}

func nodeDefaultPort(framework string) int {
	switch framework {
	case string(catalog.NextJS):
		return 3000
	case "nestjs":
		return 3000
	case string(catalog.Vite):
		return 5173
	default:
		return 3000
	}
}

func scriptOrDefault(scripts map[string]string, key, manager string, fallback string) string {
	if scripts == nil {
		return fallback
	}
	if value, ok := scripts[key]; ok && strings.TrimSpace(value) != "" {
		return fmt.Sprintf("%s run %s", manager, key)
	}
	return fallback
}
