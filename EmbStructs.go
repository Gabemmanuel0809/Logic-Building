package main
import "fmt"

type Product struct {
   Name string
   Price int 
}

type Brand struct {
   Class string
   Product
}

func main() {
	  // H&M's greatest rival
    var HnR = Brand {
		Class: "Clothes",
		Product: Product {
            Name: "Sweat Shirt",
			Price: 25,
		},
	}
    
	// Defenitely not Nissan
	var Neesan = Brand {
		Class: "Vehicle",
		Product: Product {
		   Name: "HJ-56",  // Just imagine it as a small van
		   Price: 25000,
		},
	}

	fmt.Println(HnR.Product)
	fmt.Println("Vehicle: ", Neesan.Product.Name)
}
