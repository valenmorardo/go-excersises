package main

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        return len(layers)*2
    }
    return len(layers)*time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var gOfNoodles int
    var lOfSauce float64
    for _, v := range layers {
        if v == "sauce" {
            lOfSauce += 0.2
        }
        if v == "noodles" {
            gOfNoodles += 50
        }
    }
    return gOfNoodles, lOfSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(listOfMyFriend []string, myList []string) {
    myList[len(myList)-1] = listOfMyFriend[len(listOfMyFriend)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsFor2Portions []float64, portionsWantedToCook int) []float64{
    amountForPortionsFinal := make([]float64, len(amountsFor2Portions))
    
    for i:=0;i<len(amountForPortionsFinal);i++ {
		amountForPortionsFinal[i] = amountsFor2Portions[i]*(float64(portionsWantedToCook)/2.0)
    }
    return amountForPortionsFinal
}

