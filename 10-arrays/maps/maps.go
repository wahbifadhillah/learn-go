package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	// []string{"https://google.com", "https://ubahgambar.id"}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["UbahGambar.ID"] = "https://www.ubahgambar.id"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
