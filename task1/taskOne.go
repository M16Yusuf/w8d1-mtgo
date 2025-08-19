package task1

type Data struct {
	Nama        string
	Foto        string
	Email       string
	Umur        uint
	NoHp        string
	StatusNikah bool
	ListSchool  map[string]string
}

func TaskOne() Data {
	dataDiri := Data{
		Nama:        "Muhammad Yusuf",
		Foto:        "/img/profile_2017.jpg",
		Email:       "yusufsmd58@gmail.com",
		Umur:        24,
		NoHp:        "082240563847",
		StatusNikah: false,
		ListSchool: map[string]string{
			"unikom":      "Teknik Informatika",
			"ma al-irfan": "ipa",
		},
	}
	return dataDiri
}
