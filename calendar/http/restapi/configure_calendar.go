// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/swag"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/handlers"
	"github.com/ntnbrtnkv/otus-golang/calendar/logger"
	"github.com/ntnbrtnkv/otus-golang/calendar/storage/inmemory"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/ntnbrtnkv/otus-golang/calendar/http/restapi/operations"
)

//go:generate swagger generate server --target ..\..\http --name Calendar --spec ..\..\docs\swagger.yml --principal interface{}

var options struct {
	LogLevel string `long:"log-level" description:"logging level" default:"info"`
	LogFile  string `long:"log-file" description:"logging to file" default:"logs.txt"`
}

func configureFlags(api *operations.CalendarAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Logging options",
			LongDescription:  "",
			Options:          &options,
		},
	}
}

func configureAPI(api *operations.CalendarAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	log := logger.InitLogger(logger.Config{
		LogFile:  options.LogFile,
		LogLevel: options.LogLevel,
	})

	api.Logger = log.Infof

	storage := &inmemory.InMemoryStorage{}

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	//api.UrlformConsumer = runtime.JSONConsumer()
	//
	//api.JSONProducer = runtime.JSONProducer()

	api.PostCreateEventHandler = handlers.NewCreateEventHandler(log, storage)

	api.PostUpdateEventHandler = handlers.NewUpdateEventHandler(log, storage)

	api.PostDeleteEventHandler = handlers.NewDeleteEventHandler(log, storage)

	api.GetEventsForDayHandler = handlers.NewGetEventsForDayHandler(log, storage)

	api.GetEventsForWeekHandler = handlers.NewGetEventsForWeekHandler(log, storage)

	api.GetEventsForMonthHandler = handlers.NewGetEventsForMonthHandler(log, storage)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
