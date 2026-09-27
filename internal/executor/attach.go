package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func attachSession(sessionName string) error {
	fmt.Printf("Attaching to session %q...\n", sessionName)

	cmd := exec.Command("tmux", "attach", "-t", sessionName)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to attach to session %q: %w", sessionName, err)
	}

	return nil
}
