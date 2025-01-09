package main

import "fmt"


func main(){
	fmt.Println("hi")

	func(){
		fmt.Println("anonmy func")
	}()
}




func main1() {
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(increment1())
	fmt.Println(increment1())
	fmt.Println(increment1())
}


func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}



func increment1() int {
	count := 0 
	count++
	return count
}



func main2() {
	message := "hi my name is bobur "
	printMessage(message)	
}


func printMessage (message string) {
	message += "i'll be golang ninja"
	fmt.Println(message)

}