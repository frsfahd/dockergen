package generator

import "fmt"

func pythonImage(version string) string {
	if version == "" {
		version = "3.12"
	}
	return fmt.Sprintf("python:%s-slim", version)
}
