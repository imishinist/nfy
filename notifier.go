package nfy

// Notifier has Notify method
type Notifier interface {
	Notify(string, string) error
}
