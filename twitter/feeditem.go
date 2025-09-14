package main

import (
	"fmt"
	"time"
)


type FeedItem struct {
	tweetId string 
	authorId string 
	content string 
	likesCount int 
	usersLiked map[string]struct{}
	comments []*Comment
	timestamp time.Time 
}

func NewFeedItem(tweedId string, authorId string, content string, likesCount int, usersLiked map[string]struct{}, comments []*Comment, timestamp time.Time) *FeedItem {
	return &FeedItem{
		tweetId: tweedId,
		authorId: authorId,
		content: content,
		likesCount: likesCount,
		usersLiked: usersLiked,
		comments: comments,
		timestamp: timestamp,
	}
}

// Display FeedItem in a clear formatting
func(f *FeedItem) Display() {
	usersLiked := ""
	comments := ""
	for userId := range f.usersLiked {
		usersLiked += fmt.Sprintf("%s ", userId)
	}
	for _, comment := range f.comments {
		comments += fmt.Sprintf("Added by: %s, content: %s, time: %s -- ", comment.GetUserId(), comment.GetContent(), comment.GetTimestamp().String())
	}

	feed := fmt.Sprintf("TweetID: %s, authorID: %s, time: %s\n", f.tweetId, f.authorId, f.timestamp.String())
	feed += fmt.Sprintf("content: %s, Total Likes: %d, usersLiked: [ %s ], Comments: [ %s ]\n", f.content,  f.likesCount, usersLiked, comments)
	fmt.Println(feed)
}