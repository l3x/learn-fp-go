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
	DataFilepath     string   `toml:"data_filepath"`
	AppEnv           string   `toml:"app_env"`
	LogLevel         string   `toml:"log_level"`
	LogVerbose       bool     `toml:"log_verbose"`
	LogDebugChars    string   `toml:"log_debug_chars"`
	LogDebugInfo     bool     `toml:"log_debug_info"`
	LogTimeTrack     bool     `toml:"log_timetrack"`
	I18nFilename     string   `toml:"i18n_filename"`
}

var Config Conf

func GetOptions() bool {
	var configFile string
	flag.StringVar(&configFile, "config", "", "Configuration file")
	flag.StringVar(&Config.DataFilepath, "data-filepath", "./data/cars.base64", "Full path to the file containing the base64 car strings")
	flag.StringVar(&Config.AppEnv, "development", "development", "Runtime environment. Determines whether to run semver scripts or expect env vars")
	flag.StringVar(&Config.LogLevel, "log-level", "INFO", "Log level options: PANC, FATL, ERRO, WARN, INFO, DEBG")
	flag.BoolVar(&Config.LogVerbose, "log-verbose", false, "Whether to include timestamps and log level on all log entries")
	flag.StringVar(&Config.LogDebugChars, "log-debug-chars", ">>", "The character(s) used to preface debug lines")
	flag.BoolVar(&Config.LogDebugInfo, "log-debug-info", true, "Whether to log debug output to the log (set to true for debug purposes)")
	flag.BoolVar(&Config.LogTimeTrack, "log-timetrack", true, "Enable or disable logging of utils/TimeTrack() (For benchmarking/debugging)")
	flag.StringVar(&Config.I18nFilename, "i18n-filename", "en-us.all.json", "i18n translation file name, see github.com/nicksnyder/go-i18n")
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

