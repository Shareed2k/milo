package internal

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"encoding/json"
	"github.com/imdario/mergo"
)

type Core interface {
	OnStop()
	GetSettings() Settings
	GetLog() *log.Logger
	GetMaster() MasterOperator
	GetMinion() MinionOperator
}

type LogFormatter struct {
	prefix string
}

type core struct {
	Settings
	log    *log.Logger
	master MasterOperator
	minion MinionOperator
}

func NewCore(s Settings) Core {
	c := &core{Settings: s}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
			os.Exit(1)
		}
	}()

	// Read config
	if err := c.ReadConfig(); err != nil {
		panic(err)
	}

	// Initialize all services
	c.initializeForeground()

	// Bootstrap
	if err := c.initBootstrap(); err != nil {
		panic(err)
	}

	return c
}

func (c *core) initializeForeground() error {
	// Set Logger
	c.log = log.New()
	c.log.SetLevel(log.InfoLevel | log.DebugLevel | log.ErrorLevel)

	//Set Operator
	if c.MasterMode == true {
		c.master = NewMaster(c)
	} else {
		c.minion = NewMinion(c)
	}

	return nil
}

func (c *core) initBootstrap() error {
	var err error

	if c.MasterMode == true {
		err = c.master.InitBootstrap()
	} else {
		err = c.minion.InitBootstrap()
	}

	return err
}

func (c *core) GetSettings() Settings {
	return c.Settings
}

func (c *core) OnStop() {
	c.log.Out = nil
}

func (c *core) GetMaster() MasterOperator {
	return c.master
}

func (c *core) GetMinion() MinionOperator {
	return c.minion
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

func (c *core) ReadConfig() error {
	// Set Settings from config file
	if c.ConfigFilePath != "" {
		var configFileSettings Settings
		configFile, err := os.Open(c.ConfigFilePath)
		defer configFile.Close()

		if err != nil {
			return err
		}
		if err := json.NewDecoder(configFile).Decode(&configFileSettings); err != nil {
			return err
		}
		// Merge in command line settings (which overwrite respective config file settings)
		if err := mergo.Merge(&c.Settings, configFileSettings); err != nil {
			return err
		}

		fmt.Println("tmp", configFileSettings)
		fmt.Println("original", c.Settings)
	}

	return nil
}