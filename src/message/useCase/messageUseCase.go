package useCase

import (
	message "message"
	msrp "message/repository"
	"time"
)

type MessageUseCase interface {
	GetMessages(startAt time.Time, endAt time.Time) []message.Message
	SaveMessage(user string, text string)
}

type messageUsecase struct {
	repository             msrp.MessageRepository
	functionToRecoverImage func(string) []byte
}

func NewMessageUseCase(repo msrp.MessageRepository, f func(string) []byte) MessageUseCase {

	return messageUsecase{repository: repo, functionToRecoverImage: f}
}

func (useCase messageUsecase) GetMessages(startAt time.Time, endAt time.Time) []message.Message {
	messages := useCase.repository.GetMessages(startAt, endAt)

	for _, element := range messages {
		//TODO: CONVERT IT INTO GOROUTINES
		addImageToMessage(&element, useCase.functionToRecoverImage)
	}
	return messages

}

func (useCase messageUsecase) SaveMessage(user string, text string) {
	auxMessage := message.Message{}
	auxMessage.Name = user
	auxMessage.Date = time.Now()
	auxMessage.Text = text
	useCase.repository.SaveMessage(auxMessage)

}
func addImageToMessage(input *message.Message, functionToRecoverImage func(string) []byte) {

}
