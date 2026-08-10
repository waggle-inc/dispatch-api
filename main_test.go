package main

import "testing"

func TestAssignDeliveries(t *testing.T) {
	tests := []struct {
		name    string
		drivers int
		want    int
	}{
		{"three drivers", 3, 12},
		{"one driver", 1, 4},
		{"no drivers", 0, 0},
		{"negative drivers", -2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssignDeliveries(tt.drivers); got != tt.want {
				t.Errorf("AssignDeliveries(%d) = %d, want %d",
					tt.drivers, got, tt.want)
			}
		})
	}
}
