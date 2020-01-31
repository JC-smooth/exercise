package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var allStudents map[int64]*student

func showAllStudents() {
	for k, v := range allStudents {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func newStudent(id int64, name string) *student{
	return &student{
		id:   id,
		name: name,
	}
}

func addStudents() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	allStudents[id] = newStu
}

func removeStudents(){
	var deleteID int64
	fmt.Print("请输入学生学号:")
	_, _ = fmt.Scanln(&deleteID)
	delete(allStudents, deleteID)
}


func main() {
	for {
		var choice int8
		fmt.Println("欢迎使用学生管理系统！")
		fmt.Println("1.查看学生")
		fmt.Println("2.新增学生")
		fmt.Println("3.删除学生")
		fmt.Println("4.退出")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudents()
		case 3:
			removeStudents()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("没有该选项！")
		}
	}
}
