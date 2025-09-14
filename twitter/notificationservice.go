package main 


type NotificationService struct {
	notifications map[string][]*Notification // key: userId, value: []*Notification
}

func NewNotificationService() *NotificationService {
	return &NotificationService{notifications: make(map[string][]*Notification)}
}

func (n *NotificationService) AddNotifcation(userId string, notificationType NotificationType, refId string) {
	notif := NewNotification(userId, notificationType, refId)
	n.notifications[userId] = append(n.notifications[userId], notif)
}

func (n *NotificationService) GetNotifications(userId string) []*Notification {
	return n.notifications[userId]
}

func (n *NotificationService) GetUnreadNotification(userId string)[]*Notification {
	var notifs []*Notification

	for _, notif := range n.notifications[userId] {
		if !notif.IsRead() {
			notifs = append(notifs, notif)
		}
	}
	return notifs
}

func (n *NotificationService) MarkAllRead(userId string)  {
	for _, notif := range n.notifications[userId] {
		notif.MarkRead()
	}
}



