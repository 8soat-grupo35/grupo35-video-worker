package repository

type SNS interface {
	SendMessage(message interface{}) error
}
