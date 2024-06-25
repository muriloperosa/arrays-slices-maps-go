package make

import "fmt"

func main() {

	// make slices
	userNames := make([]string, 2, 4) // len, cap
	userNames[0] = "Bernardo"
	userNames[1] = "Juno"
	// userNames[2] = "Simone" // error
	userNames = append(userNames, "Murilo")
	userNames = append(userNames, "Rafaela")

	userNames = append(userNames, "Gilmar") // re-allocate cap

	fmt.Println(userNames)

	// make maps
	courseRatings := make(map[string]float64, 3) // cap
	courseRatings["go"] = 10
	courseRatings["php"] = 10
	courseRatings["js"] = 9

	courseRatings["vue"] = 8 // re-allocate cap

	fmt.Println(courseRatings)
}
