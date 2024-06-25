package maps

import "fmt"

// map is key -> value

func main() {
	// websites := []string{"https://google.com", "https://aws.com"} // slice
	websites := map[string]string{
		"Amazon Web Services": "https://aws.com",
		"Google":              "https://google.com",
	} // map

	fmt.Println(websites, websites["Google"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites["LinkedIn"])

	websites["LinkedIn"] = "https://linkedin.com/murilo-perosa"
	fmt.Println(websites["LinkedIn"])

	delete(websites, "Google")
	fmt.Println(websites)

}
