package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"redis.hive/pkg/boat"
	"redis.hive/pkg/options"
)

func main() {
	settings, err := options.ParseCommandLineOptions()
	if err != nil {
		fmt.Println("Erreur lors de l'analyse des options de la ligne de commande:", err)
		return
	}

	fmt.Println("HOST: ", settings.Host)
	fmt.Println("PORT: ", settings.Port)
	fmt.Println("R: ", settings.RefreshRate)
	fmt.Println("Boat: ", settings.NumberOfBoats)

	redisAddr := settings.Host + ":" + fmt.Sprint(settings.Port)
	fmt.Println("Result concact: ", redisAddr)

	// Initialisation du client Redis
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // aucun mot de passe
		DB:       0,  // base de données par défaut
	})
	defer client.Close()

	err = client.FlushDB(ctx).Err()
    if err != nil {
        fmt.Println("Could not clean db")
    }
    fmt.Println("DB cleaned")

	notOk := options.SaveSettingToRedis(ctx, client, *settings)
	if notOk != nil {
		fmt.Println("Error when saving settings in redis")
		return
	}

	boatsArray, err := boat.CreateBoatsParallel(ctx, client, settings.NumberOfBoats)
	if err != nil {
		fmt.Println("Error when creating or saving boats to redis")
		return
	}

	bu := boat.NewBoatUpdater(context.Background(), client, settings.RefreshRate, boatsArray)
	bu.LoopUpdatePosition()

}
