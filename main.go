package main

import (
	"fmt"
	"github.com/employee_manager/internal/handler/rest"
	employeeRepo "github.com/employee_manager/internal/repo/employee"
	employeeService "github.com/employee_manager/internal/service/employee"
	"github.com/employee_manager/pkg/middlewares"
	"github.com/employee_manager/pkg/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func main() {
	config := server.NewConfig()
	app := server.NewApp(config)
	g := &errgroup.Group{}
	g.Go(func() error {
		gin.SetMode(gin.DebugMode)
		engine := gin.New()

		endpoint := engine.Group("")
		engine.NoRoute(middlewares.NoRoute)
		engine.NoMethod(middlewares.NoMethod)

		app.RestHandler = SetupHandler()

		app.SetupRoutesAndMiddleware(endpoint, app.RestHandler)

		engine.HandleMethodNotAllowed = true
		engine.NoRoute(middlewares.NoRoute)
		engine.NoMethod(middlewares.NoMethod)

		log.Info("starting employee manager service server host:", config.Host, " port:", config.Port)

		err := startServer(engine, config.Host, config.Port)
		if err != nil {
			log.Panic("Fatal: service failed to start. ", err.Error())
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
		return
	}
}

func SetupHandler() rest.RestHandler {
	repo := employeeRepo.NewEmployeeRepo()
	service := employeeService.NewEmployeeService(repo)
	handler := rest.NewRestHandler(service)
	return handler
}

func startServer(engine *gin.Engine, host string, port int) error {
	httpServer, err := createHttpServer(engine, host, port)
	if err != nil {
		return err
	}
	err = httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

// Keeping an option to add error in case mTLS required in future
func createHttpServer(engine *gin.Engine, host string, port int) (*http.Server, error) {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      engine,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}, nil
}
