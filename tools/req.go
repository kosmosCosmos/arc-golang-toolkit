package tools

import (
	"fmt"
	"github.com/imroc/req/v3"
	"strings"
)

func NewRequest(method, url string, header, payload map[string]string) (*req.Response, string, error) {
	client := req.C().SetCommonHeaders(header)

	r := client.R()
	if payload != nil {
		r.SetFormData(payload)
	}

	var resp *req.Response
	var err error

	switch strings.ToUpper(method) {
	case "GET":
		resp, err = r.Get(url)
	case "POST":
		resp, err = r.Post(url)
	case "PUT":
		resp, err = r.Put(url)
	case "DELETE":
		resp, err = r.Delete(url)
	case "PATCH":
		resp, err = r.Patch(url)
	default:
		return nil, "", fmt.Errorf("unsupported method: %s", method)
	}

	if err != nil {
		return nil, "", fmt.Errorf("request failed: %w", err)
	}

	if !resp.IsSuccessState() {
		return nil, "", fmt.Errorf("request did not return OK. Status: %s, Body: %s", resp.Status, resp.String())
	}

	return resp, resp.String(), nil
}
