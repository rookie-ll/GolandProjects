package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int    `json:"id" db:"title" xml:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

//Class 班级
type Class struct {
	Title    string    `json:"title"`
	Students []Student `json:"students"`
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}

	// Json 序列化
	_, err := json.Marshal(c)
	if err != nil {
		return
	}
	//fmt.Printf("%s\n", data)

	// Json 反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := Class{}
	err = json.Unmarshal([]byte(str), &c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

}
