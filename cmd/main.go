package main

import (
	"flag"
	"github.com/birdlinks/gmqtt/internal/config"
	"github.com/birdlinks/gmqtt/internal/persistence/bolt"
	"go.etcd.io/bbolt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/birdlinks/gmqtt"
	"github.com/birdlinks/gmqtt/internal/listeners"
	"github.com/birdlinks/gmqtt/internal/log"
)

func main() {

	f := flag.String("f", "config.yaml", "config file path")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	conf, err := config.Load(*f)
	if err != nil {
		panic(err)
	}

	log.Init(conf.Log)

	// server options...
	options := &mqtt.Options{
		BufferSize:      conf.Mqtt.BufferSize,      // Use default values 1024 * 256
		BufferBlockSize: conf.Mqtt.BufferBlockSize, // Use default values 1024 * 8
	}

	log.Info("MQTT Broker initializing...")
	server := mqtt.NewServer(options)
	tcp := listeners.NewTCP("t1", conf.Mqtt.TCP)
	err = server.AddListener(tcp, nil)
	if err != nil {
		log.Fatal("", log.Any("err", err))
	}

	ws := listeners.NewWebsocket("ws1", conf.Mqtt.WS)
	err = server.AddListener(ws, nil)
	if err != nil {
		log.Fatal("", log.Any("err", err))
	}

	stats := listeners.NewHTTPStats("stats", conf.Mqtt.HTTP)
	err = server.AddListener(stats, nil)
	if err != nil {
		log.Fatal("", log.Any("err", err))
	}

	err = server.AddStore(bolt.New("mqtt.db", &bbolt.Options{
		Timeout: 500 * time.Millisecond,
	}))
	if err != nil {
		log.Fatal("", log.Any("err", err))
	}

	go server.Serve()
	log.Info("server started")

	<-done
	log.Info("caught signal...")

	server.Close()
	log.Info("finished")

}
