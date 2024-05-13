package options
import (
	"flag"
	"strconv"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
	"context"
	"encoding/json"
)

type Options struct {
	Host string
	Port int
	RefreshRate time.Duration
	NumberOfBoats int
}

func SaveSettingToRedis(ctx context.Context, client *redis.Client, setting Options) error {
	settingMap := map[string]interface{}{
	  "RefreshRate": int64(setting.RefreshRate),
	  "Port": setting.Port,
	  "Map": 0,
	  "Host": setting.Host,
	  "TotalBoats": setting.NumberOfBoats,
	}
	settingJson, err := json.Marshal(settingMap)
	if err != nil {
	  return err
	} 
  
	// Utiliser HSET pour ajouter le bateau au hachage
	err = client.HSet(ctx, "settings", "settings", settingJson).Err()
	if err != nil {
		return err
	}
	  return nil
  }

func ParseCommandLineOptions() (*Options, error) {
	hFlag := flag.String("h", "localhost", "Host du Redis")
	pFlag := flag.Int("p", 6379, "Port pour Redis")
	rFlag := flag.Int("r", 1, "Refresh rate in second")

	flag.Parse()

	if flag.NArg() == 0 {
		return nil, fmt.Errorf("Print usage")
	}

	remainingArgs := flag.Arg(0)
	tBoat, err := strconv.ParseInt(remainingArgs, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Print Usage")
	}

	if tBoat <= 0 {
		return nil, fmt.Errorf(("Print usage"))
	}

	options:= &Options{
		Host: *hFlag,
		Port: *pFlag,
		NumberOfBoats: int(tBoat),
	}
	options.RefreshRate = time.Duration(*rFlag) * time.Second
	
	return options, nil
}