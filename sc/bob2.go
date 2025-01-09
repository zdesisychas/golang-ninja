package main

import "fmt"

type Bob2 struct {
	name string
	numbers int
	items map[string]float64
	tip float64
}


func NewBob2(names string) Bob2 {
	b := Bob2 {
		name: names,
		numbers: 20,
		items: map[string]float64{"kofe": 6.9, "choy": 7.9, "moxito": 4.9},
	}
	return b
}


func(b *Bob2) format() string {
	fs := "Bob breakdown\n"
	var barchasi float64 = 0


		for k, v := range b.items {

			 fs += fmt.Sprintf("%v %v\n", k+":", v)
			 barchasi += v 
			}

				 fs += fmt.Sprintf("%v %0.3f\n", "tip:", b.tip)

				 fs += fmt.Sprintf("%v %0.3f", "barchasi:",  barchasi+ b.tip)

					return fs
}

func (b *Bob2) UpdateTip(tip float64) {
	b.tip = tip
}

func (b *Bob2) UpdateItems(name string, price float64){
	b.items[name] = price
}