package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/frsfahd/dockergen/internal/catalog"
	"github.com/frsfahd/dockergen/internal/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Interactively generate a Dockerfile/Containerfile",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := resolveConfig(cmd)
		if err != nil {
			return err
		}

		content, err := generator.GenerateDockerfile(cfg)
		if err != nil {
			return err
		}

		output := cfg.OutputFile
		if output == "" {
			output = "Dockerfile"
		}

		if err := os.WriteFile(output, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write output: %w", err)
		}

		var absPath, _ = filepath.Abs(output)

		fmt.Printf("✅ Generated %s in %s\n ", filepath.Base(output), absPath)
		return nil
	},
}

func init() {
	generateCmd.Flags().String("lang", "", "Runtime language (node|python)")
	generateCmd.Flags().String("output", "", "Output file path (default: Dockerfile)")
	generateCmd.Flags().String("app", "", "Application name")
	generateCmd.Flags().String("workdir", "/app", "Work directory inside container")
	generateCmd.Flags().Int("port", 0, "Application port to expose")
	generateCmd.Flags().String("node-version", "", "Node.js version (e.g. 20)")
	generateCmd.Flags().String("python-version", "", "Python version (e.g. 3.12)")
	generateCmd.Flags().String("package-manager", "", "Node.js package manager (npm|yarn|pnpm)")
	generateCmd.Flags().String("build", "", "Build command (optional)")
	generateCmd.Flags().String("start", "", "Start command")
	generateCmd.Flags().String("requirements", "requirements.txt", "Python requirements file name")
	generateCmd.Flags().Bool("no-expose", false, "Do not include EXPOSE instruction")
	generateCmd.Flags().Bool("alpine", true, "Use Alpine-based images when supported")
}

func resolveConfig(cmd *cobra.Command) (generator.Config, error) {
	cfg := generator.Config{}

	lang, _ := cmd.Flags().GetString("lang")
	output, _ := cmd.Flags().GetString("output")
	appName, _ := cmd.Flags().GetString("app")
	workdir, _ := cmd.Flags().GetString("workdir")
	port, _ := cmd.Flags().GetInt("port")
	build, _ := cmd.Flags().GetString("build")
	start, _ := cmd.Flags().GetString("start")
	requirements, _ := cmd.Flags().GetString("requirements")
	noExpose, _ := cmd.Flags().GetBool("no-expose")
	alpine, _ := cmd.Flags().GetBool("alpine")

	reader := bufio.NewReader(os.Stdin)

	if lang == "" {
		lang = promptChoice(reader, "Language", []string{string(catalog.Node), string(catalog.Python)}, string(catalog.Node))
	}
	lang = strings.ToLower(strings.TrimSpace(lang))

	if appName == "" {
		appName = prompt(reader, "App name", "my-app")
	}

	if workdir == "" {
		workdir = "/app"
	}

	if noExpose == false && port == 0 {
		port = promptInt(reader, "Expose port", 3000)
	}

	var nodeVersion string
	var pythonVersion string
	var packageManager string

	switch lang {
	case string(catalog.Node):
		nodeVersion, _ = cmd.Flags().GetString("node-version")
		packageManager, _ = cmd.Flags().GetString("package-manager")
		if nodeVersion == "" {
			nodeVersion = prompt(reader, "Node version", "20")
		}
		if packageManager == "" {
			packageManager = promptChoice(reader, "Package manager", []string{"npm", "yarn", "pnpm"}, "npm")
		}
	case string(catalog.Python):
		pythonVersion, _ = cmd.Flags().GetString("python-version")
		if pythonVersion == "" {
			pythonVersion = prompt(reader, "Python version", "3.12")
		}
	default:
		return cfg, fmt.Errorf("unsupported language: %s", lang)
	}

	if start == "" {
		start = prompt(reader, "Start command", defaultStartCommand(lang, packageManager))
	}

	cfg = generator.Config{
		Language:        lang,
		OutputFile:      output,
		AppName:         appName,
		WorkDir:         workdir,
		Port:            port,
		NodeVersion:     nodeVersion,
		PythonVersion:   pythonVersion,
		PackageManager:  packageManager,
		BuildCommand:    build,
		StartCommand:    start,
		Requirements:    requirements,
		ExposePort:      !noExpose,
		UseAlpineImages: alpine,
	}

	return cfg, nil
}

func prompt(reader *bufio.Reader, label, fallback string) string {
	fmt.Printf("%s [%s]: ", label, fallback)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback
	}
	return value
}

func promptChoice(reader *bufio.Reader, label string, options []string, fallback string) string {
	fmt.Printf("%s (%s) [%s]: ", label, strings.Join(options, "/"), fallback)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(strings.ToLower(value))
	if value == "" {
		return fallback
	}
	for _, opt := range options {
		if value == opt {
			return value
		}
	}
	return fallback
}

func promptInt(reader *bufio.Reader, label string, fallback int) int {
	fmt.Printf("%s [%d]: ", label, fallback)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func defaultStartCommand(lang string, packageManager string) string {
	switch lang {
	case string(catalog.Node):
		switch packageManager {
		case "npm":
			return "npm start"
		case "yarn":
			return "yarn start"
		case "pnpm":
			return "pnpm start"
		default:
			return "npm run start" // Default for Node.js
		}
	case string(catalog.Python):
		return "python app.py"
	default:
		return ""
	}
}
