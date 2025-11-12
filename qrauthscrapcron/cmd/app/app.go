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

func NewApp(cfg *config.Config) *App {
	a := &App{
		config: cfg,
		stop:   make(chan struct{}),
	}

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)

	go func() {
		<-channel
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
}
