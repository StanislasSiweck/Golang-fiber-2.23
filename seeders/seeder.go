package seeders

import (
	"fmt"
	"time"
)

func Seed() {
	start := time.Now()
	fmt.Println("Seeding users")
	userSeeder()
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("Seeding roles")
	roleSeeder()
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("Seeding user roles")
	userRolesSeeder()
	fmt.Println(time.Since(start))
}
