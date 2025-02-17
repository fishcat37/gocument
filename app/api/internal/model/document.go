package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Document struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT" uri:"id" form:"id"`
	UserID    uint      `json:"user_id" gorm:"user_id"`
	Title     string    `json:"title"`
	Authority int       `json:"authority"` //1为公开，0为私密
	CreateAt  time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt  time.Time `json:"update_at" gorm:"autoUpdateTime"`
}

type DocumentContent struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	DocID   uint               `bson:"doc_id"`
	Content string             `bson:"content" json:"content"`
}

type WholeDocument struct {
	Document Document
	Content  DocumentContent
}
