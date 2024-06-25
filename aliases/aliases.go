package aliases

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

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
	courseRatings := make(floatMap, 3) // cap
	courseRatings["go"] = 10
	courseRatings["php"] = 10
	courseRatings["js"] = 9

	courseRatings["vue"] = 8 // re-allocate cap

	courseRatings.output()
}
