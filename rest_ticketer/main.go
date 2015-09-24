package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/vosst/surl"
)

type RestTicketerService struct {
}

func (self *RestTicketerService) Main(host string, rw surl.CounterReaderWriter) {
	if wt, err := surl.NewWriteAheadTicketer(rw); err != nil {
		log.Fatal(err)
	} else {
		http.Handle(surl.RestTicketerNextEndpoint, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(surl.RestTicketerDocument{wt.Next()})
		}))
		log.Fatal(http.ListenAndServe(host, nil))
	}
}

var (
	flagHost = flag.String("host", ":9090", "Host address to bind to")
	flagRW   = flag.String("rw", "file", "Configures the persistence backend, only file is supported right now")
)

func createCounterReaderWriter(name string) surl.CounterReaderWriter {
	switch name {
	case "file":
		return &surl.FileCounterReaderWriter{"/var/surl/counter"}
	default:
		return nil
	}
}

func main() {
	rts := &RestTicketerService{}
	rts.Main(":9090", createCounterReaderWriter(*flagRW))
}
