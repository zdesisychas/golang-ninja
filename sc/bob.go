//  ----------------------------------  Struct --------------------------------------------

//  pasda struct yaratilgan misol tarifasida

package main

import (
	"fmt"
)

type Bob struct {
	name string
	age int
    weigh float64            
	hight float64
	items map[string]float64
	tip float64

}

func Newbob(names string) Bob {
	a := Bob{
		name: names,
		age: 19,
		weigh: 50.3,
		hight: 176.5,
		items: map[string]float64{"kofe": 5.99, "choy": 3.99},
	} 
	
	// fmt.Printf("Mening ismim %s va mening yoshim %d da, va vesim %.1f boyim esa %.1f \n",a.name, a.age, a.weigh, a.hight)
	return a
}



// formatni hisoblash
func (a *Bob) format() string {
	fs := "Bob breakdown\n"
	var hammasi float64 = 0

		//List items
		for k, v := range a.items {
			fs += fmt.Sprintf("%v %v\n", k+ ":", v) 
			hammasi += v
		}

		fs += fmt.Sprintf("%v %0.2f\n", "tip:", a.tip)

			//total
			fs += fmt.Sprintf("%v %0.2f", "total:", hammasi+a.tip)
				return fs
}

//updateTip
func (a *Bob) UpdateTip(tip float64) {
	a.tip = tip
}

//updateItems
func (a *Bob) updateItems(name string, price float64) {
	a.items[name] =price
}

