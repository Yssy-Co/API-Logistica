package main

import (
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func cotacaoTransportadoras(w http.ResponseWriter, r *http.Request) {
	// Create a span for a web request at the /posts URL.
	span := tracer.StartSpan("web.request", tracer.ResourceName("/quotar-transportadoras"))
	defer span.Finish()

	// Append span info to log messages:
	//log.Info().Msg("Olá")
	//log.Printf("Iniciando cotação de transportadora em breve... %v", span)

	log.WithField("dd", span).Print("Iniciando cotação de transportadora em breve")
	//log.Log().Interface("Iniciando cotação de transportadora em breve... %v", span)

}

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.SetReportCaller(true)

	log.SetFormatter(&log.JSONFormatter{})

	tracer.Start(
		tracer.WithEnv("yssy-demo"),
		tracer.WithService("api-logistica"),
		tracer.WithServiceVersion(strconv.Itoa(time.Now().Day())+"-"+time.Now().Month().String()),
	)
	defer tracer.Stop()

	http.HandleFunc("/quotar-transportadoras", cotacaoTransportadoras)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
