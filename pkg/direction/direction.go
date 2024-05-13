package direction

import (
	"math/rand"
)

const NoDirection = -1

func GetRandomDirection(lastDirection, speed int) int {
    var directions []int

    // Calcul du nombre total de directions possibles
    totalDirections := 8 * speed

    // Ajout de toutes les directions possibles
    for d := 0; d <= totalDirections; d++ {
        directions = append(directions, d)
    }

    // Si lastDirection est non nulle, exclure la direction opposée
    if lastDirection >= 0 {
        // Calcul de la direction opposée
        opposite := (lastDirection + (totalDirections / 2)) % totalDirections

        // Retrait de la direction opposée des directions possibles
        for i := len(directions) - 1; i >= 0; i-- {
            if directions[i] == opposite {
                directions = append(directions[:i], directions[i+1:]...)
            }
        }
    }

    // Si aucune direction n'est possible, retourner NoDirection
    if len(directions) == 0 {
        return NoDirection
    }

    // Retourner une direction aléatoire parmi les directions possibles
    return directions[rand.Intn(len(directions))]
}


// Juste a retourner direction au lieu du Stwich
func GetNewDirection(lastDirection, speed int) int {
	direction := GetRandomDirection(lastDirection, speed)
	return direction
}