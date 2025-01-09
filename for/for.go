// ************************************* For ****************************************

// go tilida for loop bu biror kod qatorida bir necha marta takrorlash uchun ishlatiladigan asosiy sikl operatori . for loop yordamida 
// malum bir shart bajarilguncha yoki aniq bir marta kodni qayta qayta bajarish mumkin 






package main

import "fmt"


// func main() {
// 	for i := 0; i <5; i++ {
// 		fmt.Println("qiymat", i)
// 	}
// }

// bu yerda for lop har bir 5 dan kichik qiymatga yugurib chiqib bizga qiymatni korsatib chiqyapti 
// masalan 0 kicik 5 dan true keyngsi, 1 kicik 5 dan keyngsi, 2 kicik 5 dan keyngsi, va shu ravishda davom etadi 


func main() {
	arr := [3]string{"olma", "banan", "ananas"}
		for i, v := range arr {
			fmt.Printf("index: %d, value:%s \n", i, v)
		}
}



// bunda for loop har birdan yugurib chiqib index nechiligi va value nimaligini korsatb chiqyapt