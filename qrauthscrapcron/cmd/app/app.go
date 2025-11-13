package app

import (
	"os"
	"os/signal"
	"qrauthscrapcron/authenticator"
	"qrauthscrapcron/config"
	"qrauthscrapcron/network"
	"qrauthscrapcron/repository"
	"qrauthscrapcron/service"
	"syscall"
)

type App struct {
	config *config.Config

	network *network.Network

	authenticator authenticator.AuthenticatorImpl
	service       service.ServiceImpl
	repository    repository.RepositoryImpl

	stop chan struct{}
}

func NewApp(config *config.Config) *App {
	a := &App{
		config: config,
		stop:   make(chan struct{}),
	}

	var err error
	if a.authenticator, err = authenticator.NewAuthenticator(config); err != nil {
		panic(err)
	}
	a.network = network.NewNetwork(config, a.service, a.authenticator)

	c := make(chan os.Signal, 1)

	//this listen user input of ctrl c
	signal.Notify(c, syscall.SIGINT)

	go func() {
		<-c
		a.exit()
	}()

	return a
}

func (a *App) Wait() {
	<-a.stop
	os.Exit(1)
}

func (a *App) exit() {
	a.stop <- struct{}{}
}

func (a *App) Run() {
	a.network.Run()
}
