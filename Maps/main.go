package main

import "fmt"

func main() {

	mp := map[int]int64{} // empty map initialzed
	// var mp2 map[string]string;
	// mp2 := make(map[string]string);

		
	// insertion
	mp[0] = 1
	mp[12] = 123
	mp[2] = -123
	mp[3] = 2131
	mp[4] = -123
	mp[5] = 23
	mp[6] = 90

	// temp := len(mp);

	mp[12] = 123123 // updation
	fmt.Println(mp[12])

	//deletion
	delete(mp, 12) // key

	// iteration
	for key, value := range mp {
		fmt.Println(key, value)
	}
}

// Notes

// When we index a map we get 2 values 1 is the value for the particular key
// and the other is a boolean that indicates weather the value exists or not
