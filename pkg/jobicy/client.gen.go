// This file provides a client with methods as well as functions to interact with the HTTP API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package jobicy

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/MarkRosemaker/jsonutil"
	"github.com/go-api-libs/api"
	"github.com/go-json-experiment/json"
)

var (
	baseURL = &url.URL{
		Host:   "jobicy.com",
		Path:   "/api/v2",
		Scheme: "https",
	}

	jsonOpts = json.JoinOptions(
		json.RejectUnknownMembers(true),
		json.WithMarshalers(json.MarshalFuncV2(jsonutil.URLMarshal)),
		json.WithUnmarshalers(json.UnmarshalFuncV2(jsonutil.URLUnmarshal)))
)

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The HTTP client to use for requests.
	cli *http.Client
}

// NewClient creates a new Client.
func NewClient() (*Client, error) {
	return &Client{cli: http.DefaultClient}, nil
}

// ListRemoteJobs defines an operation.
//
//	GET /remote-jobs
func (c *Client) ListRemoteJobs(ctx context.Context, params *ListRemoteJobsParams) (*JobsResponse, error) {
	return ListRemoteJobs[JobsResponse](ctx, c, params)
}

// ListRemoteJobs defines an operation.
// You can define a custom result to unmarshal the response into.
//
//	GET /remote-jobs
func ListRemoteJobs[R any](ctx context.Context, c *Client, params *ListRemoteJobsParams) (*R, error) {
	u := baseURL.JoinPath("/remote-jobs")

	if params != nil {
		q := make(url.Values, 4)

		if params.Count != 0 {
			q["count"] = []string{strconv.Itoa(params.Count)}
		}

		if params.Geo != "" {
			q["geo"] = []string{params.Geo}
		}

		if params.Industry != "" {
			q["industry"] = []string{params.Industry}
		}

		if params.Tag != "" {
			q["tag"] = []string{params.Tag}
		}

		u.RawQuery = q.Encode()
	}

	req := (&http.Request{
		Header:     http.Header{},
		Host:       u.Host,
		Method:     http.MethodGet,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        u,
	}).WithContext(ctx)

	rsp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		// TODO
		switch rsp.Header.Get("Content-Type") {
		case "application/json":
			var out R
			if err := json.UnmarshalRead(rsp.Body, &out, jsonOpts); err != nil {
				return nil, api.WrapDecodingError(rsp, err)
			}

			return &out, nil
		default:
			return nil, api.NewErrUnknownContentType(rsp)
		}
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}
