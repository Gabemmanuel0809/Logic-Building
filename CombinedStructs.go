package main 
import "fmt" 

type User struct {
  Id   int 
  Name string
}

type Item struct {
  Id int  
  Name string 
}

type Orders_Placed struct {
	User_name  string
	Item_name  string
}

func main() {
  var usr = User {
		 Id: 1,
		 Name: "John",
	}
    
	var usr2 = User {
		 Id: 2,
		 Name: "Jacob",
	}

	var usr3 = User {
		 Id: 3,
		 Name: "James",
	}

	var itm = Item {
		 Id: 1,
		 Name: "Doughnut",
	}

	var itm2 = Item {
		 Id: 2,
		 Name: "Ice Cream",
	}

	var itm3 = Item {
		 Id: 3,
		 Name: "Coffee",
	}

	var ordplcd1 = Orders_Placed {
        User_name: usr.Name,
		Item_name: itm3.Name, 
	}

	var ordplcd2 = Orders_Placed {
        User_name: usr2.Name,
		Item_name: itm.Name,
	}

	var ordplcd3 = Orders_Placed {
        User_name: usr3.Name,
		Item_name: itm2.Name ,
	}

	fmt.Println(ordplcd1)
	fmt.Println(ordplcd2)
	fmt.Println(ordplcd3)
}
