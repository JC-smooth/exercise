package main


func main(){
	x := 1
	// f := func(){
	// 	println(x)
	// }
	// f()  // 1
	// x = 2
	// x = 3
	// f()  // 3
	func() {
		println(&x)
	} ()
	println(&x)
}
