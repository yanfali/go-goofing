package main

import (
	"fmt"

	"github.com/yanfali/go-goofing/immutablesortof/immute"
)

type ImmutableThing struct {
	name string
	age  int
}

func New(name string, age int) *ImmutableThing {
	return &ImmutableThing{name: name, age: age}
}

func Copy(other *ImmutableThing) *ImmutableThing {
	return &ImmutableThing{name: other.name, age: other.age}
}

// noodling on the value of non pointer method
func (my ImmutableThing) Age() int {
	return my.age
}

// is there a good use case for them that I can't see?
func (my ImmutableThing) Name() string {
	return my.name
}

// If you're going to do this anyway then the data is read-only by default
// maybe this adds an extra layer of inaccessibility? It's certainly less efficient
func (my ImmutableThing) String() string {
	return fmt.Sprintf("I am %s aged %d", my.name, my.age)
}

func main() {
	thing := New("Harry Potter", 13)
	fmt.Println(thing)
	otherthing := Copy(thing)
	fmt.Println(otherthing)
	thing.name = "blah"
	fmt.Println(thing) // see this isn't that good same package

	immute1 := immute.New("Ron Weasley", 13)
	fmt.Println(immute1)
	//immute1.name = "test" // ah this is now illegal
	immute2 := immute.Copy(immute1)
	fmt.Println(immute2)
}
