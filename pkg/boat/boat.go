package boat

import (
	"fmt"
	"redis.hive/pkg/direction"
	"github.com/rs/xid"	
	"math"
	"math/rand"
	"context"
    "github.com/redis/go-redis/v9"
	"sync"
)

type Boat struct {
	ID        		string
	Latitude 		int
	Longitude		int
	Speed 			int
	Cap 			int 
	LastDirection 	int
  }

  func generateRandomSpeed() int {
    // Créez un tableau de vitesses avec une distribution préférentielle pour les valeurs spécifiées
    speeds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 24, 26, 28, 30}

    // Générez un index aléatoire pour sélectionner une vitesse dans le tableau
    randomIndex := rand.Intn(len(speeds))

    // Renvoie la vitesse correspondant à l'index aléatoire sélectionné
    return speeds[randomIndex]
}

func generateRandomPosition() (int, int) {
	min := -5000.0
	max := 5000.0
	latitude := int(math.Round(min + rand.Float64()*(max-min)))
	longitude := int(math.Round(min + rand.Float64()*(max-min)))
	return latitude, longitude
  }

func idGenerator() string {
	id := xid.New()
    idStr := id.String()
	return idStr
}

func Generator() (*Boat, error) {
	latitude, longitude := generateRandomPosition()
	boat := &Boat {
		ID: idGenerator(),
		Latitude: latitude,
		Longitude: longitude,
		Cap: direction.NoDirection,
		LastDirection: direction.NoDirection,
		Speed: generateRandomSpeed(),
		}
	return boat, nil
}

func CreateBoatsParallel(ctx context.Context, client *redis.Client, totalBoats int) ([]Boat, error) {
	var boatsArray []Boat
	var mu sync.Mutex // Mutex pour la synchronisation

    var wg sync.WaitGroup
    wg.Add(totalBoats) // Ajouter le nombre total de goroutines à attendre

	for i := 0; i < totalBoats; i++ {
		go func(){
			defer wg.Done() // Indiquer que la goroutine est terminée
			boat, err := Generator()
			if err != nil{
				fmt.Print("Error: when creating a boat:", err)
				return
			}
			fmt.Println("In CREATEBOATS:\n\tValue of boat: ", boat)
			// Verrouiller pour éviter les accès simultanés à la variable boatsArray
			mu.Lock()
			defer mu.Unlock()
			ok := SaveBoatToRedis(ctx, client, *boat)
			if ok != nil {
				fmt.Printf("Error when saving boat '%s'\n%s", boat.ID, ok)
				return
			}
			boatsArray = append(boatsArray, *boat)
			fmt.Printf("\tBoat'%s' has been saved\n", boat.ID)
		}()
	}
	wg.Wait()
	return boatsArray, nil	
}