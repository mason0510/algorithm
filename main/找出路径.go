package main

import (
	"fmt"
	_ "path/filepath"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/a//bc/d//././/"))
}
func simplifyPath(path string) string {
	reg, _ := regexp.Compile("/+")
	fmt.Println("reg",reg)
	dirs := reg.Split(path, -1)
	fmt.Println("dirs",dirs)
	res := []string{}
	for _, v := range dirs {
		p := string(v)
		if p == ".." {
			if len(res) > 0 {
				//获取0到指定位置的值
				res = res[:len(res)-1]
			}
		} else if p != "." && p != "" {
			res = append(res, p)
		}
	}
	str := strings.Join(res, "/")
	return "/" + str
}
