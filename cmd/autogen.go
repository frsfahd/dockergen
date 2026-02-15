package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/frsfahd/dockergen/internal/detector"
	"github.com/frsfahd/dockergen/internal/generator"
	"github.com/spf13/cobra"
)

var autogenCmd = &cobra.Command{
	Use:   "autogen",
	Short: "Auto-detect project and generate a Dockerfile/Containerfile",
	RunE: func(cmd *cobra.Command, args []string) error {
		info, err := detector.Detect(".")
		if err != nil {
			return err
		}

		cfg, err := resolveAutogenConfig(cmd, info)
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
	autogenCmd.Flags().String("output", "", "Output file path (default: Dockerfile)")
	autogenCmd.Flags().String("app", "", "Application name")
	autogenCmd.Flags().String("workdir", "/app", "Work directory inside container")
	autogenCmd.Flags().Int("port", 0, "Application port to expose")
	autogenCmd.Flags().String("node-version", "", "Node.js version (e.g. 20)")
	autogenCmd.Flags().String("python-version", "", "Python version (e.g. 3.12)")
	autogenCmd.Flags().String("package-manager", "", "Node.js package manager (npm|yarn|pnpm)")
	autogenCmd.Flags().String("build", "", "Build command (optional)")
	autogenCmd.Flags().String("start", "", "Start command")
	autogenCmd.Flags().String("requirements", "", "Python requirements file name")
	autogenCmd.Flags().Bool("no-expose", false, "Do not include EXPOSE instruction")
	autogenCmd.Flags().Bool("alpine", true, "Use Alpine-based images when supported")
}

func resolveAutogenConfig(cmd *cobra.Command, info detector.ProjectInfo) (generator.Config, error) {
	output, _ := cmd.Flags().GetString("output")
	appName, _ := cmd.Flags().GetString("app")
	workdir, _ := cmd.Flags().GetString("workdir")
	port, _ := cmd.Flags().GetInt("port")
	build, _ := cmd.Flags().GetString("build")
	start, _ := cmd.Flags().GetString("start")
	noExpose, _ := cmd.Flags().GetBool("no-expose")
	alpine, _ := cmd.Flags().GetBool("alpine")
	requirements, _ := cmd.Flags().GetString("requirements")

	if appName == "" {
		appName = info.ProjectName
	}
	if port == 0 {
		port = info.DefaultPort
	}
	if start == "" {
		start = info.StartCommand
	}
	if build == "" {
		build = info.BuildCommand
	}
	if requirements == "" {
		requirements = info.Requirements
	}

	nodeVersion, _ := cmd.Flags().GetString("node-version")
	pythonVersion, _ := cmd.Flags().GetString("python-version")
	packageManager, _ := cmd.Flags().GetString("package-manager")

	if nodeVersion == "" {
		nodeVersion = info.NodeVersion
	}
	if pythonVersion == "" {
		pythonVersion = info.PythonVersion
	}
	if packageManager == "" {
		packageManager = info.PackageManager
	}

	cfg := generator.Config{
		Language:        string(info.Language),
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
