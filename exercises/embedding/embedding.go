package main

import "fmt"

type Bytes int
type Celsius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celsius
}

type MemoryUsage struct {
	amount []Bytes
}

// functions to calculate averages for every dashboard component
func (b *BandwidthUsage) AverageBandwidthUsage() int {
	sum := 0
	for _, amount := range b.amount {
		sum += int(amount)
	}
	return sum / len(b.amount)
}

func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for _, temp := range c.temp {
		sum += int(temp)
	}
	return sum / len(c.temp)
}

func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for _, amount := range m.amount {
		sum += int(amount)
	}
	return sum / len(m.amount)
}

// dashboard struct
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celsius{50, 51, 32, 59, 44}}
	memory := MemoryUsage{[]Bytes{250000, 10000, 230000, 89000, 900000}}

	dashboard := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	fmt.Printf("Average bandwidth usage: %v\n", dashboard.AverageBandwidthUsage())
	fmt.Printf("Average CPU temperature: %v\n", dashboard.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", dashboard.AverageMemoryUsage())
}
