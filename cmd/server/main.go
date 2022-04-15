package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gobin/database"
	"gobin/server"
	"gobin/utils/config"
)

func main() {
	http := flag.Bool("http", false, "run in http mode")

	conf := read()

	log.Println("Initializing database")
	db, err := database.New(conf)
	if err != nil {
		log.Fatalln("new database:", err)
	}
	defer db.Close()

	srv := server.Serve(db, conf.Host)

	ch := make(chan os.Signal, 1)

	go func() {
		log.Println("Server started at " + srv.Addr)

		if *http {
			err = srv.ListenAndServe()
		} else {
			err = srv.ListenAndServeTLS(conf.Host.Cert, conf.Host.Key)
		}

		if err != nil {
			log.Println("Server:", err)
			ch <- nil
		}
	}()

	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)

	<-ch

	// Create a deadline to wait for closing all connections
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	log.Println("Shutting down server")
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Println("Error:", err)
	}
}

//

func read() *config.Config {
	var conf *config.Config

	ConfigPath := flag.String("config", "config.toml", "toml format config path")
	ReadEnv := flag.Bool("env", false, "read from environment values")
	NoRead := flag.Bool("noread", false, "skip config reading, environment values only")
	flag.Parse()

	if !*NoRead {
		c, err := config.Read(*ConfigPath)
		if err != nil {
			log.Fatalln("config:", err)
		}

		conf = c
	}

	if *NoRead || *ReadEnv {
		err := config.ReadEnv(conf)
		if err != nil {
			log.Fatalln("reading from environment:", err)
		}
	}

	return conf
}
