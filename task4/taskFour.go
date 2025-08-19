package task4

import "fmt"

var mapSkala = map[string]int{
	"C": 5,
	"R": 4,
	"K": 5,
	"F": 9,
}

var skala = []string{"C", "R", "K", "F"}

func KonversiSuhu(suhuInput int, skalaInput string) {
	fmt.Println(mapSkala)
	if skalaInput == "K" {
		suhuInput -= 273
	}
	if skalaInput == "F" {
		suhuInput -= 32
	}

	for i := 0; i < len(mapSkala); i++ {
		if skala[i] == skalaInput {
			continue
		}
		var suhuTujuan = (mapSkala[skala[i]] / mapSkala[skalaInput]) * suhuInput
		if skala[i] == "K" {
			suhuTujuan += 273
		}
		if skala[i] == "F" {
			suhuTujuan += 32
		}
		fmt.Println("Suhu dalam ", skala[i], " adalah ", suhuTujuan)
	}
}
