package immute

import "fmt"

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
