package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)


type NotificationType int 

const (
	TWEET_POSTED NotificationType = iota 
	TWEET_LIKED
	TWEET_COMMENTED
	USER_FOLLOWED
)

type Notification struct {
	id string // uuid of the notification
	userId           string
	notificationType NotificationType
	referenceId string // it could be a tweetid or userid or commentId depending on notification type 
	time             time.Time
	isRead bool 
}

func NewNotification(userId string, notificationType NotificationType, referenceId string) *Notification {
	return &Notification{
		id : uuid.New().String(),
		userId: userId,
		notificationType: notificationType,
		referenceId: referenceId,
		time: time.Now(), 
		isRead: false,
	}
}

func(n *Notification) IsRead() bool{
	return n.isRead
}

func(n *Notification) MarkRead() {
	n.isRead = true 
}

func(n *Notification) Display() {
	notif := fmt.Sprintf("For user: %s, ", n.userId)
	switch n.notificationType {
	case USER_FOLLOWED:
		notif += fmt.Sprintf("user %s followed you", n.referenceId)
	case TWEET_COMMENTED:
		notif += fmt.Sprintf("user %s commented on you tweet", n.referenceId)
	case TWEET_LIKED:
		notif += fmt.Sprintf("user %s liked your tweet", n.referenceId)
	case TWEET_POSTED:
		notif += fmt.Sprintf("user %s added a new tweet", n.referenceId)
	default:
		return 
	}
	fmt.Println(notif)
}