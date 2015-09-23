package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"github.com/vosst/surl"
	"log"
	"log/syslog"
	"net/http"
	"net/url"
	"os"
)

// Configuration bundles all parameters of the app.
type Configuration struct {
	Host        string // Host address that the service should bind to.
	Base        string // Base url for shortened URLs.
	LogToSyslog bool   // Enables logging to the os's logging facilities.
}

type Result struct {
	ShortUrl string `json:"ShortUrl"`
}

const (
	longURL            = `LongUrl`
	missingParameter   = `{reason:"Missing or invalid parameter in request"}`
	failedToShortenURL = `{reason:"Failed to shorten the input URL"}`
	urlUnknown         = `{reason:"No URL known for the input URL"}`
)

var configuration Configuration = Configuration{}

func init() {
	flag.StringVar(&configuration.Host, "host", ":9090", "Host address that the service binds to")
	flag.StringVar(&configuration.Base, "base", "http://localhost:9090", "Base URL for shortened URLs")
	flag.BoolVar(&configuration.LogToSyslog, "log-to-syslog", true, "Enables logging to the syslog if set to true")
}

func normalizeURL(url *url.URL) *url.URL {
	if len(url.Scheme) == 0 {
		url.Scheme = "http"
	}

	return url
}

func initLogger() *log.Logger {
	if configuration.LogToSyslog {
		return syslog.NewLogger(syslog.LOG_INFO, log.LstdFlags)
	} else {
		return log.New(os.Stdout, "surl ", log.LstdFlags)
	}
}

func main() {
	flag.Parse()

	ticketer := &surl.CountingTicketer{}
	store := surl.NewInMemoryStore()
	reporter := &surl.LoggingServiceReporter{initLogger()}

	service := surl.NewService(ticketer, store, reporter)
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if l, err := url.Parse(r.FormValue(longURL)); err != nil {
			http.Error(w, missingParameter, http.StatusBadRequest)
		} else {
			l = normalizeURL(l)
			if s, err := service.Put(l); err != nil {
				http.Error(w, failedToShortenURL, http.StatusBadRequest)
			} else {
				json.NewEncoder(w).Encode(Result{ShortUrl: configuration.Base + "/" + s})
			}
		}
	}).Methods("POST").Headers("Accept", "application/json")

	rtr.HandleFunc(`/{id:\w+}`, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if s, err := service.Get(r.URL.Path[1:]); err != nil {
			http.Error(w, urlUnknown, http.StatusBadRequest)
		} else {
			http.Redirect(w, r, s.String(), http.StatusMovedPermanently)
		}
	}).Methods("GET").Headers("Accept", "application/json")

	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
