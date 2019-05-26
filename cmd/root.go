package cmd

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/alexejk/nats-test/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type rootOps struct{
	port *int
	clusterPort *int
}

func RootCmd() *cobra.Command {

	o := &rootOps{
		port: new(int),
		clusterPort: new(int),
	}

	cmd := &cobra.Command{
		Use: "nats-test",
		RunE: o.runE,
	}

	cmd.Flags().IntVarP(o.port, "port", "p", -1, "port to start on")
	cmd.Flags().IntVarP(o.clusterPort, "cluster-port", "c", -1, "cluster port to start on")

	cmd.SilenceUsage = true

	return cmd
}

func (o *rootOps) runE (cmd *cobra.Command, args []string) error {

	if o.port == nil || *o.port <= 0 {
		return errors.New("port cannot be empty or below 0")
	}

	if o.clusterPort == nil || *o.clusterPort <= 0 {
		return errors.New("cluster port cannot be empty or below 0")
	}

	servers := o.lookupClusterNodes()


	a := &app.App{
		Port: *o.port,
		ClusterPort: *o.clusterPort,
		Nodes: servers,
	}

	return a.Start()
}

func (o *rootOps) lookupClusterNodes() []string{

	_, addrs, err :=net.LookupSRV("nats", "tcp", "nats.local")
	if err != nil {
		logrus.Fatal(err)
	}

	servers := make([]string, len(addrs))
	for i, addr := range addrs {
		target := strings.TrimSuffix(addr.Target, ".")
		url := fmt.Sprintf("nats://%s:%d", target, addr.Port)

		servers[i] = url
		logrus.Infof("SRV-discovered NATS Server: %s", url)
	}

	return servers
}