package app

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats-server/server"
	"github.com/sirupsen/logrus"

	"github.com/nats-io/go-nats"
)

type App struct {
	Port int
	Nodes []string
}

func (a *App) Start() error {

	logrus.Infof("Starting service on Port %d", a.Port)
	logrus.Infof("Initial cluster Nodes: %v", a.Nodes)

	opts := &server.Options{
		Port: -1, // Random port for client connections
		Host: "localhost",
		//Debug: true,
		//Trace: true,
		Cluster: server.ClusterOpts{
			Port: a.Port,
			Host: "localhost",
			Advertise: "localhost",
			ConnectRetries: 10,
		},

		Routes: server.RoutesFromStr(strings.Join(a.Nodes, ",")),
	}

	s := server.New(opts)
	s.ConfigureLogger()

	go s.Start()

	if !s.ReadyForConnections(3 * time.Second) {
		return errors.New("server did not start in time")
	}


	subj := fmt.Sprintf("client-%d", a.Port)

	ports := s.PortsInfo(1 * time.Second)
	if len(ports.Nats) < 1 {
		return errors.New("no NATS server address returned")
	}

	url := ports.Nats[0]
	copts := []nats.Option{
		nats.Name(subj),
	}

	nc, err := nats.Connect(url, copts...)
	if err != nil {
		return err
	}

	_, err = nc.Subscribe("*", func(msg *nats.Msg) {

		logrus.Infof("Got message: %s", msg.Data)
	})
	if err != nil {
		return err
	}

	a.publishLoop(nc)

	return nil
}

func (a *App) publishLoop(nc *nats.Conn) {

	msg := fmt.Sprintf("Ping from client-%d @ %d", a.Port, time.Now().Unix())

	for {

		err := nc.Publish("*", []byte(msg))
		if err != nil {
			logrus.Errorf("Failed to publish: %s", err.Error())
		}

		time.Sleep(5 * time.Second)
	}
}
