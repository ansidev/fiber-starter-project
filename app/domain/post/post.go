package post

import (
	"time"

	"github.com/ansidev/fiber-starter-project/domain/author"
)

type Post struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  int64
	Author    author.Author `json:"author"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (m Post) TableName() string {
	return "post"
}
