package main

import "fmt"

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func main() {
	cp := CopyParent(nil) // -> Parent{}

	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}

	cp = CopyParent(p)

	// при мутациях в копии "cp"
	// изначальная структура "p" не изменяется
	cp.Children[0] = Child{
		Name: "Gosha",
		Age:  30,
	}

	fmt.Println(p.Children)  // -> [{Andy 18}]
	fmt.Println(cp.Children) // -> [{Gosha 30}]
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}

	child := make([]Child, len(p.Children))
	copy(child, p.Children)

	return Parent{
		Name:     p.Name,
		Children: child,
	}
}
