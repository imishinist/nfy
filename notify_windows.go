package nfy

import (
	"github.com/go-toast/toast"
)

// Notify displays a Windows 10 Toast Notification.
func Notify(title, message string) error {
	notification := toast.Notification{
		AppID:   "ntify",
		Title:   title,
		Message: message,
		Icon:    "",
		Actions: nil}

	return notification.Push()
}
