package repo

type IMessageRepo interface {
	AddMessage(userId string, data string) (err error)
}
