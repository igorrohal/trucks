package main

import (
	"fmt"
	"log"
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
	fmt.Printf("Loading truck %+v \n", t)
	t.cargo += 2
	// return fmt.Errorf("Error while loading truck: %+v \n", t)
	return nil
}

func (t *DieselTruck) unloadCargo() error {
	fmt.Printf("Unloading truck %+v \n", t)
	t.cargo -= 2
	// return fmt.Errorf("Error while unloading truck: %+v \n", t)
	return nil
}

func (e *ElectricTruck) loadCargo() error {
	fmt.Printf("Loading truck %+v \n", e)
	e.cargo += 2
	e.battery -= 2
	// return fmt.Errorf("Error while loading truck: %+v \n", e)
	return nil
}

func (e *ElectricTruck) unloadCargo() error {
	fmt.Printf("Unloading truck %+v \n", e)
	e.cargo -= 2
	e.battery -= 1
	// return fmt.Errorf("Error while unloading truck: %+v \n", e)
	return nil
}

// loads and unloads a Truck
func process(t Truck) error {
	fmt.Printf("Processing truck %+v \n", t)

	if err := t.unloadCargo(); err != nil {
		return fmt.Errorf("Couldn't unload truck: %w", err)
	}

	if err := t.loadCargo(); err != nil {
		return fmt.Errorf("Couldn't load truck: %w", err)
	}

	return nil
}

func main() {
	dt := &DieselTruck{id: "Truck 1", cargo: 0, plates: "WN-710EM"}
	et := &ElectricTruck{id: "Truck 1", cargo: 0, plates: "WN-710EM", battery: 100}

	if err := process(dt); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	if err := process(et); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}
}
