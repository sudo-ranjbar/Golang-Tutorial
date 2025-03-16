package main

func main() {

	num := 20
	println("initial value of num is: ", num)
	println("initial address of num is: ", &num)
	changeNumByVal(num)
	println("the value of num after changeVal function: ", num)
	changeNumByRef(&num)
	println("the value of num after changeRef function: ", num)

}

func changeNumByRef(num *int) {
	println("changeRef address ", num)
	*num = *num + 2
}

func changeNumByVal(num int) {
	println("changeVal address ", &num)
	num = num + 2
}
