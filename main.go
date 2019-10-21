package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter (easy parsing by logstash or Splunk).
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	var cursor uint64
	modeMsg := "[PROD MODE]"

	server := flag.String("s", "localhost:6379", "server ip:port")
	removeKey := flag.String("k", "", "key to remove (can contain wildcard; watch out for shell expansion !)")
	timeWait := flag.Int("t", 1, "number of minutes to wait for next")
	dryRun := flag.Bool("n", false, "dry run mode, listing how many to delete")

	flag.Parse()

	if *removeKey == "" {
		fmt.Println("Error: -k must be specified")
		os.Exit(1)
	}

	if *dryRun {
		modeMsg = "[DRY RUN MODE]"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     *server,
		Password: "",
		DB:       0,
	})

	for {
		n := 0

		pong, err := client.Ping().Result()
		if err != nil {
			log.WithFields(log.Fields{
				"status": "Error",
				"server": *server,
				"err":    err,
				"pong":   pong,
			}).Error("Can't connect to server")

			time.Sleep(30 * time.Second)
			continue
		}

		log.WithFields(log.Fields{
			"status": fmt.Sprintf("%s started", modeMsg),
			"key":    *removeKey,
			"server": *server,
		}).Info("Removing key")

		for {
			var keys []string
			keys, cursor, err = client.Scan(cursor, *removeKey, 10).Result()
			if err != nil {
				log.WithFields(log.Fields{
					"status": "Error",
					"server": *server,
					"err":    err,
				}).Error("Can't retrieve keys data")

				break
			}

			if !*dryRun {
				client.Del(keys...)
			}

			n += len(keys)

			if cursor == 0 {
				break
			}
		}

		log.WithFields(log.Fields{
			"status": fmt.Sprintf("%s deleted", modeMsg),
			"key":    *removeKey,
			"server": *server,
			"total":  n,
		}).Info("Deleted key")

		time.Sleep(time.Duration(*timeWait) * time.Minute)
	}
}
