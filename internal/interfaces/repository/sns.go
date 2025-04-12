package repository

//go:generate mockgen -source=sns.go -destination=mock/sns.go
type SNS interface {
	SendMessage(message interface{}) error
}
