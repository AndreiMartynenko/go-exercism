package lasagnapro

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int, timePerLayer int) int {
	// This function should return the preparation time in minutes
	// based on the number of layers and the time per layer.
	// For now, it just panics to indicate it's not implemented.

	//by default, if no timePerLayer is provided, it should be 2 minutes
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return layers * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities() {}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient() {}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe() {}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
