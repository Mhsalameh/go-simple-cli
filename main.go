package main

import (
	"flag"
	"fmt"
	"strconv"
	// "os"
)

var (
	method string 
)

func main(){
	flag.StringVar(&method,"method", "" ,"which math method you want to use (add,sub,mult,div)")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		var res float64
		var resm float64 = 1
		for i, arg := range args {
			num,_ := strconv.ParseFloat(arg,64)

			switch method {
			case "add":
				fmt.Printf("%v",arg)
				if i != len(args)-1 {
					fmt.Printf("+")
				} else {
					fmt.Printf("=")
				}
				res+=num
			case "sub":
				fmt.Printf("%v",arg)
				if i != len(args)-1 {
					fmt.Printf("-")
				} else {
					fmt.Printf("=")
				}
				res-=num
			case "mult":
				fmt.Printf("%v",arg)
				if i != len(args)-1 {
					fmt.Printf("*")
				} else {
					fmt.Printf("=")
				}
				resm*=num
			case "div":
				if i % 2 ==0 {
					fmt.Printf("(%v",arg)
				} else {
					fmt.Printf("%v)",arg)
				}
				if i != len(args)-1 {
					fmt.Printf("/")
				} else {
					if i % 2 !=0 {

						fmt.Printf("=")
					}else {
						fmt.Printf(")=")
					}
				}
				if i == 0{
					resm=num
				} else {
					
					resm/=num
				}
			}	
		}	
		if method=="mult" || method=="div"{
			fmt.Println(resm)
		}
		if method == "add" || method == "sub"{
			fmt.Println(res)
		}
	}	
}