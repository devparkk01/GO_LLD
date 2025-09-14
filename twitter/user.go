package main 

type User struct {
	id string 
	name string 
	email string 
	followers map[string]struct{} // followers of this user 
	following map[string]struct{}  // users followed by this user 
	tweetIds []string // stores tweetids of tweets this user made  
}

func NewUser(id, name, email string) *User{
	return &User{
		id: id,
		name: name, 
		email: email,
		followers: make(map[string]struct{}),
		following: make(map[string]struct{}),
	}
}

func (u *User) AddFollower(follower *User) {
	u.followers[follower.id] = struct{}{} // empty struct 
}

func (u *User) AddFollowing(followee *User) {
	u.following[followee.id] = struct{}{}
}

func(u *User) GetTweetIds() []string{
	return u.tweetIds 
}

func(u *User) GetFollowing() map[string]struct{} {
	return u.following
}

func(u *User) GetFollowers() map[string]struct{} {
	return u.followers
}

func(u *User) GetId() string {
	return u.id
}

func(u *User) GetName() string {
	return u.name 
}

func(u *User) GetEmail() string {
	return u.email 
}

