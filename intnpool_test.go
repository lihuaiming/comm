package comm

import (
	"testing"
)


func TestNewIntnPool(t *testing.T) {
	pl := NewIntnPool(50,10000,4)
	for index := 0; index < 1000; index++ {
		guid := pl.Intn()
		t.Log(guid)
	}

}
var ip1  *IntnPool = NewIntnPool(50,1000,4)

func BenchmarkIntnPool(b *testing.B) {
	for index := 0; index < b.N; index++ {
		ip1.Intn()
	}
}