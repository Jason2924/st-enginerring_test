package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	con "github.com/Jason2924/st-enginerring_test/config"
	dtb "github.com/Jason2924/st-enginerring_test/databases"
	"github.com/Jason2924/st-enginerring_test/graph"
	hpr "github.com/Jason2924/st-enginerring_test/helper"
	rep "github.com/Jason2924/st-enginerring_test/repositories"
	svc "github.com/Jason2924/st-enginerring_test/services"

	"github.com/99designs/gqlgen/graphql/handler"
)

func main() {
	// get config from env file
	conf, erro := con.Load("./", "app", "env")
	if erro != nil {
		log.Fatalln("Failed occured while setting config", erro)
	}

	// connect database
	dtbs := dtb.NewMysqlDatabase(&conf.Mysql)
	dtbs.Connect()
	if erro = dtbs.Ping(context.Background()); erro != nil {
		log.Fatalln("Failed occured while connecting cache", erro)
	}

	// create repository
	pdtRepo := rep.NewProductRepository(dtbs)
	pdtSrvc := svc.NewProductService(pdtRepo)
	resr := &graph.Resolver{ProductService: pdtSrvc}

	// create server

	serv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resr}))

	http.Handle("/query", serv)

	if conf.Mysql.ImportData && conf.Mode != "production" {
		hpr.ImportProductData(pdtSrvc, "product-data.csv")
	}

	// set graceful shutdown
	sigChan := make(chan os.Signal, 1)
	// create the background and listen and serve
	go func() {
		if erro := http.ListenAndServe(":"+conf.Port, nil); erro != nil && erro != http.ErrServerClosed {
			log.Fatalln("Failed occured while starting server", erro)
		}
	}()
	// the signal channel to listen the Interrupt and Termination signals
	// SIGINT = Interrupt | SIGTERM = Termination
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// set timeout for closing all connection
	// ctxt, cacl := context.WithTimeout(context.Background(), 5*time.Second)
	// defer func() {
	// 	// close all connection
	// 	dtbs.Close()
	// 	cacl()
	// }()
	// shutdown the server
	// if erro := serv.Stop(ctxt); erro == context.DeadlineExceeded {
	// 	fmt.Println("Halted active connections")
	// }
	close(sigChan)
}
