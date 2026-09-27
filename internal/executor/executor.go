package executor

import (
	"errors"
	"fmt"

	"github.com/znwng/tsd/internal/config"
)

var (
	ErrSkipAll = errors.New("skip all remaining sessions")
	ErrAbort   = errors.New("execution aborted")
)

func ExecuteConfig(cfg config.Config, attach string) error {
	for _, session := range cfg.Sessions {
		err := executeSession(session)

		if errors.Is(err, ErrSkipAll) {
			fmt.Println("Skipping all remaining sessions")
			break
		}

		if errors.Is(err, ErrAbort) {
			fmt.Println("Execution aborted")
			return nil
		}

		if err != nil {
			return fmt.Errorf("failed to execute session %q: %w", session.Name, err)
		}
	}

	if attach != "" {
		if err := attachSession(attach); err != nil {
			return err
		}
	}

	return nil
}
