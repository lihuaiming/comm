package comm

import "testing"
import "fmt"
import "strconv"

func TestComm(t *testing.T) {
	v := 3.14159265351592653515926535

	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'f', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

}
