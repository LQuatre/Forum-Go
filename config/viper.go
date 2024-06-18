package config

import (
	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

var ViperConfig Configuration

func init()  {
    runtimeViper := viper.New()
    runtimeViper.AddConfigPath(".")
    runtimeViper.SetConfigName("config")
    runtimeViper.SetConfigType("json")
    err := runtimeViper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    runtimeViper.Unmarshal(&ViperConfig)

    bundle := i18n.NewBundle(language.English)
    bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
    bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active.fr.json")
    bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
    ViperConfig.LocaleBundle = bundle

    runtimeViper.WatchConfig()
    runtimeViper.OnConfigChange(func(e fsnotify.Event) {
        runtimeViper.Unmarshal(&ViperConfig)
        ViperConfig.LocaleBundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
    })
}