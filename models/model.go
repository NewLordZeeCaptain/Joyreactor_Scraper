package models

import "fmt"

// Struct Used to define User
type User struct {
	Id         int64
	Name       string
	UserRating float32
}

type Post struct {
	Id        int64
	Link      string
	Rating    int32
	CreatorId int64
	// Creator   *User
}

func (u *User) String() string {
	return fmt.Sprintf("User<%d %s %f", u.Id, u.Name, u.UserRating)
}

func (p *Post) String() string {
	// return fmt.Sprintf("User<%d %s %d %d %s", p.Id, p.Link, p.Rating, p.CreatorId, p.Creator)
	return fmt.Sprintf("User<%d %s %d %d", p.Id, p.Link, p.Rating, p.CreatorId)
}
