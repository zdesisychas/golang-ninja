
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
// ************************************ Pointers **********************************************
//    Pointer turlari : har bir point biror bir malumotlarini korsatadi masalan; *int integer qiymatini pointeri
//	  Pointer qiymatini olish : * qiymatini oladi
//    Pointerga manzilini : & operatori  orqali oddiy xotira manzilini olish
//    Xavfizlik : go tilida pointer arifmetikasi yoq bu xatolkni kamaytiradi
//    Nil pointer har bir pointerni boshlangich qiymati nil bolishi mumkin
//    Pointer asosan pass by reference usulida uzatish yoki malimtolani asl nusxasini ozgartrishga ishlatilinadi.
//	  Pointerlar yordamida siz bror qiymatning asl nusxasga (xotira joyiga ) togridan togri murojat qilasiz
//
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////



package main



import "fmt"




func main() {
	name := "xasan"
	age := 21
	fmt.Println(name)
	fmt.Println(age)

	
	fmt.Println("names' memory address is" , &name)  //<= & xotira manzilini ovoti 
	fmt.Println("age's memory address is", &age )	 //<= & xotira manzilini ovoti 

	fmt.Println("names' memory address is" , *&name)  //<= * qiymatini ovoti 
	fmt.Println("age's memory address is", *&age )	  //<= * qiymatini ovoti 
}




func main1() {
	name := "xasan"
	age := 21
	fmt.Println(name)
	fmt.Println(age)

	p:=age  //<= ozgartish
	b:=name //<= ozgartish


	fmt.Println("names' memory address is" , &b)   //<= & xotira manzilini ovoti 
	fmt.Println("age's memory address is", *&p)	   //<= * qiymatini ovoti 
}



func updateName(x *string){
	*x = "bobur"
}

func main2() {
	name := "xasan"

	// updateName(name)

	m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m)


	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}