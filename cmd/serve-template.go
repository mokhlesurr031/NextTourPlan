package cmd

import (
	"fmt"
	"github.com/NextTourPlan/internal/conn"
	_tourTemplate "github.com/NextTourPlan/tours/delivery/template"
	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// serveCmd represents the serve command
var serveTemplateCmd = &cobra.Command{
	Use:   "serve-template",
	Short: "Starting Server for template...",
	Long:  `Starting Server for template...`,
	Run:   templateServer,
}

func init() {
	rootCmd.AddCommand(serveTemplateCmd)
}

func templateServer(cmd *cobra.Command, args []string) {
	log.Println("Connecting database for template")
	if err := conn.ConnectDB(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connected successfully!")

	// boot http server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	srv := buildHTTPTemplate(cmd, args)
	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}(srv)
	<-stop
}

func buildHTTPTemplate(_ *cobra.Command, _ []string) *http.Server {
	r := chi.NewRouter()
	db := conn.DefaultDB()
	_ = db

	_tourTemplate.NewHTTPHandler(r)
	//_tourTemplate.New(db)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r,
	}
}
