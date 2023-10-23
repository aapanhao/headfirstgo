package strtoint

import (
	"strconv"
)

func StrToInt64() int64 {
	data := "123"
	i, _ := strconv.ParseInt(data, 0, 0)
	return i
}

//
//func main() {
//	fmt.Println(StrToInt64())
//}
