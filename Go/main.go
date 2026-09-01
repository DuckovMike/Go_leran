package main

import (
	"Go/models"
	"fmt"
)

func main() {
	u1 := models.CreateUser(1, "Misha", "Ingenier", 0.5)
	u1.CreateDefaultAppearence(2026, 1, 1)

	fmt.Println(u1)
}
