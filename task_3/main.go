package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PostWithCnt struct {
	Post
	Cnt int64 `gorm:"column:c_cnt"`
}

func main() {
	dsn := "root:passw0rd@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 1. 查询某用户发布的所有文章及其评论（预加载）
	userID := uint(1)
	var user User
	if err := db.Preload("Posts.Comments").First(&user, userID).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %s\n", user.Name)
	for _, p := range user.Posts {
		fmt.Printf("  Post: %s\n", p.Title)
		for _, c := range p.Comments {
			fmt.Printf("    Comment by %s: %s\n", c.Author, c.Text)
		}
	}

	// 2. 查询评论数量最多的文章
	var top PostWithCnt
	db.Raw(`
SELECT posts.*, COUNT(comments.id) AS c_cnt
FROM posts
LEFT JOIN comments ON comments.post_id = posts.id
GROUP BY posts.id
ORDER BY c_cnt DESC
LIMIT 1`).Scan(&top)
	fmt.Printf("评论最多的文章：%s (共 %d 条评论)\n", top.Post.Title, top.Cnt)
}
