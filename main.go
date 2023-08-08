package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
	IsHim     bool
}

// (info *User) - ties this function to the User struct, giving it access to info from that struct
func (info *User) greeting() string {
	return "Hello Mr. " + info.LastName
}

func main() {
	user := User{
		FirstName: "Stefan",
		LastName:  "Madzarac",
		Age:       30,
		IsHim:     false,
	}

	// map[KeyType]ValueType
	animals := make(map[string]string)

	animals["Lion"] = "Big Cat"

	profiles := make(map[string]User)

	profiles["me"] = user

	// slice_name := []datatype{values}
	cars := []string{"Ferrari", "Lamborgini"}
	cars = append(cars, "Bugatti")

	fmt.Println(user.greeting())
	fmt.Println(animals["Lion"])
	fmt.Println(profiles["me"].FirstName)
	fmt.Println(cars[1:3])

	// Decision Structures
	if user.IsHim {
		fmt.Print("Is that guy")
	}

	// Switch statements
	switch user.IsHim {
	case false:
		fmt.Println("Is not that guy")

	case true:
		fmt.Println("Is that guy")
	}
}
