package datatypes

import "fmt"

// Structs can be used for multiple datatypes collection
// Somewhat similar to classes
func testStruct() {
	type userData struct {
		firstName string
		age       int
		lastName  string
	}

	var userDataVal = userData{
		firstName: "test",
		age:       56,
		lastName:  "user",
	}

	fmt.Printf("\n Value of struct is %v", userDataVal)
	fmt.Printf("\n Age is %v", userDataVal.age)
}
