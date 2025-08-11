package lasagnapro

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	// This function should return the preparation time in minutes
	// based on the number of layers and the time per layer.
	// For now, it just panics to indicate it's not implemented.

	//by default, if no timePerLayer is provided, it should be 2 minutes
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(listOfFriend, ingredientList []string) {
	if len(listOfFriend) == 0 || len(ingredientList) == 0 {
		return
	}
	// The secret ingredient is the last one in the friend's list
	secretIngredient := listOfFriend[len(listOfFriend)-1]
	ingredientList[len(ingredientList)-1] = secretIngredient
}

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
