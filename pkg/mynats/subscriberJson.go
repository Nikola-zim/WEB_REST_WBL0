package mynats

import (
	"flag"
	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"log"
)

func printMsg(m *stan.Msg, i int, sn *SubsNats) {

	jsonFormNats := m.Data
	err := sn.services.WriteNatsJsonInCash(jsonFormNats)
	if err != nil {
		log.Printf("Ошибка записи в кеш")
		return
	}
	errDB := sn.services.WriteNatsJsonInDB(jsonFormNats)
	if errDB != nil {
		log.Printf("Ошибка записи в БД")
		return
	}
	log.Printf("[#%d] Received: %s\n", i, m)
	return
}

func Subscriber(sn *SubsNats) ([]byte, error) {
	var (
		clusterID, clientID string
		URL                 string
		userCreds           string
		showTime            bool
		qgroup              string
		unsubscribe         bool
		startSeq            uint64
		startDelta          string
		deliverAll          bool
		newOnly             bool
		deliverLast         bool
		durable             string
	)
	flag.StringVar(&URL, "s", stan.DefaultNatsURL, "The mynats server URLs (separated by comma)")
	flag.StringVar(&URL, "server", stan.DefaultNatsURL, "The mynats server URLs (separated by comma)")
	flag.StringVar(&clusterID, "c", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clusterID, "cluster", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clientID, "id", "stan-sub", "The NATS Streaming client ID to connect with")
	flag.StringVar(&clientID, "clientid", "stan-sub", "The NATS Streaming client ID to connect with")
	flag.BoolVar(&showTime, "t", false, "Display timestamps")
	// Subscription options
	flag.Uint64Var(&startSeq, "seq", 0, "Start at sequence no.")
	flag.BoolVar(&deliverAll, "all", true, "Deliver all")
	flag.BoolVar(&newOnly, "new_only", false, "Only new messages")
	flag.BoolVar(&deliverLast, "last", false, "Start with last value")
	flag.StringVar(&startDelta, "since", "", "Deliver messages since specified time offset")
	flag.StringVar(&durable, "durable", "", "Durable subscriber name")
	flag.StringVar(&qgroup, "qgroup", "", "Queue group name")
	flag.BoolVar(&unsubscribe, "unsub", false, "Unsubscribe the durable on exit")
	flag.BoolVar(&unsubscribe, "unsubscribe", false, "Unsubscribe the durable on exit")
	flag.StringVar(&userCreds, "cr", "", "Credentials File")
	flag.StringVar(&userCreds, "creds", "", "Credentials File")

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Subscriber")}
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

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", URL, clusterID, clientID)

	// Process Subscriber Options.
	startOpt := stan.StartAt(pb.StartPosition_NewOnly)

	startOpt = stan.StartWithLastReceived()

	subj, i := "foo", 0
	var resultJson []byte
	//
	mcb := func(msg *stan.Msg) {
		i++
		resultJson = msg.Data

		printMsg(msg, i, sn)
	}

	sub, err := sc.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}

	log.Printf("Listening on [%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", subj, clientID, qgroup, durable)

	if showTime {
		log.SetFlags(log.LstdFlags)
	}

	cleanupDone := make(chan bool)

	defer func() { cleanupDone <- true }()
	<-cleanupDone

	defer sc.Close()
	defer sub.Unsubscribe()

	return resultJson, nil
}
