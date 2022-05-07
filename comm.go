package comm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

/*获取当前文件执行的路径*/
func GetCurrPath() string {
	/*
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		path = strings.Replace(path, "\\", "/", 0)
		splitstring := strings.Split(path, "/")
		size := len(splitstring)
		splitstring = strings.Split(path, splitstring[size-1])
		ret := strings.Replace(splitstring[0], "\\", "/", size-1)
		return ret
	*/
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
func Wait_random(maxSleep int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nSleep := r.Intn(maxSleep)
	for {
		if nSleep != 0 {
			break
		} else {
			nSleep = r.Intn(maxSleep)
		}

	}
	fmt.Println("sleep : ", nSleep)
	time.Sleep(time.Duration(nSleep) * time.Second)
}
func Random_n(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func GetBitBool(n, nMove uint) bool {
	n = n >> nMove
	n &= 1
	if n == 1 {
		return true
	}
	return false
}

//对一个字符串进行MD5加密,不可解密
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

////获取一个Guid值
func GetGuid() string {
	u, _ := uuid.NewV4()
	return u.String()
}

//FileNameAndSuffix
func FileNameAndSuffix(filename string) string {
	return path.Base(filename)
}

//FileSuffix
func FileSuffix(filename string) string {
	return path.Ext(path.Base(filename))
}

//FileOnlyName
func FileOnlyName(filename string) string {
	filenameWithSuffix := path.Base(filename)
	fileSuffix := path.Ext(filenameWithSuffix)
	return strings.TrimSuffix(filenameWithSuffix, fileSuffix)
}

//CopyFile ...
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
