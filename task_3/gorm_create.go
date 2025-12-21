package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:64;not null"`
	PostCount int    `gorm:"default:0"` // 冗余字段，供钩子维护
	Posts     []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// Post 文章
type Post struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"size:128;not null"`
	Content      string    `gorm:"type:text"`
	UserID       uint      `gorm:"index"`
	CommentCount int       `gorm:"default:0"`   // 冗余字段
	CommentState string    `gorm:"default:有评论"` // 供钩子维护
	Comments     []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}

// Comment 评论
type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	PostID uint   `gorm:"index;not null"`
	Author string `gorm:"size:64"`
	Text   string `gorm:"type:text"`
}

func GormCreate() {
	dsn := "root:passw0rd@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 自动建表（含外键）
	if err = db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Fatal(err)
	}
}
