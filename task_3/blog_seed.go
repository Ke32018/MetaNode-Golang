//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"math/rand"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 把题目 1 的模型复制一份（或者直接把公共模型抽出去，这里图方便就再写一次）
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:64;not null"`
	PostCount int    `gorm:"default:0"`
	Posts     []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Post struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"size:128;not null"`
	Content      string    `gorm:"type:text"`
	UserID       uint      `gorm:"index"`
	CommentCount int       `gorm:"default:0"`
	CommentState string    `gorm:"default:有评论"`
	Comments     []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}

type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	PostID uint   `gorm:"index;not null"`
	Author string `gorm:"size:64"`
	Text   string `gorm:"type:text"`
}

func BlogSeed() {
	dsn := "root:passw0rd@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// rand.Seed(time.Now().UnixNano())

	// 1. 创建 3 个用户
	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Carol"},
	}
	for i := range users {
		db.Create(&users[i])
	}

	// 2. 为每个用户创建 2~4 篇文章
	var posts []Post
	for _, u := range users {
		n := 2 + rand.Intn(3) // 2~4
		for i := 0; i < n; i++ {
			p := Post{
				Title:   fmt.Sprintf("%s 的文章-%d", u.Name, i+1),
				Content: fmt.Sprintf("这是 %s 的第 %d 篇内容……", u.Name, i+1),
				UserID:  u.ID,
			}
			db.Create(&p)
			posts = append(posts, p)
		}
	}

	// 3. 为每篇文章随机添加 0~5 条评论
	for _, p := range posts {
		cCnt := rand.Intn(6) // 0~5
		for j := 0; j < cCnt; j++ {
			c := Comment{
				PostID: p.ID,
				Author: fmt.Sprintf("读者-%d", rand.Intn(100)),
				Text:   fmt.Sprintf("沙发 %d 楼！", j+1),
			}
			db.Create(&c)
		}
	}

	// 4. 手动把其中一篇文章的评论数刷成最多（方便肉眼观察“评论最多”的查询）
	var maxPost Post
	db.First(&maxPost) // 任取一篇
	for i := 0; i < 12; i++ {
		db.Create(&Comment{
			PostID: maxPost.ID,
			Author: "水军",
			Text:   fmt.Sprintf("灌水 %d", i+1),
		})
	}

	log.Println("✅ 造数据完成，可运行题目 2/3 的代码验证效果！")
}
