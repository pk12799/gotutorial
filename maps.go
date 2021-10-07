package main

import (
	"fmt"
)

func main() {
	// mapss := make(map[string]int)
	// mapss = map[string]int{
	// 	"bhopal":  12465,
	// 	"indore":  24587,
	// 	"delhi":   21457,
	// 	"mumbai":  1548254,
	// 	"kolkata": 21457,
	// }
	// fmt.Println(mapss["bhopal"])
	// mapss["chennai"] = 1254785
	// fmt.Println(mapss)
	// inmaps := make(map[int]string)
	// inmaps = map[int]string{
	// 	1: "bhopal",
	// 	2: "delhi",
	// 	3: "kolkata",
	// 	4: "mumbai",
	// }
	// inmaps[45] = "chennai"
	// inmaps[41] = "indore"

	// fmt.Println(inmaps)
	// delete(inmaps, 4)
	// inmaps[5] = "fhs"
	// fmt.Println(inmaps)

	// fmt.Println(inmaps[4])
	inmaps := make(map[int]int)
	inmaps = map[int]int{
		4:   45745,
		45:  25478,
		478: 4781,
	}
	inmaps[451] = 487525
	fmt.Println(inmaps)
	mapss := inmaps
	fmt.Println(inmaps, mapss)
	delete(inmaps, 4)
	mapss[45789] = 14785141
	fmt.Println(inmaps, mapss)
	// stmaops := make(map[string]string)
	// stmaops["hihh"] = "hjgfhdfhgsfd"

	// stmaops = map[string]string{
	// 	"sci":   "ddfdd",
	// 	"sdfdf": "sahff",
	// }
	// stmaops["dfhg"] = "sdfdf"
	// fmt.Println(stmaops)

}
