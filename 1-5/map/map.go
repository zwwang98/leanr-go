package main

import "fmt"

func main() {
	declareMap()
	traverseMap()
	checkKeyExistenceInMap()
	deleteInMap()
}

func declareMap() {
	fmt.Println("[declareMap] To declare map, use \"map[{key type}]{value type}\"")
	fmt.Println("For example,\n" +
		"m := map[string]string{\n\t\"name\":  \"ziweiwang\",\n\t\"place\": \"seattle\",\n}")
	m1 := map[string]string{
		"name":  "ziweiwang",
		"place": "seattle",
	}
	fmt.Println(m1)

	fmt.Println("\n[declareMap] To declare map, one could also use \"make(map[{key type}]{value type}\"")
	fmt.Println("For example,\n" +
		"m2 := make(map[string]int)")
	m2 := make(map[string]int)
	fmt.Println("m2 =", m2)
	fmt.Println("m2 == nil:", m2 == nil)
	fmt.Println("m2 == empty map", len(m2) == 0)

	fmt.Println("\n[declareMap] To declare map, one could also use \"var m map[{key type}]{value type}\"")
	fmt.Println("For example,\n" +
		"var m3 map[string]int")
	var m3 map[string]int
	fmt.Println("m3 =", m3)
	fmt.Println("m3 == nil:", m3 == nil)
	fmt.Println("m3 == empty map", len(m3) == 0)
}

// use for loop on the range of map to traverse the map
// this traversal is not ordered
func traverseMap() {
	fmt.Println("\n[traverseMap]")
	m1 := map[string]string{
		"name":  "ziweiwang",
		"place": "seattle",
	}
	fmt.Println(m1)

	for k, v := range m1 {
		fmt.Println("key =", k, ", value =", v)
	}
}

/*
1. Use ok to check some key's existence in the map
2. If the key is not in the map, no error message, but the returned value will be default value for the value type
*/
func checkKeyExistenceInMap() {
	fmt.Println("\n[checkKeyExistenceInMap]")
	m := map[string]string{
		"name":  "ziweiwang",
		"place": "seattle",
	}
	fmt.Println(m)

	name, ok := m["name"]
	if ok {
		fmt.Println("name =", name)
	} else {
		fmt.Printf("map does not have the key {%s}\n", "name")
	}

	occupation, ok := m["occupation"]
	if ok {
		fmt.Println("occupation =", occupation)
	} else {
		fmt.Printf("map does not have the key {%s}\n", "occupation")
	}
}

func deleteInMap() {
	fmt.Println("\n[deleteInMap]")
	m := map[string]string{
		"name":  "ziweiwang",
		"place": "seattle",
	}

	fmt.Println("Before delete key {name}")
	name, ok := m["name"]
	fmt.Printf("{name: %v}, {ok: %v}\n", name, ok)
	fmt.Println(m)

	delete(m, "name")

	fmt.Println("After delete key {name}")
	name, ok = m["name"]
	fmt.Printf("{name: %v}, {ok: %v}\n", name, ok)
	fmt.Println(m)
}

/*
- What type could be used as the key of the map in golang?
  1. All built-in types other than slice, map, and function.
  2. All structs without slice, map, or function
  No need to implement hashCode() or equals() like Java does.
*/
