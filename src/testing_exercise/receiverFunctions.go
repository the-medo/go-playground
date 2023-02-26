package main

import "fmt"

type Player struct {
	health    int
	maxHealth int
	energy    int
	maxEnergy int
	name      string
}

func (player *Player) increaseHealth(amount int) bool {
	if amount < 0 {
		fmt.Println("Can't increase by negative number!")
		return false
	}

	startingHealth := player.health

	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}

	fmt.Println("[", player.name, "] HP (", player.health-startingHealth, ") [", player.health, "/", player.maxHealth, "]")
	return true
}

func (player *Player) decreaseHealth(amount int) bool {
	if amount < 0 {
		fmt.Println("Can't decrease by negative number!")
		return false
	}

	startingHealth := player.health

	player.health -= amount
	if player.health <= 0 {
		player.health = 0
		fmt.Println("DEAD")
	}

	fmt.Println("[", player.name, "] HP (", player.health-startingHealth, ") [", player.health, "/", player.maxHealth, "]")
	return true
}

func (player *Player) increaseEnergy(amount int) bool {
	if amount < 0 {
		fmt.Println("Can't increase by negative number!")
		return false
	}

	startingEnergy := player.energy

	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}

	fmt.Println("[", player.name, "] EN (", player.energy-startingEnergy, ") [", player.energy, "/", player.maxEnergy, "]")
	return true
}

func (player *Player) decreaseEnergy(amount int) bool {
	if amount < 0 {
		fmt.Println("Can't decrease by negative number!")
		return false
	}

	startingEnergy := player.energy

	player.energy -= amount
	if player.energy <= 0 {
		player.energy = 0
		fmt.Println("EXHAUSTED")
	}

	fmt.Println("[", player.name, "] EN (", player.energy-startingEnergy, ") [", player.energy, "/", player.maxEnergy, "]")
	return true
}

func (player *Player) printStats() {
	fmt.Println("[", player.name, "] HP: ", player.health, "/", player.maxHealth, " ; EN: ", player.energy, "/", player.maxEnergy, "")
}

func main() {

	p := Player{
		name:      "Zed",
		health:    1000,
		maxHealth: 1000,
		energy:    100,
		maxEnergy: 100,
	}

	p.increaseHealth(100)
	p.decreaseEnergy(50)
	p.decreaseHealth(500)
	p.increaseEnergy(20)
	p.increaseHealth(200)
	p.increaseEnergy(200)
	p.increaseHealth(500)
	p.decreaseEnergy(90)
	p.decreaseEnergy(90)
	p.decreaseHealth(200)
	p.decreaseHealth(2000)
	p.printStats()

}
