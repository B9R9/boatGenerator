package boat

import (
	"fmt"
    "context"
    "time"
    "github.com/redis/go-redis/v9"
	"redis.hive/pkg/direction"
	"sync"
    "bufio"
    "os"
    "encoding/json"
    "math"
)

type BoatUpdater struct {
    ctx         context.Context
    client      *redis.Client
    updateRate  time.Duration
	boatsArr		[]Boat
    redisKey        string
}

// NewBoatUpdater crée une nouvelle instance de BoatUpdater avec les paramètres donnés.
func NewBoatUpdater(ctx context.Context, client *redis.Client, updateRate time.Duration, boatsArray []Boat) *BoatUpdater {
    return &BoatUpdater{
        ctx:        ctx,
        client:     client,
        updateRate: updateRate,
        redisKey: "boats",
        boatsArr: boatsArray,
    }
}

func (bu *BoatUpdater) LoopUpdatePosition() {
    ticker := time.NewTicker(bu.updateRate)
    defer ticker.Stop()
	var wg sync.WaitGroup
    
    exit := make(chan struct{})

    go func() {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Println("You can exit by taping quit")
        for scanner.Scan() {
            if scanner.Text() == "quit" {
                // Envoyer un signal pour terminer le programme
                close(exit)
                bu.ctx.Done()
                return
            }
        }
    }()
        
    for {                
        select {
            case <-bu.ctx.Done():
                return
            case <-exit:
                // Annuler le contexte pour signaler l'arrêt
                bu.ctx.Done()
                return
            case <-ticker.C:
                for i := range bu.boatsArr {
                    wg.Add(1)
                    go bu.updatePositions(&bu.boatsArr[i], &wg)
                    fmt.Println("**********************")
                    }
                wg.Wait()
        }
    }
}
        
func (bu *BoatUpdater) updatePositions(boat *Boat, wg *sync.WaitGroup) {
    defer wg.Done()
    newCap := direction.GetNewDirection(boat.LastDirection, boat.Speed)
    err := UpdateBoatPosition(bu.ctx, bu.client, boat, newCap)
    if err != nil {
        fmt.Printf("Erreur lors de la mise à jour de la position du bateau %s: %v\n", boat.ID, err)
        return
    }
}

func UpdateBoatPosition(ctx context.Context, client *redis.Client, boat *Boat, newCap int) error {
    // Vérifier si le bateau existe dans le hachage
    exists, err := client.HExists(ctx, "boats", boat.ID).Result()
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("Boat with ID '%s' not found in Redis", boat.ID)
    }
    fmt.Println(boat)
    lastDirection := boat.Cap
    // fmt.Println("LastDirection: ", lastDirection)
    latitude, longitude := GetNewPosition(newCap, boat.Speed ,boat.Latitude, boat.Longitude)
  
    // Mettre à jour la latitude et la longitude du bateau dans le hachage
    fields := map[string]interface{}{
        "Latitude":  latitude,
        "Longitude": longitude,
        "LastDirection": lastDirection,
        "Cap": newCap,
        "Speed": boat.Speed,
    }

    fieldsJson, err := json.Marshal(fields)
    if err != nil {
        return err
    }

    fmt.Println("fieldsJson:\n", string(fieldsJson))

    boat.Latitude = latitude
    boat.Longitude = longitude
    boat.LastDirection = lastDirection
    boat.Cap = newCap

    fmt.Println(boat)

    err = client.HSet(ctx, "boats", boat.ID, fieldsJson).Err()
    if err != nil {
        return err
    }
    return nil
  }

func GetNewPosition( direction int, speed, latitude, longitude int) (int, int){
    newLatitude := latitude
    newLongitude := longitude

        // La vitesse détermine le pas de déplacement
    // Par exemple, si la vitesse est de 1, chaque pas déplace d'une unité en latitude et longitude
    // Si la vitesse est de 2, chaque pas déplace de deux unités, etc.

    // Le nombre total de directions possibles est 8 fois la vitesse
    totalDirections := 8 * speed

    // Le pas de latitude et de longitude pour chaque direction
    latStep := float64(180) / float64(totalDirections)
    // longStep := float64(360) / float64(totalDirections)

    // Calcul de l'angle en radians pour la direction spécifiée
    angle := float64(direction) * latStep * math.Pi /180

    // Mise à jour des coordonnées en fonction de l'angle
    newLatitude += int(math.Sin(angle) * float64(speed))
    newLongitude += int(math.Cos(angle) * float64(speed))

    // Verifier si les valeurs ne sortent pas de la map
    if newLatitude > 5000 {
        newLatitude = -5000 + (newLatitude - 5000)
    } else if newLatitude < -5000 {
        newLatitude = 5000 - (-5000 - newLatitude)
    }
    if newLongitude > 5000 {
        newLongitude = -5000 + (newLongitude - 5000)
    } else if newLongitude < -5000 {
        newLongitude = 5000 - (-5000 - newLongitude)
    }
    
    return newLatitude, newLongitude
  }
