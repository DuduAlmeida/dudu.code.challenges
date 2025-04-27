package notifier

import "fmt"

type SlackNotifier struct {
	Notifier   NotifierService
	SlackToken string
	Channel    string
}

type NewSlackNotifierInput struct {
	Notifier   NotifierService
	SlackToken string
	Channel    string
}

func NewSlackNotifier(input NewSlackNotifierInput) NotifierService {
	return &SlackNotifier{
		Notifier:   input.Notifier,
		SlackToken: input.SlackToken,
		Channel:    input.Channel,
	}
}

func (s *SlackNotifier) Notify(message string) {
	s.Notifier.Notify(message)

	//Lógica para notificar por slack

	fmt.Println("[SlackNotifier] Notifiy :", message)
}
