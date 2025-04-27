package notifier

import "fmt"

type BaseNotifier struct {
}

func NewBaseNotifier() NotifierService {
	return &BaseNotifier{}
}

func (s *BaseNotifier) Notify(message string) {
	//Poderia ser um log no cloudwatch//New relic

	fmt.Println("[BaseNotifier] Notifiy :", message)
}
