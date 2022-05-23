package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
	var time int
	if avg == 0 {
		time = len(layers) * 2
	} else {
		time = len(layers) * avg
	}
	return time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, v := range layers {
		if v == "noodles" {
			noodles++
		}
		if v == "sauce" {
			sauce++
		}
	}
	return (noodles * 50), (sauce * 0.2)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amount []float64, portions int) []float64 {
	var scaleAmount = make([]float64, len(amount))
	var scale float64 = float64(portions)
	if portions == 2 && len(amount) != 0 {
		return amount
	} else if len(amount) != 0 {
		for i := 0; i < len(amount); i++ {
			scaleAmount[i] = amount[i] * (scale / 2)
		}
		return scaleAmount
	} else {
		return scaleAmount
	}
}
