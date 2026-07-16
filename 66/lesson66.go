package main

import (
	"fmt"
	"os"
)

func main() {
	movePirate(false)
	movePirate(true)
	movePirate(false)
	movePirate(false)
	movePirate(true)
	movePirate(false)
	movePirate(false)
	movePirate(false)
	movePirate(false)
	movePirate(false)
}

var (
	move  int
	life  int = 3
	score int
)

func movePirate(isTrap bool) {
	if isTrap {
		life--
		move++
		score++
		fmt.Println("Пират переместился на плиту", move)
		if life == 0 {
			fmt.Println("Пират убит")
			os.Exit(0)
		} else {
			fmt.Println("Пират ранен")
		}
	}
	if !isTrap {
		move++
		score++
		fmt.Println("Пират переместился на плиту", move)
	}
	if score == 10 && life >= 1 {
		fmt.Println("Пират преодолел все ловушки")
		return
	}
}
