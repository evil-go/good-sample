package main

import (
	"errors"
	"fmt"
	"github.com/evil-go/good-sample/config"
	"github.com/evil-go/good-sample/greet"
	"github.com/evil-go/good-sample/server"
	"net/http"
	"os"
)

type Config struct {
	DefaultMessage string
	BobMessage     string
	JuliaMessage   string
	Path           string
}

func main() {
	c, err := loadProperties()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dao := greet.Dao{
		DefaultMessage: c.DefaultMessage,
		BobMessage:     c.BobMessage,
		JuliaMessage:   c.JuliaMessage,
	}
	svc := greet.Service{GreetingFinder: dao}
	controller := greet.Controller{Greeter: svc}

	err = server.Start(server.Endpoint{c.Path, http.MethodGet, controller})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadProperties() (Config, error) {
	if len(os.Args) == 1 {
		return Config{}, errors.New("must specify config file to launch")
	}
	c := Config{}
	m, err := config.LoadPropertiesFile(os.Args[1])
	if err != nil {
		return c, err
	}
	c.DefaultMessage, err = m.GetString("message.default")
	if err != nil {
		return c, err
	}
	c.BobMessage, err = m.GetString("message.bob")
	if err != nil {
		return c, err
	}
	c.JuliaMessage, err = m.GetString("message.julia")
	if err != nil {
		return c, err
	}
	c.Path, err = m.GetString("controller.path.hello")
	if err != nil {
		return c, err
	}
	return c, nil
}
