package main

import (
	"testing"
)

func TestProcessTruck(t *testing.T) {

	t.Run("Diesel Truck Load/Unload Tes", func(t *testing.T) {
		dt := &DieselTruck{id: "DT1", cargo: 10, plates: "ABC-123"}
		err := process(dt)

		if err != nil {
			t.Fatalf("process returned error for DieselTruck: %v", err)
		}
		if dt.cargo != 0 {
			t.Errorf("DieselTruck cargo expected 0, got %d", dt.cargo)
		}
	})

	t.Run("Electric Truck Load/Unload Test", func(t *testing.T) {
		et := &ElectricTruck{id: "ET1", cargo: 10, plates: "ABC-123", battery: 100.0}
		err := process(et)

		if err != nil {
			t.Fatalf("process returned error for ElectricTruck: %v", err)
		}
		if et.cargo != 0 {
			t.Errorf("ElectricTruck cargo expected 0, got %d", et.cargo)
		}
		if et.battery != 97.0 {
			t.Errorf("ElectricTruck cargo expected 97.0, got %f", et.battery)
		}
	})
}
