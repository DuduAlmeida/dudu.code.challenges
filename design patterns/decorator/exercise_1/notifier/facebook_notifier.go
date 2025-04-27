package notifier

import "fmt"

type FacebookNotifier struct {
	Notifier          NotifierService
	FacebookToken     string
	DestinataryUserID string
}

type NewFacebookNotifierInput struct {
	Notifier          NotifierService
	FacebookToken     string
	DestinataryUserID string
}

func NewFacebookNotifier(input NewFacebookNotifierInput) NotifierService {
	return &FacebookNotifier{
		Notifier:          input.Notifier,
		FacebookToken:     input.FacebookToken,
		DestinataryUserID: input.DestinataryUserID,
	}
}

func (s *FacebookNotifier) Notify(message string) {
	s.Notifier.Notify(message)

	//Lógica para notificar pelo Facebook

	fmt.Println("[FacebookNotifier] Notifiy :", message)
}
