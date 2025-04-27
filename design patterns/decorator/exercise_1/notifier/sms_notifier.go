package notifier

import "fmt"

type SMSNotifier struct {
	Notifier    NotifierService
	Sender      string
	Destination string
}

type NewSMSNotifierInput struct {
	Notifier    NotifierService
	Sender      string
	Destination string
}

func NewSMSNotifier(input NewSMSNotifierInput) NotifierService {
	return &SMSNotifier{
		Notifier:    input.Notifier,
		Sender:      input.Sender,
		Destination: input.Destination,
	}
}

func (s *SMSNotifier) Notify(message string) {
	s.Notifier.Notify(message)

	//Lógica para notificar por SMS

	fmt.Println("[SMSNotifier] Notifiy :", message)
}
