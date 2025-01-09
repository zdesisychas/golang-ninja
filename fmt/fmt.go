package main

import (
	"fmt"      
)



//  <-- go dasturlash tilida fmt - bu satandart kutbxona paketi bo'lib , u matni malumotlarini formatlash va chiqarishni boshqarish uchun ishlatiladi .
//  fmt paketini asosiy funksyasi malummotlarni konsolga chiqarish

			// *****************************  Malumotlarni konsolga chiqarish   ********************************

//  fmt.Print <= oddiy matni chiqaradi formatsiz
//  fmt.Println <= matni chiqaradi oxrda yangi qator qoshadi 
//  fmt.Printf <= formatlanga matni chiqaradi (masalan %s , %d, %f kabi ozgaruvchilar)

// *****************************************  Malumotlarni oqish (Scan)  ******************************************

// foydalananuvchiladan malumot kiritshni oqib olish ishlatiladi.
// fmt.Scan <=  malumotni oddiy oqish 
// fmt.Scanln <= qatorni oqib olish 
// fmt..Scanf <= formatalngan malumotni oqish



// ********************************************** String bilan ishlash ********************************************

// konsolgayam chiqaradi (string) ni qaytarishi ham mumkin 

//  fmt.Sprint <= ody matni qaytaradi 
//  fmtSprintf <= formatlangan satrni qaytaradi formatalangan satrni qaytaradi konsolga chiqarmaydi
//  fmt.Sprintln <= matni yagi qator bilan qaytaradi
 
//    misol uchun :


// **************************** fmt.Print *****************************************


// func main() {
// 	name := "Bobur"
// 	age := 19 
// 	fmt.Print("Salom " )
// 	fmt.Println(name)
// 	fmt.Printf( "yoshi %d\n " , age)
// }


// ********************************* fmt.Scan  ******************************************

func main() {
	var name string
	var age int
	fmt.Println("Ismingizni kiriting")
	fmt.Scan(&name)
	fmt.Println("Yoshungizni kriting")
	fmt.Scan(&age)
	fmt.Printf("Salom %s!\n Yoshingiz %d\n", name, age)

}




// ****************************************** fmt.Sprint *****************************************

// func main() {
// 	name := "bobur"
// 	age := 19
// 	message := fmt.Sprintf("Salomo %s \n yoshingiz %d\n", name, age)
// 	fmt.Println(message)
// }