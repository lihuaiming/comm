package comm

import "github.com/satori/go.uuid"

//GUIDPool GUID池
type GUIDPool struct {
	size uint32      //缓存池大小
	pipe chan string //缓存池
}

//InitPool ...
func (p *GUIDPool) InitPool() {
	p.pipe = make(chan string, p.size)
	go func() {
		for {
			u, _ := uuid.NewV4()
			p.pipe <- u.String()
		}
	}()
}

//GUID ...
func (p *GUIDPool) GUID() string {
	return <-p.pipe
}

//NewGUIDPool ...
func NewGUIDPool(i uint32) *GUIDPool {
	g := &GUIDPool{size: i}
	g.InitPool()
	return g
}

//NewGUIDPool100 ...
func NewGUIDPool100() *GUIDPool {
	return NewGUIDPool(100)
}

//NewGUIDPool1000 ...
func NewGUIDPool1000() *GUIDPool {
	return NewGUIDPool(1000)
}

//NewGUIDPool10000 ...
func NewGUIDPool10000() *GUIDPool {
	return NewGUIDPool(10000)
}
