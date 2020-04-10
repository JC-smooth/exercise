package main

// boiling 输出水的沸点
import "fmt"

const boilingF = 212.0

func main(){
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g℉ or %g℃\n", f, c)
}
