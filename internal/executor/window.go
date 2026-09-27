package executor

import (
	"fmt"
	"os/exec"

	"github.com/znwng/tsd/internal/config"
	"github.com/znwng/tsd/internal/printer"
)

func executeWindow(sessionName string, window config.Window) error {
	cmd := exec.Command("tmux", "new-window", "-t", sessionName, "-n", window.Name)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create window %q in session %q: %w", window.Name, sessionName, err)
	}

	printer.PrintIndented(1, "Created window: %s\n", window.Name)

	for _, command := range window.Run {
		if err := runInWindow(sessionName, window.Name, command); err != nil {
			return err
		}
	}

	for _, command := range window.Send {
		if err := sendInWindow(sessionName, window.Name, command); err != nil {
			return err
		}
	}

	return nil
}

func sendInWindow(sessionName string, windowName string, command string) error {
	target := sessionName + ":" + windowName

	cmd := exec.Command("tmux", "send-keys", "-t", target, command)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command %q in window %q of session %q: %w", command, windowName, sessionName, err)
	}

	printer.PrintIndented(2, "Sent command: %s\n", command)

	return nil
}

func runInWindow(sessionName string, windowName string, command string) error {
	target := sessionName + ":" + windowName

	cmd := exec.Command("tmux", "send-keys", "-t", target, command, "C-m")

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command %q in window %q of session %q: %w", command, windowName, sessionName, err)
	}

	printer.PrintIndented(2, "Executed command: %s\n", command)

	return nil
}
