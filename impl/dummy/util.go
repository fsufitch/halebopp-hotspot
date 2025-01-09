package dummy

import "fmt"

func newDummyError(msg string) error {
	if msg == "" {
		return nil
	}
	return fmt.Errorf("dummy error: %s", msg)
}
