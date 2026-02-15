package generator

import "fmt"

func nodeImage(version string, alpine bool) string {
	if version == "" {
		version = "20"
	}
	if alpine {
		return fmt.Sprintf("node:%s-alpine", version)
	}
	return fmt.Sprintf("node:%s", version)
}

func nodeInstallCommand(manager string) string {
	switch manager {
	case "yarn":
		return "yarn install --frozen-lockfile"
	case "pnpm":
		return "pnpm install --frozen-lockfile"
	default:
		return "npm ci"
	}
}

func nodeCopyDepsCommand(manager string) string {
	switch manager {
	case "yarn":
		return "COPY --from=deps /app/node_modules ./node_modules"
	case "pnpm":
		return "COPY --from=deps /app/node_modules ./node_modules"
	default:
		return "COPY --from=deps /app/node_modules ./node_modules"
	}
}
