package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	AddToTruckCargo(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
	sync.RWMutex
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (m *truckManager) AddTruck(id string, cargo int) error {
	m.Lock()
	defer m.Unlock()

	if _, exists := m.trucks[id]; exists {
		return errors.New("truck already exists")
	}
	m.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (m *truckManager) GetTruck(id string) (Truck, error) {
	m.Lock()
	defer m.Unlock()

	if truck, exists := m.trucks[id]; exists {
		return *truck, nil
	}
	return Truck{}, ErrTruckNotFound
}

func (m *truckManager) RemoveTruck(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, exists := m.trucks[id]; exists {
		delete(m.trucks, id)
		return nil
	}
	return ErrTruckNotFound
}

func (m *truckManager) UpdateTruckCargo(id string, cargo int) error {
	m.Lock()
	defer m.Unlock()

	if truck, exists := m.trucks[id]; exists {
		truck.Cargo = cargo
		return nil
	}
	return ErrTruckNotFound
}

func (m *truckManager) AddToTruckCargo(id string) error {
	m.Lock()
	defer m.Unlock()

	if truck, exists := m.trucks[id]; exists {
		truck.Cargo += 1
		// log.Printf("Updated truck %s to cargo %d", truck.ID, truck.Cargo)
		return nil
	}
	return ErrTruckNotFound
}
