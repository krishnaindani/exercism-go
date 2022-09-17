package lasagna

const (
	Noodles = "noodles"
	Sauce   = "sauce"
)

//PreparationTime ...
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

//Quantities ...
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == Noodles {
			noodles += 50
		}
		if v == Sauce {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

//AddSecretIngredient ...
func AddSecretIngredient(ingredientsFromFriend, ingredientsOfMine []string) {
	ingredientsOfMine[len(ingredientsOfMine)-1] = ingredientsFromFriend[len(ingredientsFromFriend)-1]
}

//ScaleRecipe ...
func ScaleRecipe(amountNeeded []float64, portions int) []float64 {
	desiredNumberOfPortions := make([]float64, len(amountNeeded))
	for i, v := range amountNeeded {
		desiredNumberOfPortions[i] = (v / float64(2)) * float64(portions)
	}
	return desiredNumberOfPortions
}
