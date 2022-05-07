package comm

import (
	"testing"
)

var (
	gp1    *GUIDPool = NewGUIDPool(1)
	gp100  *GUIDPool = NewGUIDPool100()
	gp1000 *GUIDPool = NewGUIDPool1000()
)

func TestGUIDPool1(t *testing.T) {
	for index := 0; index < 200000; index++ {
		guid := gp1.GUID()
		t.Log(guid)
	}

}
func TestGUIDPool2(t *testing.T) {
	for index := 0; index < 200000; index++ {
		guid := gp1000.GUID()
		t.Log(guid)
	}

}

//func TestGUIDPool3(t *testing.T) {
//	for index := 0; index < 200000; index++ {
//		guid := uuid.NewV4().String()
//		t.Log(guid)
//	}
//
//}
func BenchmarkGUIDPool1(b *testing.B) {
	for index := 0; index < b.N; index++ {
		gp1.GUID()
	}
}
func BenchmarkGUIDPool2(b *testing.B) {
	for index := 0; index < b.N; index++ {
		gp1000.GUID()
	}

}

//func BenchmarkGUIDPool3(b *testing.B) {
//	for index := 0; index < b.N; index++ {
//		uuid.NewV4().String()
//	}
//
//}
