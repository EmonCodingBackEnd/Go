package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	//map[name:ccmouse course:golang site:imooc quality:notbad]
	fmt.Println(m)

	m2 := make(map[string]int)
	//map[]
	fmt.Println(m2) //m2 == empty map

	var m3 map[string]int
	//map[]
	fmt.Println(m3) //m3 == nil

	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	courseName, ok := m["course"]
	//golang true
	fmt.Println(courseName, ok)
	courseName, ok = m["couse"]
	// false
	fmt.Println(courseName, ok)

	if courseName, ok = m["cause"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key deso not exist")
	}

	name, ok := m["name"]
	//ccmouse true
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	// false
	fmt.Println(name, ok)
}
