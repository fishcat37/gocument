package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Document struct {
	ID       uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID   uint      `json:"user_id"`
	Title    string    `json:"title"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time `json:"update_at" gorm:"autoUpdateTime"`
}

type DocumentContent struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	DocID   uint               `bson:"doc_id"`
	Content string             `bson:"content"`
}

type WholeDocument struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Title    string    `json:"title"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Content  string    `json:"content"`
}
