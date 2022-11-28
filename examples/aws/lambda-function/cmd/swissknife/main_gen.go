package main

import (
	"fmt"

	"test/pkg/swissknife_gen"
	"github.com/aws/aws-lambda-go/lambda"
	swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)


func initSvc() swissknife.SwissknifeSvcServer {
	return swissknife.New()
}


func HandleRequest(req swissknife_gen.Request) (swissknife_gen.Response, error) {
    switch req.RequestType {
    case "ConvHexa":
        return swissknife_gen.SvcConvHexa(req.ConvHexa.Input, initSvc())
    case "ConvBase64":
        return swissknife_gen.SvcConvBase64(req.ConvBase64.Input, initSvc())
	default:
		return swissknife_gen.Response{}, fmt.Errorf("unknow request type, %s", req.RequestType)
    }
}

func main() {
	lambda.Start(HandleRequest)
}
