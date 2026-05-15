package main
import "fmt"

func main() {
	numbers := [...]int{10, 2, 50, 93, -5, -7, 13, 22, 85, 73, 21, 18, 11, 14, 15}
    odd := []int{}
	even := []int{}
	filter := []int{}
    filter2 := []int{}
	filter3 := []int{}

    fmt.Print("All Numbers: ")
	for _, value := range numbers {
		 fmt.Print(" ",value)
	}

	fmt.Println("----------")

	fmt.Println("----------")

	for _, value := range numbers {
	    if(value %2 != 0) {
   	      odd = append(odd, value)
	    }
	}
	fmt.Println("Odd Numbers: ", odd)

	fmt.Println("----------")

	for _, value := range numbers {
		 if(value %2 == 0) {
			 even = append(even, value)
		 }
	}
	fmt.Println("Even Numbers: ", even);

	fmt.Println("----------")

	for _, value := range numbers {
		 if(value <= 15) {
			 filter = append(filter, value)
		 }
	}
	fmt.Println("15 or lesser than 15: ", filter)

	fmt.Println("----------")

	for _, value := range numbers {
	   if(value > 15 && value <= 50) {
		  filter2 = append(filter2, value)
	   }
	}
	fmt.Println("50 or lesser than 50: ", filter2)

	fmt.Println("----------")

	for _, value := range numbers {
		 if(value > 50) {
			  filter3 = append(filter3, value)
		 }
	}
	 fmt.Println("Greater than 50: ", filter3)

}
