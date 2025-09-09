package problems

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"nil prices", nil, 0},
		{"empty prices", []int{}, 0},
		{"single day", []int{5}, 0},
		{"two days decreasing", []int{7, 1}, 0},
		{"two days increasing", []int{1, 7}, 6},
		{"all increasing", []int{1, 2, 3, 4, 5}, 4},
		{"all decreasing", []int{5, 4, 3, 2, 1}, 0},
		{"peak in middle", []int{3, 2, 6, 5, 0, 3}, 4},    // buy at 2, sell at 6
		{"multiple peaks", []int{2, 1, 2, 1, 0, 3, 4}, 4}, // buy at 0, sell at 4
		{"duplicates", []int{1, 1, 1, 1}, 0},
		{"zeros and rise", []int{0, 0, 0, 1}, 1},
		{"large values", []int{100000000, 1, 100000000}, 99999999},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxProfit(tc.prices)
			if got != tc.want {
				t.Fatalf("MaxProfit(%v) = %d, want %d", tc.prices, got, tc.want)
			}
		})
	}
}
