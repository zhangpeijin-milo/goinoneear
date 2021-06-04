package designpattern

import (
	"errors"
	"fmt"
)

type coffee interface {
	taste() string
	cost() int8
}

type cappuccino struct {}

type latte struct {}

func (c cappuccino) taste() string {
	return "cappuccino"
}

func (c cappuccino) cost() int8 {
	fmt.Println("cost $12" )
	return 12
}

func (c latte) taste() string {
	return "latte"
}

func (c latte) cost() int8 {
	fmt.Println("cost $8" )
	return 8
}

func createCoffee(name string) (coffee,error){
	switch name {
		case "C" :
			c := cappuccino{}
			defer c.cost()
			return c , nil
		case "L" :
			l := latte{}
			defer l.cost()
			return l , nil
		default:
			return nil , errors.New(name + " is not supported yet")
	}
}

func FactoryExample() {
	list := []string{"C","L","EEE"}

	for _, c := range list {
		coffee, err := createCoffee(c)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("got " + coffee.taste())
		}
	}
}