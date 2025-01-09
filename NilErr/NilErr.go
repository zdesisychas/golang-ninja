// ********************************* nil and err **********************************

// nil va err ular xatolarni boshqarish va xotira bilan ishlashda keng qollaniladi 

// nil nima?
// go tilida hech qanday qiymatga ega bolmagan korsatkich 
// pointer (interfeys), map, channel yoki func uchun ishlatiladi "bosh", "aniqlanmagan" degan manoni bldradi


// error (err) interfeysi xatolarni ifodalash uchun go tilida har bir xato error ga mos bolishi kerak.

// agar xatolik bolmasa error qiymati nil boladi 


package main


import (
	"fmt"
	"errors"
)

func divide(a, b float64) (float64, error){
	if b == 0 {
		return 0, errors.New("bolinish 0 ga bolinmaydi")
	}

	return a / b, nil
}



func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Xato:" , err)
	} else {
		fmt.Println("Natija:" , result)
	}
}




