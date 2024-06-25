package lists

import "fmt"

// Time to practice what you learned!

//  1. Create a new array (!) that contains three hobbies you have
//     Output (print) that array in the command line.
//  2. Also output more data about that array:
//     - The first element (standalone)
//     - The second and third element combined as a new list
//  3. Create a slice based on the first element that contains
//     the first and second elements.
//     Create that slice in two different ways (i.e. create two slices in the end)
//  4. Re-slice the slice from (3) and change it to contain the second
//     and last element of the original array.
//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
//  7. Bonus: Create a "Product" struct with title, id, price and create a
//     dynamic list of products (at least 2 products).
//     Then add a third product to the existing list of products.

type Product struct {
	id    int32
	title string
	price float64
}

func exercises() {
	// #1
	hobbies := [3]string{"Youtube", "Sports", "Games"}
	fmt.Println(hobbies)

	// #2
	fmt.Println(hobbies[0], hobbies[1:])

	// #3
	hobbiesFS1 := hobbies[0:2]
	hobbiesFS2 := hobbies[:2]

	fmt.Println(hobbiesFS1, hobbiesFS2)

	// #4
	fmt.Println(cap(hobbiesFS1))
	hobbiesFS3 := hobbiesFS1[1:3]
	fmt.Println(hobbiesFS3)

	// #5
	courseGoals := []string{"Learn the fundamentals of Go Lang", "Create practical projects"}

	fmt.Println(courseGoals)

	// #6
	courseGoals[1] = "Improve tech skills"
	courseGoals = append(courseGoals, "Create practical projects")

	fmt.Println(courseGoals)

	// #7
	product1 := Product{
		id:    1,
		title: "1kg Rice",
		price: 10.99,
	}

	product2 := Product{
		id:    2,
		title: "1kg Beans",
		price: 12.99,
	}

	products := []Product{product1, product2}
	fmt.Println(products)

	product3 := Product{
		id:    3,
		title: "3kg Sugar",
		price: 20.99,
	}

	products = append(products, product3)

	fmt.Println(products)
}
