package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return time * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodle := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodle += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	secret := friendList[len(friendList)-1]
	myList = myList[0 : len(myList)-1]
	myList = append(myList, secret)
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	var newQuantities []float64
	for _, q := range quantities {
		ratio := (q / 2) * float64(scale)
		newQuantities = append(newQuantities, ratio)
	}
	return newQuantities
}
