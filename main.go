package main

import (
	"fmt"
	"os"

	"github.com/deranjer/tinyMonitor/config"

	"github.com/rs/zerolog"
	"nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/pub"

	// register transports
	_ "nanomsg.org/go/mangos/v2/transport/tcp" //TODO change to /all to register all transports
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

var (
	err  error
	sock mangos.Socket
	//Logger is the global Logger variable
	Logger zerolog.Logger
)

func main() {
	serverSettings, Logger := config.SetupServer() //setup logging and all server settings
	Logger.Info().Msg("Server and Logger configuration complete")
	sock, err = pub.NewSocket()
	if err != nil {
		Logger.Fatal().Err(err).Msg("Can't get new pub socket")
	}
	err = sock.Listen(serverSettings.ListenAddr)
	if err != nil {
		Logger.Fatal().Err(err).Str("Connection String", serverSettings.ListenAddr).Msg("Failed to open listen socket")
	}
	Logger.Info().Str("Address", serverSettings.ListenAddr).Msg("Listen socket opened with no errors")
	fmt.Println("Done!")

}