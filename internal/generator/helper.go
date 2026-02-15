package generator

import (
	"fmt"
	"strings"
)

func jsonCommand(command string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "[\"sh\", \"-c\", \"echo 'No start command provided' && sleep 3600\"]"
	}
	quoted := make([]string, 0, len(parts))
	for _, part := range parts {
		quoted = append(quoted, fmt.Sprintf("\"%s\"", part))
	}
	return "[" + strings.Join(quoted, ", ") + "]"
}

func shellCommand(command string) string {
	return command
}
