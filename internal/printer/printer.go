package printer

import (
	"github.com/znwng/tsd/internal/config"
)

func PrintConfig(cfg config.Config) {
	for _, session := range cfg.Sessions {
		printSession(session, 0)
	}
}

func printSession(session config.Session, depth int) {
	PrintIndented(depth, "Session: %s\n", session.Name)

	for _, window := range session.Windows {
		printWindow(window, depth+1)
	}
}

func printWindow(window config.Window, depth int) {
	PrintIndented(depth, "Window: %s\n", window.Name)

	for _, command := range window.Run {
		PrintIndented(depth+1, "Run: %s\n", command)
	}

	for _, command := range window.Send {
		PrintIndented(depth+1, "Send: %s\n", command)
	}
}
