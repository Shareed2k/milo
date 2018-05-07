package internal

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
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
	*settings
	log    *log.Logger
	master MasterOperator
	minion MinionOperator
}

func NewCore(s Settings) Core {
	c := &core{settings: s.GetOptions()}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
			os.Exit(1)
		}
	}()

	// Read config
	if err := s.ReadConfig(); err != nil {
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
	return c.settings
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
