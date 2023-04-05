package main

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
	Implement Sort
		- User should be able to sort by any attribute
		- IMPORTANT : Must use sort.Sort()
*/

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
}
