package swissknife_gen //nolint

import (
	swissknife "github.com/pmg-tools/adapterkit-module-swissknife"
)

type Request struct {
	RequestType string                    `json:"requestType"`
	ConvHexa    *swissknife.ConvHexaReq   `json:"ConvHexa"`
	ConvBase64  *swissknife.ConvBase64Req `json:"ConvBase64"`
}

type Response struct {
	ConvHexa   *swissknife.ConvHexaRes   `json:"ConvHexa"`
	ConvBase64 *swissknife.ConvBase64Res `json:"ConvBase64"`
}

func SvcConvHexa(input string, svc swissknife.SwissknifeSvcServer) (Response, error) {

	res, err := svcConvHexa(input, svc)

	if err != nil {
		return Response{}, err
	}

	return Response{ConvHexa: res}, nil
}

func SvcConvBase64(input string, svc swissknife.SwissknifeSvcServer) (Response, error) {

	res, err := svcConvBase64(input, svc)

	if err != nil {
		return Response{}, err
	}

	return Response{ConvBase64: res}, nil
}
