// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// }

//NOTE: A map maps keys to values. The zero value of a map is nil. A nil map has no keys, nor can keys be added. To initialize a map, use the builtin make: make(map[keyType]valType).

///////////////////////////////////////////////////////////

// Map Literals

// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

// func main() {
// 	fmt.Println(m)
// }

// NOTE: A map literal is like an array or slice literal, but the keys are required. The elements are key:value pairs.


/////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37.42202, -122.08408},
// }

// func main() {
// 	fmt.Println(m)
// }

// NOTE: If the top-level type is just a type name, you can omit it from the elements of the literal.

/////////////////////////////////////////////////////////////////

// Mutating Maps

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// NOTE: The builtin function delete removes key/value pairs from a map.
// The optional second return value when retrieving a value from a map indicates if the key was present in the map.
// If the key was not present, then the zero value for the map's value type is returned.
// For example, in the above code, the zero value for an int is 0.
// The variable ok is true if the key was present in the map, and false otherwise.
// The statement v, ok := m["Answer"] is an example of a "comma ok" idiom that is used widely in Go code.
// The "comma ok" idiom is also used when retrieving values from channels and when asserting types.