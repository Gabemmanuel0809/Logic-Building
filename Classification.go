package main 
import "fmt"

func main() {
  // just for practice purpose only[reading input + slice + map]
	var adultCitizen = []string{} 
	var minorCitizen = []string{}

  var name string 
	fmt.Println("Input your name")
	fmt.Scanf("%s", &name)

	var age int 
	fmt.Println("Input your age: ")
	fmt.Scanf("%d", &age)

	var classification string 
	if(age < 18) {
		 classification = "Minor"
		 fmt.Println(classification)
	} else {
		 classification = "Adult"
		 fmt.Println(classification)
	}
     
	// super very needed(nope not at all to be honest)
	citizen := make(map[string]string)
	citizen[name] = classification
	value, exist := citizen[name]

	fmt.Println(value, exist)
    
	if(exist && classification == "Adult") {
        adultCitizen = append(adultCitizen, name)
		fmt.Println("Inserted into Adult class")
	} else if(exist && classification == "Minor") {
        minorCitizen = append(minorCitizen, name)
		fmt.Println("Inserted into Minor class")
	} else {
		// of course this will never run in the first place
		fmt.Println("Error, Missing keys/values detected")
	}
    
}
