package main

import (
	"fmt"
)

type User struct {
	Name string
	Focus int
	Signature string
	VIP bool

}

func main()  {
	  a :=User{
		Name:      string("国际华语辩论赛邀请撒"),
		Focus:     int(78000),
		VIP:       false,
		Signature: string("辩论竞技专业化,尽在国际华语辩论赛"),
	}

	  var b User
	  b.Name="leetcode"
	  b.Focus=6165
	  b.Signature="力加官方账号"
	  b.VIP=false
fmt.Println(a,b)
}


