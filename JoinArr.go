package main
import "fmt"

func main() {
    var users = [...]string{"John", "Bob", "Peter"}
	var orders = [...]string{"Burger", "Fries", "Doughnout"}

  // Joining two arrays in a map
	var usr_ordr = make(map[string]string)
	usr_ordr[users[0]] = orders[2]
	usr_ordr[users[1]] = orders[0]
	usr_ordr[users[2]] = orders[1]

	for key, value := range usr_ordr {
		fmt.Println(key, ":", value)
	}
}
