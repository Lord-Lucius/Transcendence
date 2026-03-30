package model

import (
	"time"
)

type Posts struct {
	ID string `json:"id"`
	UserID string `json:"userid"`
	RepostID string `json:"repostid"`
	ReplyID string `json:"replyid"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	DeletedAt time.Time `json:"deletedat"`
}
