// Package factorial separate the functionality of factorial calculation into its own package under `internal` folder
// to keep the implementation details private to the application
package factorial

// Calculate computes the factorial of n using recursion.
func Calculate(n int) int {
	if n == 0 {
		return 1
	}
	return n * Calculate(n-1)
}
