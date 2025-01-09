// ****************************************Functions*******************************************


// --------------- len()---------------------


len() bu funksiya matn, array , slice , map , yoki kanalning uzunligni aniqLaydi 
 masalan str := "salom" bu sozni uzunligi 5 ta 
 uzunlik := len(str) // uzunlik =5 


//  ---------------- cap() ---------------


cap() funksyasi slice, array, channel , ni sigimini aniqlaydi
 slice := make([]int, 3, 5) bu yerda uznlik 3 , lekin sigim 5



// --------------- append() --------------------
 

append() funksyasi (slice) ga yangi elementlar qoshish uchun ishlatiladi agrar sigim yetalri blmasa go avtamatik ravishd akatta sigimga
ega kesma yaratadi 
slice := []in{2, 2, 3}
slice = append(slice, 4)
fmt.Println(slice)      bu yerda kesma oxirida 4 qoshildi Natija : [1 2 3 4]