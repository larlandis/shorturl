package cmd

import "flag"

var (
	Port           string
	ConfigFile     string
	StorageCluster string
	MetricsServer  string
)

func ParseArgs() {
	flag.StringVar(&Port, "port", "8080", "application listen port")
	flag.StringVar(&ConfigFile, "config", "/cmd/config/config.yml", "config yaml file")
	flag.StringVar(&StorageCluster, "storage-cluster", "cassandra", "storage cluster name")
	flag.StringVar(&MetricsServer, "metrics-server", "http://influxdb:8086", "metrics server address")
	flag.Parse()
}
