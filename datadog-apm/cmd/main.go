package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var (
	port                = flag.String("port", "8080", "Exposed port")
	targetURL           = flag.String("targetURL", "", "Next request hop")
	datadogAgentAddress = flag.String("datadogAgentAddress", "127.0.0.1:8126", "Datadog agent address")
	timeout             = flag.Int("timeout", 0, "Timeout to fake task/work")
)

func traceCtx(r *http.Request, operationName string) (ddtrace.Span, context.Context) {
	ctx := r.Context()
	sctx, err := tracer.Extract(tracer.HTTPHeadersCarrier(r.Header))
	if err != nil {
		span, ctx := tracer.StartSpanFromContext(ctx, operationName)
		return span, ctx
	}
	span := tracer.StartSpan(operationName, tracer.ChildOf(sctx))
	return span, ctx
}

func req(ctx context.Context, span ddtrace.Span, method, url string, body io.Reader) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(span, "error", err)
		return err
	}
	req = req.WithContext(ctx)
	err = tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(req.Header))
	if err != nil {
		log.Println(span, "error", err)
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(span, "error", err)
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	log.Println(span, "body", string(b))
	return nil
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	span, ctx := traceCtx(r, "web.request")
	span.SetTag("port", *port)
	log.Println(span, "request coming in")
	defer span.Finish()
	time.Sleep(time.Duration(*timeout) * time.Second)
	if *targetURL == "" {
		w.Write([]byte("Ok"))
		return
	}
	err := req(ctx, span, "GET", *targetURL, nil)
	if err != nil {
		w.Write([]byte("error"))
	}
	w.Write([]byte("Ok"))
}

func main() {
	flag.Parse()
	tracer.Start(
		tracer.WithAgentAddr(*datadogAgentAddress),
		tracer.WithEnv("test"),
		tracer.WithService(fmt.Sprintf("host:%s", *port)),
		tracer.WithServiceVersion("v0.0.1"),
	)
	defer tracer.Stop()
	mux := httptrace.NewServeMux()
	mux.HandleFunc("/", actionHandler)
	log.Println("Server serving at port", *port)
	log.Println(*port, *targetURL, *datadogAgentAddress, *timeout)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), mux)
}
