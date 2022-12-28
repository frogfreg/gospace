package main

// var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

func main() {

	products := []Product{{"Kayak", 279}, {"Lifejacket", 49.95}, {"Soccer ball", 19.50}}

	ProductSlices(products)

	for _, p := range products {

		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)

	}

}
