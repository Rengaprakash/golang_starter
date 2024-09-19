package datatypes

import "fmt"

/**
Map DataType
maps can be used for single datatype of key and values pair
**/

/*
*

	Package level variable
	can be accessed within this package

*
*/
var mapValues = make(map[string]int)

func initMap() {
	// defining map with key as string and values as int type
	// make is used for creating empty map
	mapValues["Age"] = 5
	mapValues["size"] = 4
}

func getMapDetails() {
	mapValues["length"] = 25
	fmt.Printf("\n Map: %v", mapValues)
	// Accessing single value by key
	fmt.Printf("Age: %v", mapValues["length"])
}
