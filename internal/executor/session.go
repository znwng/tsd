package executor

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/znwng/tsd/internal/config"
)

func sessionExists(session config.Session) bool {
	cmd := exec.Command("tmux", "list-sessions", "-F", "#S")

	output, err := cmd.Output()
	if err != nil {
		return false
	}

	for name := range strings.SplitSeq(strings.TrimSpace(string(output)), "\n") {
		if name == session.Name {
			return true
		}
	}

	return false
}

func executeSession(session config.Session) error {
	if sessionExists(session) {
		fmt.Printf("\nSession %q already exists\n", session.Name)
		fmt.Println("├── [s] Skip this session")
		fmt.Println("├── [q] Skip all remaining sessions")
		fmt.Println("├── [d] Delete the old session and create this one")
		fmt.Println("├── [r] Attach to existing session")
		fmt.Println("└── [a] Abort mission")
		fmt.Print("Answer: ")

		reader := bufio.NewReader(os.Stdin)

		for {
			answer, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}

			answer = strings.ToLower(strings.TrimSpace(answer))

			switch answer {
			case "s":
				fmt.Printf("Skipped session %q\n", session.Name)
				return nil

			case "q":
				return ErrSkipAll

			case "d":
				cmd := exec.Command("tmux", "kill-session", "-t", session.Name)

				if err := cmd.Run(); err != nil {
					return fmt.Errorf("failed to delete existing session %q: %w", session.Name, err)
				}

				fmt.Printf("Deleted session %q\n", session.Name)

			case "r":
				cmd := exec.Command("tmux", "attach-session", "-t", session.Name)

				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				fmt.Printf("Attaching to session %q...\n", session.Name)

				if err := cmd.Run(); err != nil {
					return fmt.Errorf(
						"failed to attach to existing session %q: %w",
						session.Name,
						err,
					)
				}

			case "a":
				return ErrAbort

			default:
				fmt.Println("Invalid answer. Please enter s, q, d, or a.")
				fmt.Print("Answer: ")
				continue
			}

			break
		}
	}

	cmd := exec.Command("tmux", "new-session", "-d", "-s", session.Name)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create session %q: %w", session.Name, err)
	}

	fmt.Printf("Created session %q\n", session.Name)

	for _, window := range session.Windows {
		if err := executeWindow(session.Name, window); err != nil {
			return err
		}
	}

	// Remove the automatically-created initial window.
	cmd = exec.Command(
		"sh",
		"-c",
		`tmux kill-window -t "$(tmux list-windows -F '#{window_index}' | head -n1)" && tmux move-window -r`,
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to remove initial window: %w", err)
	}

	// Change focus in session
	cmd = exec.Command("tmux", "select-window", "-t", session.Focus)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to change focus window: %w", err)
	}

	return nil
}
