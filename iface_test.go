package pwn

import "testing"
import "fmt"

func TestGetAllIfaceAddrs(t *testing.T) {
	// TODO: better tests, currently only tests the function not crash / panicing
	// Not much we can 'test' though since it doesnt take input except
	// GetIFaceByName and GetIFaceByIndex
	GetAllIfaceAddrs()

	// Few strings for testing this function
	testcasesName_str := [4]string{"lo", "eth0", "wlan0", "fake"}
	
	for i := 0; i < len(testcasesName_str); i++ {
		fmt.Printf("\nTestcase: %d | GetIFaceByName(\"%s\") | \n", i, testcasesName_str[i])
		fmt.Println( GetIFaceByName(testcasesName_str[i]) )
	}


	// Few ints for testing this function
	testcasesIndex_int := [8]int{0, 1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(testcasesIndex_int); i++ {
		fmt.Printf("\nTestcase: %d | GetIFaceByIndex(%d) | \n", i, testcasesIndex_int[i])
		fmt.Println( GetIFaceByIndex(testcasesIndex_int[i]) )
	}

}
