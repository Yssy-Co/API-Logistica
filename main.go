package main

import (
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func cotacaoTransportadoras(w http.ResponseWriter, r *http.Request) {
	sctx, err := tracer.Extract(tracer.HTTPHeadersCarrier(r.Header))

	span := tracer.StartSpan("web.request", tracer.ResourceName("/quotar-transportadoras"), tracer.ChildOf(sctx))
	defer span.Finish()

	if err != nil {
		log.WithField("dd", span).Errorln(err)
	}
	log.WithField("dd", span).Println("Iniciando cotação de transportadora em breve")

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
