package main

import (
	"fmt"
)


type Bob1 struct {
	name string
	number int 
	items map[string]float64
	tip float64
}

func NewBob1(names string) Bob1 {
	s := Bob1 {
	name: names,
	number: 30,
	items: map[string]float64{"kofe": 4.5, "choy": 5.5,},
}
return s
}

func (s *Bob1) format() string {
	fs := "Hisob kitob:\n"
	var barchasi float64 = 0

	for k, v := range s.items {

		fs += fmt.Sprintf("%v $%0.3f\n", k+ ":",v ,)
		barchasi += v
	}

		fs += fmt.Sprintf("%v %v\n", "tip:", s.tip)

		fs += fmt.Sprintf("%v $%0.3f\n", "barchasi:" , barchasi)

		return fs

}


func (s *Bob1) UpdateTip(tip float64) {
	s.tip = tip
}

func (s *Bob1) updateItems(name string, price float64){
	s.items[name] = price
}