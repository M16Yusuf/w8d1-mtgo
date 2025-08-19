package main

import (
	"fmt"

	"github.com/m16yusuf/w8d1-mtgo/task1"
	"github.com/m16yusuf/w8d1-mtgo/task2"
	"github.com/m16yusuf/w8d1-mtgo/task3"
	"github.com/m16yusuf/w8d1-mtgo/task4"
)

func main() {
	// Tugas 1
	fmt.Println("============================")
	fmt.Println("==========TASK 1============")
	fmt.Println(task1.TaskOne())
	fmt.Println("============================")
	fmt.Println("==========TASK 2============")

	// Tugas 2
	isLogin := false
	for isLogin != true {
		var emailInput string
		var passInput string
		fmt.Println("masukan email : ")
		fmt.Scanln(&emailInput)
		fmt.Println("masukan Password : ")
		fmt.Scanln(&passInput)
		fmt.Println("============================")
		status, resIsLogin := task2.TryLogin(emailInput, passInput)
		fmt.Println(status)
		isLogin = resIsLogin
	}

	// Tugas 3
	fmt.Println("============================")
	fmt.Println("==========TASK 3============")
	fmt.Println(task3.WhichBigger(task3.MyBulatPositif))

	// Tugas 4
	fmt.Println("============================")
	fmt.Println("==========TASK 4============")
	var suhuInput int
	var skalaInput string
	fmt.Println("masukan nilai suhu : ")
	fmt.Scanln(&suhuInput)
	fmt.Println("masukan skala suhu tersebut : ")
	fmt.Scanln(&skalaInput)
	task4.KonversiSuhu(suhuInput, skalaInput)
}
