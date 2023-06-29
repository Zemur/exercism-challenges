package lasagna

func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2
    }

    return len(layers) * avgPrepTime
}

func Quantities(layers []string) (int, float64) {
    var gramsOfNoodles int
    var litersOfSauce float64

    for _, e := range layers {
        if e == "sauce" {
            litersOfSauce += 0.2
        } else if e == "noodles" {
        	gramsOfNoodles += 50
        }
    }

    return gramsOfNoodles, litersOfSauce
}

func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numPortions int) []float64 {
    var scaledQuantities []float64
    for _, e := range quantities {
        // Ususal amount is for 2 portions, but we want the amount for one.
    	var amountPerPortion float64 = e/2
        scaledQuantities = append(scaledQuantities, float64(numPortions) * amountPerPortion)
    }

    return scaledQuantities
}