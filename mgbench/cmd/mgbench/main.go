package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection string = "test"

func main() {
	c := flag.Uint("c", 0, "Number of clients")
	t := flag.Uint("t", 0, "Number of transactions per client")
	p := flag.Uint("p", 27017, "Server port")
	h := flag.String("h", "localhost", "Server host")
	u := flag.String("u", "", "User name")
	d := flag.String("d", "", "Database name")

	flag.Parse()

	usage(*c, *t, *u, *d)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	passwd := os.Getenv("MONGO_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", *u, passwd, *h, *p, *d)

	log.Println(uri)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Printf("error connecting to MongoDB: %v", err)
		return
	}

	defer func() {
		if err != nil {
			if err = client.Disconnect(ctx); err != nil {
				log.Printf("error disconnecting from MongoDB: %v", err)
			}
		}
	}()

	latencies := []float64{}
	latencyChan := make(chan float64, 10000)
	var mainWG sync.WaitGroup
	mainWG.Add(1)
	go func() {
		defer mainWG.Done()

		for latency := range latencyChan {
			latencies = append(latencies, latency)
		}
	}()

	var i uint

	fmt.Println("Starting benchmark...")

	start := time.Now()
	var wg sync.WaitGroup
	for i = 0; i < *c; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var j uint
			for j = 0; j < *t; j++ {
				latency := time.Now()

				_, err := client.Database(*d).Collection(collection).InsertOne(ctx, map[string]interface{}{"test": "test"})
				if err != nil {
					log.Printf("error inserting document: %v", err)
				}

				elapsed := time.Since(latency)
				latencyChan <- elapsed.Seconds()
			}

			filter := bson.D{{Key: "test", Value: "test"}}
			opts := options.Find().SetLimit(100)

			for j = 0; j < *t; j++ {
				start := time.Now()

				_, err := client.Database(*d).Collection(collection).Find(ctx, filter, opts)
				if err != nil {
					log.Printf("error fetching documents: %v", err)
				}

				elapsed := time.Since(start)
				latencyChan <- elapsed.Seconds()
			}
		}()
	}

	wg.Wait()
	totalTime := time.Since(start).Seconds()

	close(latencyChan)
	mainWG.Wait()

	err = client.Database(*d).Collection(collection).Drop(ctx)
	if err != nil {
		log.Printf("error dropping collection: %v", err)
	}

	rps, avg := calculateResult(latencies, totalTime)

	fmt.Printf("total time: %v\n", getUnit(totalTime))
	fmt.Println("number of clients: ", *c)
	fmt.Println("number of write ops per client: ", *t)
	fmt.Println("number of read ops per client: ", *t)
	fmt.Println("number of transactions per client: ", (*t)*2)
	fmt.Println("number of transactions actually processed: ", len(latencies))
	fmt.Printf("latency average: %v\n", getUnit(avg))
	fmt.Printf("tps: %.3f\n", rps)
}

func getUnit(aTime float64) string {
	if aTime < 0.001 {
		return fmt.Sprintf("%.3f ns", aTime*1000000)
	} else if aTime < 1 {
		return fmt.Sprintf("%.3f ms", aTime*1000)
	} else {
		return fmt.Sprintf("%.3f s", aTime)
	}
}

func calculateResult(latencies []float64, totalTime float64) (float64, float64) {
	var sum float64
	lenght := len(latencies)

	for _, latency := range latencies {
		sum += latency
	}

	return float64(lenght) / totalTime, sum / float64(lenght)
}

func usage(c, t uint, u, d string) {
	if c == 0 || t == 0 || u == "" || d == "" {
		fmt.Println("Usage: MONGO_PASSWORD='' mgbench -h <host> -p <port> -c <# of clients> -t <# of transactions per client> -u <user> -d <database>")
		os.Exit(1)
	}
}
