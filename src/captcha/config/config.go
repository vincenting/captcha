package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Config struct {
	InitialCount  int `json:"initial_count"`
	CheckInterval int `json:"check_interval"`
	Threshold     int `json:"threshold"`
	UpdateCount   int `json:"update_count"`
}

var (
	config     *Config
	configLock = new(sync.RWMutex)
)

func loadConfig(fail bool) {
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Println("open config: ", err)
		if fail {
			os.Exit(1)
		}
	}
	temp := new(Config)
	if err = json.Unmarshal(file, temp); err != nil {
		log.Println("parse config: ", err)
		if fail {
			os.Exit(1)
		}
	}
	configLock.Lock()
	config = temp
	configLock.Unlock()
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

// go calls init on start
func init() {
	loadConfig(true)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			loadConfig(false)
			log.Print(GetConfig())
			log.Println("Reloaded")
		}
	}()
}
