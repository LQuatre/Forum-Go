package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type App struct {
    Address      string
    Static       string
    Log          string
    Locale       string
    Language     string
}

type Database struct {
    Driver      string
    Address        string
    Database    string
    User        string
    Password    string
}

type Configuration struct {
    App App
    Db  Database
    LocaleBundle *i18n.Bundle
}

var config *Configuration
var once sync.Once

func LoadConfig() *Configuration {
    once.Do(func() {
        file, err := os.Open("config.json")
        if err != nil {
            log.Fatalln("Cannot open config file", err)
        }
        decoder := json.NewDecoder(file)
        config = &Configuration{}
        err = decoder.Decode(config)
        if err != nil {
            log.Fatalln("Cannot get configuration from file", err)
        }
        bundle := i18n.NewBundle(language.French)
        bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
        bundle.MustLoadMessageFile(config.App.Locale + "/active.fr.json")
        bundle.MustLoadMessageFile(config.App.Locale + "/active." + config.App.Language + ".json")
        config.LocaleBundle = bundle
    })
    return config
}