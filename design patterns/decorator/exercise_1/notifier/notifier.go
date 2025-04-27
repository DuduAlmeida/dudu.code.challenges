package notifier

type NotifierService interface {
	Notify(message string)
}
