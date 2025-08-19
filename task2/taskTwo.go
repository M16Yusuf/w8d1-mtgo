package task2

const realEmail string = "yusufsmd58@gmail.com"
const realPassword string = "@admin123"

func TryLogin(email string, password string) (string, bool) {
	if email != realEmail && password != realPassword {
		return "email dan password salah", false
	} else if email == realEmail && password != realPassword {
		return "password salah", false
	} else if email != realEmail && password == realPassword {
		return "email salah", false
	} else {
		return "Berhasil Login", true
	}
}
