package utils

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"reflect"
	"strconv"
	"github.com/pkg/errors"
)

type Conf struct {
	Port	string   `toml:"port"`
	LogDebugInfo         bool     `toml:"log_debug_info"`
	MaxConcurrentConnections	int   `toml:"max_concurrent_connections"`
	MaxNumber	int   `toml:"max_number"`
	UseNumberHandler	bool   `toml:"use_number_handler"`
}

var Config Conf

func GetOptions() bool {
	var configFile string

	flag.StringVar(&configFile, "config", "", "Configuration file")
	flag.StringVar(&Config.Port, "port", "8080", "Port that the API listens on")
	flag.BoolVar(&Config.LogDebugInfo, "log-debug-info", false, "Whether to log debug output to the log (set to true for debug purposes)")
	flag.IntVar(&Config.MaxConcurrentConnections, "max-concurrent-connections", 6, "Maximum number of concurrent connections (not currently used)")
	flag.IntVar(&Config.MaxNumber, "max_number", 10, "Maximum number that user is allowed to enter")
	flag.BoolVar(&Config.UseNumberHandler, "use-number-handler", true, "Use number handler (to display number, optionally with FormatNumber applied) else display files in project root")

	flag.Parse()

	if configFile != "" {
		if _, err := toml.DecodeFile(configFile, &Config); err != nil {
			HandlePanic(errors.Wrap(err, "unable to read config file"))
		}
	}
	return true
}

type Datastore interface {}

func UpdateConfigVal(d Datastore, key, val string) (oldValue string) {
	Debug.Printf("key (%s), val (%v)\n", key, val)
	value := reflect.ValueOf(d)
	if value.Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	valElem := value.Elem()
	for i := 0; i < valElem.NumField(); i++ {
		tag := valElem.Type().Field(i).Tag
		field := valElem.Field(i)
		switch tag.Get("toml") {
		case key:
			if fmt.Sprintf("%v", field.Kind()) == "int" {
				oldValue = strconv.FormatInt(field.Int(), 10)
				intVal, err := strconv.Atoi(val)
				if err != nil {
					fmt.Printf("could not parse int, key(%s) val(%s)", key, val)
				} else {
					field.SetInt(int64(intVal))
				}
			} else if fmt.Sprintf("%v", field.Kind()) == "bool" {
				oldValue = strconv.FormatBool(field.Bool())
				b, err := strconv.ParseBool(val)
				if err != nil {
					fmt.Printf("could not parse bool, key(%s) val(%s)", key, val)
				} else {
					field.SetBool(b)
				}
			} else {
				// Currently only supports bool, int and string
				oldValue = field.String()
				field.SetString(val)
			}
		}
	}
	return
}

