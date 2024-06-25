package lists

import "fmt"

func main() {

	// exercises()

	prices := []float64{10.99, 8.99}
	fmt.Println(prices, prices[1], prices[0:1])

	prices = append(prices, 12.88, 20.99)
	fmt.Println(prices)

	discountPrices := []float64{6.99, 3.99}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// prices := [4]float64{10.00, 5.00, 40.88, 20.00}
// fmt.Println(prices)
// fmt.Println(prices[2])

// productNames := [4]string{"Book", "Coffee", "Pen", "Glasses"}
// productNames[2] = "Card"

// fmt.Println(productNames[2])
// fmt.Println(productNames)

// featuredPrices := prices[1:3]
// highlightedPrices := featuredPrices[:1]
// fmt.Println(featuredPrices, highlightedPrices)

// // featuredPrices2 := prices[1:] // to the end
// // fmt.Println(featuredPrices2)

// fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// highlightedPrices = highlightedPrices[:3]
// fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
