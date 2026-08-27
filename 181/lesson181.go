package main

import (
	"errors"
	"fmt"
	"log"
)

type Engine struct {
	Started    bool
	HorsePower int
}

type Car struct {
	Engine Engine
	Model  string
}

func (c *Car) Start() error {
	if err := c.Engine.Start(); err != nil {
		return fmt.Errorf("engine start: %w", err)
	}
	fmt.Println("Включаются фары. Двигатель рычит!")
	return nil
}

func (e *Engine) Start() error {
	if e.Started {
		return errors.New("engine is already started")
	}
	e.Started = true
	return nil
}

func (c Car) Drive() error {
	if !c.Engine.Started {
		return errors.New("off engine, can't start")
	}
	fmt.Println("Едем в соседнее село.")
	return nil
}

func main() {
	car := Car{
		Engine: Engine{
			Started:    false,
			HorsePower: 150,
		},
		Model: "Toyota",
	}
	if err := car.Start(); err != nil {
		log.Fatalf("start car: %v", err)
	}

	fmt.Printf("Машина запущена.\nМодель: %s.\nЛошадок: %d.\n", car.Model, car.Engine.HorsePower)
	fmt.Printf("%+v\n", car)

	if err := car.Drive(); err != nil {
		log.Fatalf("drive car: %v", err)
	}

	if err := car.Start(); err != nil {
		log.Fatalf("start car: %v", err)
	}
}
