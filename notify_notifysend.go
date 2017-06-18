// +build linux freebsd

package nfy

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Notify displays a libnotify notification.
func Notify(title, message string) error {
	_, err := exec.LookPath("notify-send")
	if err != nil {
		return errors.New("Install 'notify-send' and try again")
	}

	cmd := exec.Command("notify-send", title, message)
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("Banner: %s", err)
	}

	return nil
}
