package main

import "fmt"

type Movie struct {
	Name string
}

func main() {
	fmt.Printf("请输入你的命令\n1.获得名字\n2.退出程序\n")
	var option int
	for {
		fmt.Scanf("%d", &option)
		if option == 1 {
			m := Movie{Name: "西线无战事"}
			fmt.Println(m.Name)
		} else if option == 2 {
			return
		}
	}
}
