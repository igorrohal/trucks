package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type contextKey string

var UserIDKey contextKey = "UserID"

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
	fmt.Printf("Loading %+v \n", t)
	t.cargo += 2
	return nil
}

func (t *DieselTruck) unloadCargo() error {
	fmt.Printf("Unloading %+v \n", t)
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) loadCargo() error {
	fmt.Printf("Loading %+v \n", e)
	e.cargo += 2
	e.battery -= 2.0
	return nil
}

func (e *ElectricTruck) unloadCargo() error {
	fmt.Printf("Unloading %+v \n", e)
	e.cargo = 0
	e.battery -= 1.0
	return nil
}

// loads and unloads a Truck
func process(ctx context.Context, t Truck) error {
	fmt.Printf("Processing truck %+v \n", t)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	delay := 1 * time.Second
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	if err := t.loadCargo(); err != nil {
		return fmt.Errorf("couldn't load truck: %w", err)
	}
	if err := t.unloadCargo(); err != nil {
		return fmt.Errorf("couldn't unload truck: %w", err)
	}

	return fmt.Errorf("couldn't process truck: %+v", t)
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 42)

	fmt.Printf("Starting truck processing... %v", ctx)

	trucks := []Truck{
		&DieselTruck{id: "Truck 1", cargo: 0, plates: "WN-710EM"},
		&ElectricTruck{id: "Truck 2", cargo: 0, plates: "WN-710EM", battery: 100},
		&DieselTruck{id: "Truck 3", cargo: 0, plates: "WN-710EM"},
		&ElectricTruck{id: "Truck 4", cargo: 0, plates: "WN-710EM", battery: 100},
	}

	var wg sync.WaitGroup
	errsChan := make(chan error, len(trucks))

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := process(ctx, t); err != nil {
				errsChan <- err
			}
			wg.Done()
		}(t)
	}

	wg.Wait()
	close(errsChan)

	for err := range errsChan {
		fmt.Printf("Error processing truck (%s)\n", err)
	}
}
