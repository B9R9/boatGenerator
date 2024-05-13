package boat

import (
    "fmt"
    "context"
    "github.com/redis/go-redis/v9"
    "encoding/json"
    // "errors"
)

func SaveBoatToRedis(ctx context.Context, client *redis.Client, boat Boat) error {
    // Convertir le bateau en map[string]interface{} pour le stockage dans le hachage
    boatMap := map[string]interface{}{
      "Latitude":       boat.Latitude,
      "Longitude":      boat.Longitude,
      "Speed":          boat.Speed,
      "Cap":            boat.Cap,
      "LastDirection":  boat.LastDirection,
  }

  boatJson, err := json.Marshal(boatMap)
  if err != nil {
    return err
  } 

  // Utiliser HSET pour ajouter le bateau au hachage
  err = client.HSet(ctx, "boats", boat.ID, boatJson).Err()
  if err != nil {
      return err
  }
    return nil
}

// Récupération d'un bateau
func GetBoatFromRedis(ctx context.Context, client *redis.Client, boatID string) (Boat, error) {
  // Récupérer les données du bateau à partir du hachage
  boatData, err := client.HGet(ctx, "boats", boatID).Result()
  if err != nil {
      return Boat{}, err
  }

  fmt.Println("Value boatData: ", boatData)

  var boatUnmarshal Boat
  err = json.Unmarshal([]byte(boatData), &boatUnmarshal)
  if err != nil {
    fmt.Println("Error when turn boatdata to boatJson in GetBoatFromRedis")
  }

//   Convertir les données du bateau en une structure Boat
  var boat Boat
  boat.Latitude = boatUnmarshal.Latitude
  boat.Longitude = boatUnmarshal.Longitude
  boat.Speed = boatUnmarshal.Speed
  boat.Cap = boatUnmarshal.Cap
  boat.LastDirection = boatUnmarshal.LastDirection

  return boat, nil
}

// Récupération de tous les  bateaux
func GetAllBoatsFromRedis(ctx context.Context, client *redis.Client, key string) (map[string]string, error) {
  // Récupérer les données du bateau à partir du hachage
  boatData, err := client.HGetAll(ctx, "boats").Result()
  if err != nil {
      return nil, err
  }
  return boatData, nil
}