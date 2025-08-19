package task5

import "fmt"

var SliceInte = []int{50, 75, 66, 20, 32, 90}

func Menambahkan(value int) {
	fmt.Println("Sebelum di append : ")
	fmt.Println(SliceInte)
	// dimana insertnya? setelah 66 = index 3
	index := 3
	SliceInte = append(SliceInte, 0)
	copy(SliceInte[index+1:], SliceInte[index:])
	SliceInte[index] = value
	fmt.Println("Setelah di append :")
	fmt.Println(SliceInte)
}
