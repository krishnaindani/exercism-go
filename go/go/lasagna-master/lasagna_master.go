package lasagna

//PreparationTime ...
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
