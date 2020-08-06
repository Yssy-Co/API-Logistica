package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func cotacaoTransportadoras(w http.ResponseWriter, r *http.Request) {
	// Create a span for a web request at the /posts URL.
	span := tracer.StartSpan("web.request", tracer.ResourceName("/quotar-transportadoras"))
	defer span.Finish()

	// Append span info to log messages:
	log.Printf("Iniciando cotação de transportadora em breve... %v", span)
}

func main() {
	tracer.Start(
		tracer.WithEnv("yssy-demo"),
		tracer.WithService("api-logistica"),
		tracer.WithServiceVersion(strconv.Itoa(time.Now().Day())+"-"+time.Now().Month().String()),
	)
	defer tracer.Stop()

	http.HandleFunc("/quotar-transportadoras", cotacaoTransportadoras)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
