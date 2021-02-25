/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"log"
	"net/http"
	"time"

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
			Debug: true,
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

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("root."))
		})
		r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
			hub := sentry.GetHubFromContext(r.Context())
			hub.CaptureException(errors.New("test error"))
		})
		r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
			panic("server panic")
		})

		http.ListenAndServe("localhost:3333", r)
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
