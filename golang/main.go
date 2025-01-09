package main

import (
	"fmt"
)




func main(){
	message := make([]string, 10)
	fmt.Println(len(message))
	fmt.Println(cap(message))
	message = append(message, "11")
	fmt.Println(len(message))
	fmt.Println(cap(message))
}





// func main() {
// 	phonebook := map[int]string{
// 		98329873987: "bobur",
// 		3768763286: "xasan",
// 		8937498237:"waxboz",
// 	}

// 	fmt.Println(phonebook[98329873987])
// 	phonebook[98329873987] = "nari"
// 	fmt.Println(phonebook)
	
// 	for k, v  := range phonebook {
// 	fmt.Println(k, "-", v)
// }


// }











// func main () {
// 	menu := map[string]float64 {
// 		"tort" : 20.55,
// 		"bulochka" : 10.99,
// 		"choy" : 5.11,
// 	}





// 	fmt.Println(menu)
// }





// var score = 95.5
 
// func main(){
// 	sayHello("bobur")


// 	for _, v := range points {
// 		fmt.Println(v)
// 	}

// 	showScore()

// }


// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}
// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}
	
// 	return initials[0], "_"	
// }
// func main() {
// 	fn1 , sn1 :=getInitials("bobur muhammadiyev")
// 	fmt.Println(fn1, sn1)
// 	fn2, sn2 := getInitials("xasan davronbekov")
// 	fmt.Println(fn2, sn2)
// 	fn3, sn3 := getInitials("   turayev")
// 	fmt.Println(fn3, sn3)
// }






// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }
// func sayBye(n string)  {
// 	fmt.Printf("Goodbye %v \n", n)
	
// }

// func qtname(n[]string, f func(string)){
// 	for _, v := range n {
// 		f(v)
// 	}

// }

// func main(){

// 	qtname([]string{"azi", "sevinc", "volida"}, sayGreeting)
// 	qtname([]string{"bobur", "xasan", "waxboz"}, sayBye)
// 	// sayGreeting("bobur")
// 	// sayGreeting("xasan")
// 	// sayBye("waxboz")
	
// }





// names := []string{"bobur", "xasan" , "waxboz", "narimon"}

// // sort.Strings(names)
// // fmt.Println(names)

// for index, value := range names {
// 	if index <= 1 {
// 		fmt.Println("salom xsuh kelibsan", value)
// 		continue
// 	}
// 	if index >1 {
// 		fmt.Println("nega keldin", value)
// 	}
		
// 	}


// ages := []int{32, 67, 87, 36, 1, 78}

// sort.Ints(ages)
// fmt.Println(ages)


// index:= sort.SearchInts(ages , 67)
// fmt.Println(index)



// names := []string {"bobur va", "xasan" , "waxbzo", "narimon"}

// sort.Strings(names)
// fmt.Println(names)

// indexn := sort.SearchStrings(names , "xasan")
// fmt.Println(indexn)


// 


// func main () {
// 	var ism string
// 	var yosh uint8


// 	fmt.Scanln(&ism)
// 	fmt.Scanln(&yosh)


// 	fmt.Println(ism,yosh,"yoshda")
// }







// func main () {

	
// 	fmt.Println(prediction("md"))
	
// }


//  func prediction(dayOfweek string) (string , error) {
// 	switch dayOfweek {
// 	 case "mn":
// 		return "Monday is first day a week", nil
// 	 case "sd":
// 		return "Saturday is sunny day" , nil
// 	 case "wd":
// 		return "Wednesday is hard day " , nil
// 	 case "td":
// 		return "Thursday is lovely day" , nil
// 	 case "fd":
// 		return "The best day is Friday" , nil
// 	 default :
// 	    return "error", errors.New("invalid day of the week")

// 	 }
// }











// func main () {
// 	someInt , message , err := kriditMarkeeet(12)
// 	if err != nil {
// 	 fmt.Println(err)
// 	 return
// 	}
// 	fmt.Println(message)
// 	fmt.Println(someInt)
// }





// func kriditMarkeeet (age int) (int , string , error) {
// if age >=19 && age < 50 {
// 	return 200 , "success" , nil
// } else if age > 50 {
// 	return 501 , "unsuccsesfull" , errors.New("Sizning yoshingiz to'g'ri kelmaydi!")
// } else {
// 	return 501, "unsuccsesfull" , errors.New("Yoshingiz kichik")
// }
// }






// func main () {
// 	someInt, message, err := smokeShop(12)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(someInt)
// 	fmt.Println(message)
// }


// func smokeShop (age int)(int , string , error) {
// 	if age >=21 && age < 50 {
// 		return 200 , "success", nil 
// 	}else if age > 51 {
// 		return 401 , "it's bad for old humans health!" , errors.New("It's bad for your health")
// 	}else { 
// 		return 500 , "too young for smoke" , errors.New("go to gym don't smoke!")
// 	}
// }


// func main() {
// 	// message := "Hello Bobur"

	

// 	// message = "bobur qonday"

// 	// var message int


// 	// var message string



// 	// var message int
// 	// message = 12
// 	// var b bool 

// 	// b = true


// 	//       
	
	

// 	fmt.Println(message)
// 	// fmt.Println(b)
// }





// func main (){
// 	message, err := enterTheClub(12)
// 	if  err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(message) 

// }

// func enterTheClub (age int )( string , error) {

// 	if age >= 18 && age < 45 {
// 		return "Xush kelibsiz!" ,  nil
// 	} else if age >= 45 && age < 65 {
// 		return "Sizga anniq yoqdimi?" , nil
// 	} else if age >= 65 {
// 		return "Sizga bu sayt togri kelmaydi!" , errors.New("you are too old!!")
// 	} else {
// 		return "Hali yoshsan!" ,  errors.New("Sani yoshin tori kemid tez chiq!!!")

// 	}

// }

//---------------------------IF------------------------------------

