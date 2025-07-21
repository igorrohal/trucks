package main

import (
	"fmt"
	"sync"
	"time"
)

type DieselTruck struct {
	id     string
	cargo  int
	plates string
}

type ElectricTruck struct {
	id      string
	cargo   int
	plates  string
	battery float64
}

type Truck interface {
	unloadCargo() error
	loadCargo() error
}

func (t *DieselTruck) loadCargo() error {
	t.cargo += 2
	return nil
}

func (t *DieselTruck) unloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) loadCargo() error {
	e.cargo += 2
	e.battery -= 2.0
	return nil
}

func (e *ElectricTruck) unloadCargo() error {
	e.cargo = 0
	e.battery -= 1.0
	return nil
}

// loads and unloads a Truck
func process(t Truck) error {
	fmt.Printf("Processing truck %+v \n", t)

	time.Sleep(time.Second)

	if err := t.loadCargo(); err != nil {
		return fmt.Errorf("Couldn't load truck: %w", err)
	}
	if err := t.unloadCargo(); err != nil {
		return fmt.Errorf("Couldn't unload truck: %w", err)
	}

	return nil
}

func main() {
	trucks := []Truck{
		&DieselTruck{id: "Truck 1", cargo: 0, plates: "WN-710EM"},
		&ElectricTruck{id: "Truck 2", cargo: 0, plates: "WN-710EM", battery: 100},
		&DieselTruck{id: "Truck 3", cargo: 0, plates: "WN-710EM"},
		&ElectricTruck{id: "Truck 4", cargo: 0, plates: "WN-710EM", battery: 100},
	}

	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := process(t); err != nil {
				fmt.Printf("Error processing truck %s: %v\n", t, err)
			}

			wg.Done()
		}(t)
	}

	wg.Wait()
}
