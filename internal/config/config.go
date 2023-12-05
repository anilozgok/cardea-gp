package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

var (
	configWatcher = viper.New()
	secretWatcher = viper.New()
)

func Get() (*Config, error) {
	config := Default()
	err := watchFile[Config](config, "configs", configWatcher)
	if err != nil {
		return nil, err
	}

	err = watchFile[Secrets](config.Secrets, "secrets", secretWatcher)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func watchFile[T any](bindTo *T, file string, watcher *viper.Viper) error {
	watcher.AddConfigPath("./configs")    // local folder
	watcher.AddConfigPath("/app/configs") // required for container
	watcher.SetConfigName(file)
	watcher.SetConfigType("json")
	watcher.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := watcher.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read in config. err: %s", err)
	}

	if err := watcher.Unmarshal(bindTo); err != nil {
		return fmt.Errorf("failed to unmarshal config. err: %s", err)
	}

	watcher.WatchConfig()
	watcher.OnConfigChange(func(event fsnotify.Event) {
		if err := watcher.Unmarshal(bindTo); err == nil {
			zap.L().Info(fmt.Sprintf("config %s updated", file))
		} else {
			zap.L().Error(fmt.Sprintf("could not unmarshal %s config.", file), zap.Error(err))
		}
	})

	return nil
}
