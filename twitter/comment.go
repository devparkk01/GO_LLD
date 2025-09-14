package main

import "time"

type Comment struct {
	id        string // commentid 
	tweetId string 
	userId   string // user who added that comment
	content   string
	timestamp time.Time
}

func NewComment(id, userId, tweetId, content string) *Comment {
	return &Comment{
		id: id, 
		userId: userId,
		tweetId: tweetId,
		content: content,
		timestamp: time.Now(),
	}
}

func (c *Comment) GetContent() string {
	return c.content
}

func(c *Comment) GetTimestamp() time.Time {
	return c.timestamp
}

func (c *Comment) GetUserId() string {
	return c.userId 
}
