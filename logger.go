package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"flag"
	"log"

	"github.com/MarinX/keylogger"
)

const (
	sleepSeconds = 2
)

var (
	expectedKey string
	keyboardDevices, mouseDevices arrayFlags
)

func main() {
	initFlags()
	enforceFlags()

	// Wait few seconds before to start following devices
	time.Sleep(time.Second * sleepSeconds)

	devs, err := keylogger.NewDevices()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range devs {
		for _, kd := range keyboardDevices {
			if strings.Trim(val.Name, "\"") == kd {
				go followKeyboardDevice(val)
			}
		}
		for _, md := range mouseDevices {
			if strings.Trim(val.Name, "\"") == md {
				go followMouseDevice(val)
			}
		}
	}
	for {
	}
}

// followKeyboardDevice exit(0) if an specific key is clicked. Otherwise, exit(1)
func followKeyboardDevice(inputDevice *keylogger.InputDevice) {
	rd := keylogger.NewKeyLogger(inputDevice)
	in, err := rd.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	first := true
	for i := range in {
		if first {
			first = false
		} else {
			if i.Type == 1 {
				if i.KeyString() == expectedKey {
					os.Exit(0)
				} else {
					os.Exit(1)
				}
			}
		}
	}
}

// followMouseDevice exit(1) if a click is detected
func followMouseDevice(inputDevice *keylogger.InputDevice) {
	rd := keylogger.NewKeyLogger(inputDevice)
	in, err := rd.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range in {
		if i.Code == 272 {
			os.Exit(1)
		}
	}
}

func initFlags() {
	flag.StringVar(&expectedKey, "expectedKey", ".", "Expected key to be pressed")
	flag.Var(&keyboardDevices, "keyboardDevice", "A keyboard device to follow")
	flag.Var(&mouseDevices, "mouseDevice", "A mouse device to follow")
	flag.Parse()
}

func enforceFlags() {
	if expectedKey == "" {
		flag.Usage()
		log.Fatal("expectedKey flag is required")
	}

	if len(keyboardDevices) < 1 {
		flag.Usage()
		log.Fatal("at least one keyboardDevice flag is required")
	}

	if len(mouseDevices) < 1 {
		flag.Usage()
		log.Fatal("at least one mouseDevice flag is required")
	}
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return ""
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
