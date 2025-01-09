// **************************************** struct ************************************************************
// struct bu murrakab malumot tuzilamalarni yaratish uchun ishlatilinadi
// Struct bir nechta turli malumotlarni bitta birlikda birlashtirish imkonyuatni beradi
// struct kopincha interfeyslar bilan ishlaydi

// ---------------------------- KEYINNGI FAYLNI OCH ---------------------------------------

package main

import "fmt"

func main() {

	myBob1 := NewBob1("bobur")

	myBob1.UpdateTip(50)

	myBob1.updateItems("cola", 9.9,)
	myBob1.updateItems("mastava", 19.9,)
	myBob1.updateItems("worbo", 10.9,)


	fmt.Println(myBob1.format())

}

