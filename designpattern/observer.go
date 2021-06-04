package designpattern

import "fmt"

type Subject struct {
	name string
	obs  []*Observer
}

func (s *Subject) Add(ob *Observer) {
	s.obs = append(s.obs, ob)
}

func (s *Subject) Notify(msg string) {
	for _,ob := range s.obs {
		(*ob).Receive(msg, s.name)
	}
}

type Observer interface {
	Receive(msg string, sub string)
}

type Observer1 struct {
	obName string
}

func (ob Observer1) Receive(msg string, sub string) {
	fmt.Println(ob.obName + " has received " + msg + " from " + sub)
}

type Observer2 struct {
	obName string
}

func (ob *Observer2) Receive(msg string, sub string) {
	fmt.Println(ob.obName + " has received " + msg + " from " + sub)
}

func ObserverExample()  {
	subject := Subject{name: "zpj", obs: make([]*Observer, 0, 2)}

	var ob1 Observer = Observer1{obName: "ob1"}
	subject.Add(&ob1)

	// 因为实现Observer的是Observer2的指针类型，所以要 &
	var ob2 Observer = &Observer2{obName: "ob2"}
	subject.Add(&ob2)

	subject.Notify("hello")
}