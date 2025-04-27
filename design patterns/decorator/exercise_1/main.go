package main

import "dudu.decorators/exercise_1/notifier"

// Esse poderia ser uma service que notifica um usuário conforme as suas permissões
func main() {
	user := User{
		UUID:  "MANGA_123",
		Name:  "Manguito",
		Phone: "1140028922",
		NotificationChannels: NotificationChannels{
			Slack: SlackNotificationChannel{
				IsEnabled: true,
				Channel:   "#Top_Fofocas",
			},
			Facebook: FacebookNotificationChannel{
				IsEnabled: true,
				UserID:    "facebook_manguito",
			},
			SMS: SMSNotificationChannel{
				IsEnabled: true,
			},
		},
	}

	var notifier_service notifier.NotifierService

	notifier_service = notifier.NewBaseNotifier()

	if user.NotificationChannels.IsSlackEnabled() {
		notifier_service = notifier.NewSlackNotifier(notifier.NewSlackNotifierInput{
			Notifier:   notifier_service,
			SlackToken: "token_seguro_do_slack",
			Channel:    user.NotificationChannels.Slack.Channel,
		})
	}

	if user.NotificationChannels.IsFacebookEnabled() {
		notifier_service = notifier.NewFacebookNotifier(notifier.NewFacebookNotifierInput{
			Notifier:          notifier_service,
			FacebookToken:     "token_seguro_do_facebook",
			DestinataryUserID: user.NotificationChannels.Facebook.UserID,
		})
	}

	if user.NotificationChannels.IsSMSEnabled() {
		notifier_service = notifier.NewSMSNotifier(notifier.NewSMSNotifierInput{
			Notifier:    notifier_service,
			Sender:      "1540028922",
			Destination: user.Phone,
		})
	}

	notifier_service.Notify("Você foi Du Mangado!")
}

// Domain
type User struct {
	UUID                 string
	Name                 string
	Phone                string
	NotificationChannels NotificationChannels
}

// VO
type NotificationChannels struct {
	Slack    SlackNotificationChannel
	Facebook FacebookNotificationChannel
	SMS      SMSNotificationChannel
}

func (n NotificationChannels) IsSlackEnabled() bool {
	return n.Slack.IsEnabled
}

func (n NotificationChannels) IsFacebookEnabled() bool {
	return n.Facebook.IsEnabled
}

func (n NotificationChannels) IsSMSEnabled() bool {
	return n.SMS.IsEnabled
}

type SlackNotificationChannel struct {
	IsEnabled bool
	Channel   string
}

type FacebookNotificationChannel struct {
	IsEnabled bool
	UserID    string
}

type SMSNotificationChannel struct {
	IsEnabled bool
}
