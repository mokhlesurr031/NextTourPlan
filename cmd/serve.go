package cmd

import (
	"fmt"
	"github.com/NextTourPlan/internal/conn"
	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"

	_toursHttp "github.com/NextTourPlan/tours/delivery/http"
	_toursRepository "github.com/NextTourPlan/tours/repository"
	_toursUseCase "github.com/NextTourPlan/tours/usecase"

	_authHttp "github.com/NextTourPlan/auth/delivery/http"
	_authRepository "github.com/NextTourPlan/auth/repository"
	_authUseCase "github.com/NextTourPlan/auth/usecase"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starting Server...",
	Long:  `Starting Server...`,
	Run:   server,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func server(cmd *cobra.Command, args []string) {
	log.Println("Connecting database")
	if err := conn.ConnectDB(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connected successfully!")

	// boot http server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	srv := buildHTTP(cmd, args)
	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}(srv)
	<-stop
}

func buildHTTP(_ *cobra.Command, _ []string) *http.Server {
	r := chi.NewRouter()

	db := conn.DefaultDB()
	_ = db

	toursRepo := _toursRepository.New(db)
	toursUseCase := _toursUseCase.New(toursRepo)
	_toursHttp.NewHTTPHandler(r, toursUseCase)

	authRepo := _authRepository.New(db)
	authUsecase := _authUseCase.New(authRepo)
	_authHttp.NewHTTPHandler(r, authUsecase)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", 8081),
		Handler: r,
	}
}
