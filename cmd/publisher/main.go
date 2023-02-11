package main

import (
	"WEB_REST_exm0302/configs"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var usageStr = `
Usage: stan-pub [options] <subject> <message>

Options:
	-s,  --server   <url>            NATS Streaming server URL(s)
	-c,  --cluster  <cluster name>   NATS Streaming cluster name
	-id, --clientid <client ID>      NATS Streaming client ID
	-a,  --async                     Asynchronous publish mode
	-cr, --creds    <credentials>    NATS 2.0 Credentials
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func readJson() [][]byte {

	res := make([][]byte, 10, 10)

	if err := godotenv.Load("pub.env"); err != nil {
		logrus.Println("No .env file found")
	}
	//"ABSOLUTE_PATH" - полный путь к папке с json файлами для отправки
	dir := os.Getenv("ABSOLUTE_PATH")
	files, errReadDir := os.ReadDir(dir)
	if errReadDir != nil {
		fmt.Println(errReadDir)
		return nil
	}
	for _, json := range files {
		if json.IsDir() {
			continue
		}
		filename := filepath.Join(dir, json.Name())
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}
		res = append(res, data)
	}
	return res
}

func main() {
	var (
		clusterID string
		clientID  string
		URL       string
		async     bool
		userCreds string
	)

	flag.StringVar(&URL, "s", stan.DefaultNatsURL, "The mynats server URLs (separated by comma)")
	flag.StringVar(&URL, "server", stan.DefaultNatsURL, "The mynats server URLs (separated by comma)")
	flag.StringVar(&clusterID, "c", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clusterID, "cluster", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clientID, "id", "stan-pub", "The NATS Streaming client ID to connect with")
	flag.StringVar(&clientID, "clientid", "stan-pub", "The NATS Streaming client ID to connect with")
	flag.BoolVar(&async, "a", false, "Publish asynchronously")
	flag.BoolVar(&async, "async", false, "Publish asynchronously")
	flag.StringVar(&userCreds, "cr", "", "Credentials File")
	flag.StringVar(&userCreds, "creds", "", "Credentials File")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}
	// Use UserCredentials
	if userCreds != "" {
		opts = append(opts, nats.UserCredentials(userCreds))
	}

	// Connect to NATS
	nc, err := nats.Connect(URL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer sc.Close()

	//Название канала
	subj := configs.Mysubj
	//Чтение json файлов, создаём слайс из сообщений для отправки
	msg := make([][]byte, 10, 10)
	msg = readJson()

	if !async {

		for _, unitInMsg := range msg {
			if unitInMsg != nil {
				time.Sleep(1000 * time.Millisecond)
				err = sc.Publish(subj, unitInMsg)
				if err != nil {
					log.Fatalf("Error during publish: %v\n", err)
				}
				log.Printf("Published [%s] : '%s'\n", subj, msg)
			}
		}

	}
}
