package repo

import (
	domain "MessagingServer/domain"
)

type IMessageRepo interface {
	AddMessage(userId string, data string) (err error)
	GetMessageByUserId(userId string) (record []*domain.Message, err error)
}
