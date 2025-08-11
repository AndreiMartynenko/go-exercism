package lasagna

// PreparationTime calculates the total prep time.
// If timePerLayer == 0, default to 2 minutes per layer.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// Quantities computes the amounts needed based on layers.
// Noodles: +50g per "noodles"; Sauce: +0.2l per "sauce".
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

// AddSecretIngredient replaces the last item in your ingredientList
// with your friend's last item (their secret ingredient), in place.
func AddSecretIngredient(friendList, ingredientList []string) {
	if len(friendList) == 0 || len(ingredientList) == 0 {
		return
	}
	ingredientList[len(ingredientList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe returns a NEW slice with quantities scaled from a base of 2 portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	factor := float64(portions) / 2.0
	for i, q := range quantities {
		scaled[i] = q * factor
	}
	return scaled
}
