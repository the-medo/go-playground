package main

import "testing"

func createPlayer() Player {
	return Player{
		name:      "Player",
		health:    1000,
		maxHealth: 1000,
		energy:    100,
		maxEnergy: 100,
	}
}

func TestIncreaseHealth(t *testing.T) {
	p := createPlayer()

	success := p.increaseHealth(100)
	if !success {
		t.Errorf("Positive number returned false")
	}

	if p.health > p.maxHealth {
		t.Errorf("Health is over maximum")
	}

	success = p.increaseHealth(-100)
	if success {
		t.Errorf("Negative number returned true")
	}
}

func TestDecreaseHealth(t *testing.T) {
	p := createPlayer()

	success := p.decreaseHealth(100)
	if !success {
		t.Errorf("Positive number returned false")
	}

	p.decreaseHealth(20000)
	if p.health < 0 {
		t.Errorf("Health is less than 0")
	}

	success = p.decreaseHealth(-100)
	if success {
		t.Errorf("Negative number returned true")
	}
}

func TestIncreaseEnergy(t *testing.T) {
	p := createPlayer()

	success := p.increaseEnergy(100)
	if !success {
		t.Errorf("Positive number returned false")
	}

	if p.energy > p.maxEnergy {
		t.Errorf("Energy is over maximum")
	}

	success = p.increaseEnergy(-100)
	if success {
		t.Errorf("Negative number returned true")
	}
}
