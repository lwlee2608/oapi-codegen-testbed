package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	Spec "oapi-codegen-testbed/pkg/generated"

	"github.com/gdexlab/go-render/render"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type SpecServer struct {
}

func (m *SpecServer) Get(ctx echo.Context) error {
	r := ctx.Request()
	reqBody, _ := ioutil.ReadAll(r.Body)
	prettyJson, _ := prettyprint(reqBody)
	log.Infof("request json %s", prettyJson)

	request := Spec.RequestObject{}
	json.Unmarshal(reqBody, &request)
	log.Infof("request object %s", render.AsCode(request))

	return ctx.String(http.StatusOK, "")
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	swagger, err := Spec.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	specServer := SpecServer{}

	e := echo.New()
	// e.Use(echomiddleware.Logger())

	//e.Use(middleware.OapiRequestValidator(swagger))

	Spec.RegisterHandlers(e, &specServer)

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
