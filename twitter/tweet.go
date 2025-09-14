package main

import "time"

type Tweet struct {
	id        string
	authorId    string 
	content   string
	mediaUrls []string
	timestamp time.Time
	usersLiked map[string]struct{} // stores userId of those users who has liked this tweet 
	comments []*Comment 
}

func NewTweet(id string, authorId string, content string, mediaUrls []string) *Tweet{
	return &Tweet{
		id: id, 
		authorId: authorId,
		content: content,
		mediaUrls: mediaUrls,
		timestamp: time.Now(),
		usersLiked: make(map[string]struct{}),
	}
}

func (t *Tweet) AddComment(comment *Comment) {
	t.comments = append(t.comments, comment )
}

func (t *Tweet) AddLike(userId string) {
	t.usersLiked[userId] = struct{}{}
}

func (t *Tweet) GetId() string {
	return t.id 
}
func(t *Tweet) GetAuthorId() string {
	return t.authorId
}

func (t *Tweet) GetContent() string{
	return t.content 
}

func (t *Tweet) GetLikeCount() int {
	return len(t.usersLiked)
}