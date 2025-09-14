package main

import (
	"fmt"
	"sort"
	"sync"
)

type TweetService struct {
	users  map[string]*User
	tweets map[string]*Tweet
	notificationService *NotificationService 
	mu     sync.RWMutex
}

func NewTweetService() *TweetService {
	return &TweetService{
		users: make(map[string]*User),
		tweets: make(map[string]*Tweet),
		notificationService: NewNotificationService(),
	}
}

func(t *TweetService) CreateUser(id, name, email string) *User {
	t.mu.RLock()
	user, exists := t.users[id] 
	t.mu.RUnlock()

	if exists {
		fmt.Println("User already exists")
		return user 
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	user = NewUser(id, name, email)
	t.users[id] = user 
	return user 
}



func(t *TweetService) Follow(follwerId, followeeId string) {
	t.mu.RLock()
	follower, ok1 := t.users[follwerId]
	followee, ok2 := t.users[followeeId]
	t.mu.RUnlock()

	if !ok1 {
		fmt.Println("Follower does not exist ", follwerId)
	}
	if !ok2 {
		fmt.Println("followee does not exist ", followeeId)
	}

	followee.AddFollower(follower)
	follower.AddFollowing(followee)
	// Add notification to followee 
	t.notificationService.AddNotifcation(followee.id, USER_FOLLOWED, follwerId)
}

func (t *TweetService) PostTweet(id string, authorId string, content string, mediaUrls []string) *Tweet {
	user, ok := t.users[authorId]
	if !ok {
		fmt.Println("User does not exist ", authorId)
		return nil 
	}
	tweet := NewTweet(id, authorId, content, mediaUrls)
	t.tweets[id] = tweet 
	user.tweetIds = append(user.tweetIds, id)

	// Notify all followers 
	for followerId := range user.GetFollowers() {
		t.notificationService.AddNotifcation(followerId, TWEET_POSTED, user.GetId())
	}

	return tweet 
}


func (t *TweetService) AddComment(id, userId, tweetId, content string) *Comment {
	_ , ok := t.users[userId]
	if !ok {
		fmt.Println("User does not exist ", userId)
		return nil 
	}

	tweet, ok := t.tweets[tweetId]
	if !ok {
		fmt.Println("Tweet does not exist ", tweetId)
		return nil
	}

	comment := NewComment(id, userId, tweetId, content)
	tweet.AddComment(comment) 
	// notify author
	t.notificationService.AddNotifcation(tweet.GetAuthorId(), TWEET_COMMENTED, userId)
	return comment
}

func (t *TweetService) LikeTweet(userID, tweetID string) {
	if tweet, exists := t.tweets[tweetID]; exists {
		tweet.AddLike(userID)
		// Notify author
		t.notificationService.AddNotifcation(tweet.GetAuthorId(), TWEET_LIKED, userID)
	}
}
 
func (t *TweetService) GetNewsFeed(userId string) []*FeedItem {
	user, exist := t.users[userId]
	if !exist {
		fmt.Println("User does not exist ", userId)
		return nil 
	}

	var tweets []string // stores tweetids 

	// Get all tweets of this user 
	tweets = append(tweets, user.GetTweetIds()... )


	// Collect tweets of users this user follows
	for followeeId := range user.GetFollowing() {
		followeeUser := t.users[followeeId]
		tweets = append(tweets, followeeUser.GetTweetIds()... )
	}

	// build feedItem
	var feeds []*FeedItem 
	for _, tweetId := range tweets  {
		tweet := t.tweets[tweetId]
		feeds = append(feeds, NewFeedItem(tweetId, tweet.authorId, tweet.content, len(tweet.usersLiked), tweet.usersLiked, tweet.comments, tweet.timestamp))
	}

	sort.Slice(feeds, func(i int, j int) bool {
		return feeds[i].timestamp.After(feeds[j].timestamp) 
	} )

	return feeds
}

func(t *TweetService) PrintFeed(feeds []*FeedItem) {
	for _, feed := range feeds {
		feed.Display()
	}
} 

func(t *TweetService) GetNotifications(userId string) []*Notification {
	_, ok := t.users[userId]
	if !ok {
		fmt.Println("user does not exist ", userId )
		return nil
	}
	return t.notificationService.GetNotifications(userId)
}

func(t *TweetService) GetUnreadNotification(userId string) []*Notification {
	_, ok := t.users[userId]
	if !ok {
		fmt.Println("user does not exist ", userId )
		return nil
	}
	return t.notificationService.GetUnreadNotification(userId)
}

func(t *TweetService) MarkAllRead(userId string) {
	_, ok := t.users[userId]
	if !ok {
		fmt.Println("user does not exist ", userId )
	}
	t.notificationService.MarkAllRead(userId)
}