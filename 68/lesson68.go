package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println(Attack())

	attackWithDamageBoost := DamageBoosterDecorator(Attack)
	fmt.Println(attackWithDamageBoost())

	attackWithCriticalHit := CriticalHitDecorator(Attack)
	fmt.Println(attackWithCriticalHit())

	attackWithSlowEffect := slowEffectDecorator(Attack)
	fmt.Println(attackWithSlowEffect())

	attackSuper := slowEffectDecorator(CriticalHitDecorator(DamageBoosterDecorator(Attack)))
	fmt.Println(attackSuper())
}

func Attack() string {
	return "Атака выполнена!"
}

func DamageBoosterDecorator(attackFunc func() string) func() string {
	return func() string {
		return "Вам улыбнулась удача, нанесение урона увеличено на 10%! " + attackFunc()
	}
}

func CriticalHitDecorator(attackFunc func() string) func() string {
	return func() string {
		if rand.IntN(100) <= 25 {
			return "Критический урон! Удар удвоен! " + attackFunc()
		}
		return attackFunc()
	}
}

func slowEffectDecorator(attackFunc func() string) func() string {
	return func() string {
		return attackFunc() + " Цель замедлена на 2 хода!"
	}
}
