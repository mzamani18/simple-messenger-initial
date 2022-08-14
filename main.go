package main

import (
	"errors"
)

type Bale interface {
	AddUser(username string, isBot bool) (int, error)
	AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error)
	SendMessage(userId, chatId int, text string) (int, error)
	SendLike(userId, messageId int) error
	GetNumberOfLikes(messageId int) (int, error)
	SetChatAdmin(chatId, userId int) error
	GetLastMessage(chatId int) (string, int, error)
	GetLastUserMessage(userId int) (string, int, error)
}

type User struct {
	username string
	IsBot    bool
	ID       int
}

type message struct {
	text     string
	ID       int
	likes    int
	likers   []int
	senderId int
}

// var AllMessages []message

type chat struct {
	chatName string
	isGroup  bool
	owner    int
	admins   []int
	ID       int
	messages []message
}

type BaleImpl struct {
	Users       []User
	Chats       []chat
	AllMessages []message
}

func NewBaleImpl() *BaleImpl {
	return &BaleImpl{}
}

func ValidUserName(username string) bool {
	if len(username) <= 3 {
		return false
	}
	var isDigigt bool = false
	var isStr bool = false
	for _, v := range username {
		// fmt.Println(v)
		if 48 <= v && v <= 57 {
			isDigigt = true
		} else if (97 <= v && v <= 122) || (65 <= v && v >= 90) {
			isStr = true
		}
	}

	// fmt.Println(isDigigt, isStr)
	return isStr && isDigigt
}

func (bi *BaleImpl) AddUser(username string, isBot bool) (int, error) {
	id := len(bi.Users) + 1
	if ValidUserName(username) {
		for _, v := range bi.Users {
			if v.username == username {
				return 0, errors.New("invalid username")
			}
		}
		var newUser = User{
			username: username,
			IsBot:    isBot,
			ID:       id,
		}
		bi.Users = append(bi.Users, newUser)
		return id, nil
	}
	return 0, errors.New("invalid username")
}

func (bi *BaleImpl) AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error) {
	id := len(bi.Chats) + 1
	if bi.Users[creator-1].IsBot {
		return 0, errors.New("could not create chat")
	}
	NewChat := chat{
		chatName: chatname,
		isGroup:  isGroup,
		owner:    creator,
		admins:   admins,
		ID:       id,
	}
	bi.Chats = append(bi.Chats, NewChat)
	return id, nil
}

func (bi *BaleImpl) SendMessage(userId, chatId int, text string) (int, error) {
	if bi.Chats[chatId-1].isGroup == false {
		for _, v := range bi.Chats[chatId-1].admins {
			if userId == v {
				id := len(bi.AllMessages) + 1
				newMessage := message{
					text:     text,
					ID:       id,
					senderId: userId,
					likes:    0,
				}
				bi.AllMessages = append(bi.AllMessages, newMessage)
				bi.Chats[chatId-1].messages = append(bi.Chats[chatId-1].messages, newMessage)
				return id, nil
			}
		}
		return 0, errors.New("user could not send message")
	}
	id := len(bi.AllMessages) + 1
	newMessage := message{
		text:     text,
		ID:       id,
		senderId: userId,
		likes:    0,
	}
	bi.AllMessages = append(bi.AllMessages, newMessage)
	bi.Chats[chatId-1].messages = append(bi.Chats[chatId-1].messages, newMessage)
	return id, nil
}

func (bi *BaleImpl) SendLike(userId, messageId int) error {
	if messageId <= 0 || messageId > len(bi.AllMessages) {
		return errors.New("message not found")
	}
	for _, v := range bi.AllMessages[messageId-1].likers {
		if v == userId {
			return errors.New("this user has liked this message before")
		}
	}
	bi.AllMessages[messageId-1].likes++
	bi.AllMessages[messageId-1].likers = append(bi.AllMessages[messageId-1].likers, userId)
	return nil
}

func (bi *BaleImpl) GetNumberOfLikes(messageId int) (int, error) {
	n := bi.AllMessages[messageId-1].likes
	return n, nil
}

func (bi *BaleImpl) SetChatAdmin(chatId, userId int) error {
	for _, v := range bi.Chats[chatId-1].admins {
		if v == userId {
			return errors.New("user is already admin")
		}
	}
	bi.Chats[chatId-1].admins = append(bi.Chats[chatId-1].admins, userId)
	return nil
}

func (bi *BaleImpl) GetLastMessage(chatId int) (string, int, error) {
	text := bi.Chats[chatId-1].messages[len(bi.Chats[chatId-1].messages)-1].text
	id := bi.Chats[chatId-1].messages[len(bi.Chats[chatId-1].messages)-1].ID
	return text, id, nil
}

func (bi *BaleImpl) GetLastUserMessage(userId int) (string, int, error) {
	text := ""
	id := 0
	for _, v := range bi.AllMessages {
		if v.senderId == userId {
			text = v.text
			id = v.ID
		}
	}
	return text, id, nil
}
