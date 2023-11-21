package config

import (
	"fmt"
	"os"
	"pack-sizes/pkg/split"
	"strconv"
	"strings"
)

const (
	envPort             = "APP_PORT"
	envHost             = "APP_HOST"
	envBuckets          = "APP_BUCKETS"
	envDynamicDepth     = "APP_USE_DYNAMIC_ALGORITHM_WITH_DEPTH"
	envBucketsSeparator = ","
	defaultPort         = 12345
	defaultHost         = "localhost"
)

var (
	defaultBuckets = []int{250, 500, 1000, 2000, 5000}
)

type Config struct {
	Buckets []int
	Port    int
	Host    string
	Options split.Options
}

func Get() Config {
	c := Config{
		Buckets: defaultBuckets,
		Host:    defaultHost,
		Port:    defaultPort,
	}

	var err error
	if os.Getenv(envPort) != "" {
		portStr := os.Getenv(envPort)
		c.Port, err = strconv.Atoi(portStr)
		if err != nil {
			fmt.Printf("error parsing %s env var: %s", envPort, err)
			os.Exit(1)
		}
	}
	if os.Getenv(envHost) != "" {
		c.Host = os.Getenv(envPort)
	}
	if os.Getenv(envBuckets) != "" {
		bucketsStr := os.Getenv(envBuckets)
		for _, e := range strings.Split(bucketsStr, envBucketsSeparator) {
			bucket, err := strconv.Atoi(e)
			if err != nil {
				fmt.Printf("error parsing %s env var: %s", envBuckets, err)
				os.Exit(1)
			}
			c.Buckets = append(c.Buckets, bucket)
		}
		c.Host = os.Getenv(envBuckets)
	}
	if os.Getenv(envDynamicDepth) != "" {
		depthStr := os.Getenv(envDynamicDepth)
		c.Options.Depth, err = strconv.Atoi(depthStr)
		if err != nil {
			fmt.Printf("error parsing %s env var: %s", envDynamicDepth, err)
			os.Exit(1)
		}
	}

	fmt.Printf("using config: %+v\n", c)
	fmt.Printf("to use other config override using %s, %s and comma (%s) separated %s\n", envPort, envHost, envBucketsSeparator, envBuckets)
	fmt.Printf("to use other than Greedy algorith set %s depth for Dymamic algorithm\n", envDynamicDepth)
	return c
}
