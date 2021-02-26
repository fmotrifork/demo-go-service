package cmd

import (
	"errors"
	"log"
	"net/http"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	// Load auto generated swagger docs
	_ "github.com/fmotrifork/demo-go-service/swaggerdocs"
	"github.com/fmotrifork/demo-go-service/version"

	"github.com/fmotrifork/demo-go-service/utils"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := sentry.Init(sentry.ClientOptions{
			Debug:   true,
			Release: version.Version,
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
		defer sentry.Flush(time.Second)

		sentryMiddleware := sentryhttp.New(sentryhttp.Options{
			Repanic: true,
		})

		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		// Important: Chi has a middleware stack and thus it is important to put the
		// Sentry handler on the appropriate place. If using middleware.Recoverer,
		// the Sentry middleware must come afterwards (and configure it with
		// Repanic: true).
		r.Use(sentryMiddleware.Handle)

		r.Mount("/swagger", httpSwagger.WrapHandler)
		r.Get("/", RootHandler)
		r.Get("/error", ErrorHandler)
		r.Get("/panic", PanicHandler)

		address := "localhost:3333"
		log.Printf("Starting webserver on: %s", address)
		err = http.ListenAndServe(address, r)
		log.Printf("Serving err: %s", err)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type response struct {
	Message    string `example:"All is good"`
	StatusCode int    `example:"200"`
}

// RootHandler - Returns all the available APIs
// @Summary This API can be used as health check for this application.
// @Description Tells if the chi-swagger APIs are working or not.
// @Accept  json
// @Produce  json
// @Success 200 {object} response "api response"
// @Router / [get]
func RootHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, utils.SuccessResponse, &response{
		Message:    "Working OK",
		StatusCode: 200,
	})
}

// ErrorHandler - Always return errors. ;)
// @Summary This API always returns an error, and sends an error report to Sentry.io
// @Description
// @Tags error
// @Accept  json
// @Produce  json
// @Success 200 {object} response "api response"
// @Failure 500 {object} utils.ErrorResponseModel "api response"
// @Router /error [get]
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Internal Server Error")

	hub := sentry.GetHubFromContext(r.Context())
	hub.CaptureException(err)

	utils.WriteResponse(w, utils.ErrorResponse, err)
}

// PanicHandler - Always Panics ! :o
// @Summary This API always panics, and sends a stack trace to Sentry.io
// @Description
// @Tags error
// @Accept  json
// @Produce  json
// @Success 200 {object} response "api response"
// @Failure 500 {object} utils.ErrorResponseModel "api response"
// @Router /panic [get]
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("server panic")
}
