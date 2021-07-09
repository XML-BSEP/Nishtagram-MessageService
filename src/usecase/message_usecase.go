package usecase

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"message-service/domain"
	"message-service/repository"
	"os"
	"strings"
)

type messageUsecsae struct {
	MessageRepository repository.MessageRepository
}


type MessageUsecase interface {
	GetMessages(ctx context.Context, receiver, sender string) ([]*domain.Message, error)
	GetUsers(ctx context.Context, userId string) ([]string, error)
	Create(ctx context.Context, message domain.Message) (*domain.Message, error)
	EncodeBase64(media string, messageId string, ctx context.Context) (string, error)
	DecodeBase64(media string, messageId string, ctx context.Context) (string, error)
	UpdateSeenStatus(ctx context.Context, messageId string, value bool) error
	IsAllowedToSee(ctx context.Context, messageId string) bool
}

func NewMessageUsecase(messageRepository repository.MessageRepository) MessageUsecase {
	return &messageUsecsae{
		MessageRepository: messageRepository,
	}
}

func (m *messageUsecsae) GetMessages(ctx context.Context, receiver, sender string) ([]*domain.Message, error) {
	return m.MessageRepository.GetMessages(ctx, receiver, sender)
}

func (m *messageUsecsae) GetUsers(ctx context.Context, userId string) ([]string, error) {
	return m.MessageRepository.GetUsers(ctx, userId)
}

func (m *messageUsecsae) Create(ctx context.Context, message domain.Message) (*domain.Message, error) {
	fmt.Println("Usao u create")
	return m.MessageRepository.Create(ctx, message)
}

func (m *messageUsecsae) DecodeBase64(media string, messageId string, ctx context.Context) (string, error) {

	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	path1 := "./assets/"
	err := os.Chdir(path1)
	fmt.Println(err)

	err = os.Chdir(messageId)
	fmt.Println(err)

	spliced := strings.Split(media, "/")
	var f *os.File
	if len(spliced) > 1 {
		f, _ = os.Open(spliced[1])
	} else {
		f, _ = os.Open(spliced[0])
	}


	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)


	encoded := base64.StdEncoding.EncodeToString(content)

	os.Chdir(workingDirectory)

	return "data:image/jpg;base64," + encoded, nil
}

func (m *messageUsecsae) EncodeBase64(media string, messageId string, ctx context.Context) (string, error) {

	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	path1 := "./assets/"
	err := os.Chdir(path1)
	if err != nil {
		return "", err
	}
	err = os.Mkdir(messageId, 0755)
	fmt.Println(err)

	err = os.Chdir(messageId)
	fmt.Println(err)

	s := strings.Split(media, ",")
	a := strings.Split(s[0], "/")
	format := strings.Split(a[1], ";")
	dec, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return "", err
	}
	uuid := uuid.NewString()
	f, err := os.Create(uuid + "." + format[0])
	if err != nil {
		return "", err
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return "", err
	}
	if err := f.Sync(); err != nil {
		return "", err
	}

	os.Chdir(workingDirectory)
	ret := messageId + "/" + uuid + "." + format[0]
	return ret, nil
}

func (m *messageUsecsae) UpdateSeenStatus(ctx context.Context, messageId string, value bool) error {
	return m.MessageRepository.UpdateSeenStatus(ctx, messageId, value)
}

func (m *messageUsecsae) IsAllowedToSee(ctx context.Context, messageId string) bool {
	return m.MessageRepository.IsAllowedToSee(ctx, messageId)
}

