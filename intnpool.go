package comm

import (
	"math/rand"
	"time"
)


//InitPool
type IntnPool struct {
	size uint32      //缓存池大小 default 1w
	pipe chan int
	pthrs int //计算线程数 default 2
	nrange int //计算线程数 default 100
}

//InitPool ...
func (p *IntnPool) InitPool() {
	if p.size <=0 {
		p.size = 10000
	}
	if p.pthrs <=0 {
		p.pthrs = 2
	}
	if p.nrange <=0 {
		p.nrange = 100
	}

	p.pipe = make(chan int, p.size)

	for thr :=0;thr < p.pthrs  ;thr++  {
		go func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				p.pipe <- r.Intn(p.nrange)
			}
		}()
		time.Sleep(864*time.Nanosecond)
	}

}

//GUID ...
func (p *IntnPool) Intn() int {
	return <-p.pipe
}

//NewIntnPool ...
func NewIntnPool(nrange int,poolsize uint32 ,calcThrs int ) *IntnPool {
	g := &IntnPool{
		size :poolsize,
		pthrs :calcThrs,
		nrange:nrange,
	}
	g.InitPool()
	return g
}
