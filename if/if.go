// ******************************* if *************************************************

// if sharti yordamida turli holatlarni teksjirish va ularga asoslangan holatlardni amalga oshirish mumkin
// if sharti faqatgina berilgan shart togri (true) bolgandad bajariladi

// Shart faqat true yoki false qiymatiga ega bo'lishi kerak.

// go tilida if sharti sodda va kuchli bolib , shartli mantqni boshqarishda keng qollaniladi

package main

import (
	"fmt"
)

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("Siz voyaga yetgansiz!")
	} else {
		fmt.Println("Siz voyaga yetmagansiz!")
	}
}
