package internal

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/imdario/mergo"
	"os"
)

type Core interface {
	OnStop()
	GetSettings() Settings
	GetLog() *log.Logger
}

type LogFormatter struct {
	prefix string
}

type core struct {
	settings
	log *log.Logger
}

func NewCore(s Settings) Core {
	c := &core{settings: *s.GetOptions()}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	// Initialize all services
	if err := c.initializeForeground(); err != nil {
		panic(err)
	}

	// Bootstrap
	if err := c.initBootstrap(); err == nil {
		panic(err)
	}

	return c
}

func (c *core) initializeForeground() error {

	return nil
}

func (c *core) initBootstrap() error {

	return nil
}

func (c *core) GetSettings() Settings {
	return &c.settings
}

func (c *core) OnStop() {
	c.log.Out = nil
}

func (c *core) GetLog() *log.Logger {
	c.log.Formatter = &LogFormatter{"Core"}
	return c.log
}

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	entry.Message = fmt.Sprintf("%s: %s", f.prefix, entry.Message)
	tf := log.TextFormatter{}
	return tf.Format(entry)
}

func (c *core) readConfig() error {
	// Set Settings from config file
	if c.ConfigFilePath != "" {
		var configFileSettings settings
		configFile, err := os.Open(c.ConfigFilePath)
		defer configFile.Close()

		if err != nil {
			return err
		}
		if err := json.NewDecoder(configFile).Decode(&configFileSettings); err != nil {
			return err
		}
		// Merge in command line settings (which overwrite respective config file settings)
		if err := mergo.Merge(&c.settings, configFileSettings); err != nil {
			return err
		}
	}

	return nil
}
