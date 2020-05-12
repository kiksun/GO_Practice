package main

import "fmt"

func Checknumber(x int)bool{
	if x==3  {return false 
	}else {return true}
}
func main() {
	msg :="hello world"
	a,b :=10,15
	var (
		c int
		d string
	)
	c=20
	d="hoge"
	f:=Checknumber(3)
	fmt.Println(a,b,c,d,msg)
	fmt.Println(f)
}