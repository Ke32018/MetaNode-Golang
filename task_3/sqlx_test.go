package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/* ---------- 结构体定义 ---------- */
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

/* ---------- 工具：获取 sqlx.DB ---------- */
func mustDB() *sqlx.DB {
	dsn := "root:passw0rd@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("connect DB failed: %v", err)
	}
	return db
}

/* ---------- 题目1 ---------- */

// 1. 查询部门为“技术部”的所有员工
func queryTechEmployees(db *sqlx.DB) ([]Employee, error) {
	var emps []Employee
	err := db.Select(&emps, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	return emps, err
}

// 2. 查询工资最高的员工
func queryHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var emp Employee
	err := db.Get(&emp, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	return emp, err
}

/* ---------- 题目2 ---------- */

// 查询价格大于 50 元的书籍
func queryExpensiveBooks(db *sqlx.DB) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	return books, err
}

/* ---------- 简单测试 ---------- */
func SqlxTest() {
	db := mustDB()
	defer db.Close()

	// 1
	emps, err := queryTechEmployees(db)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("技术部员工:", emps)
	}

	// 2
	topEmp, err := queryHighestSalaryEmployee(db)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("工资最高员工:", topEmp)
	}

	// 3
	books, err := queryExpensiveBooks(db)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("价格>50 的书籍:", books)
	}
}
