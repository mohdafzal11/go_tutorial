package main

import "fmt"

type Notifier interface {
	Notify(to string, message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Notify(to string, message string) {
	fmt.Println("Sending email to", to , "with message:", message)
}

type SlackNotifier struct{}

func (s SlackNotifier) Notify(to string, message string) {
	fmt.Println("Sending message to", to , "with message:", message)
}

func main() {
  email := EmailNotifier{}
  slack := SlackNotifier{}

  email.Notify("afzal-mohd@gmail.com", "Hi Mohdafzal!")
  slack.Notify("afzal-mohd@gmail.com", "Hi Mohdafzal!")
}
