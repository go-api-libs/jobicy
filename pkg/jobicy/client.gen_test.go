// This file provides tests for the client package.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package jobicy_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-api-libs/api"
	"github.com/go-api-libs/jobicy/pkg/jobicy"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

type testRoundTripper struct {
	rsp *http.Response
	err error
}

func (t *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.rsp, t.err
}

func TestClient_Error(t *testing.T) {
	ctx := context.Background()

	c, err := jobicy.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Do", func(t *testing.T) {
		testErr := errors.New("test error")
		http.DefaultClient.Transport = &testRoundTripper{err: testErr}

		if _, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
			Count:    20,
			Geo:      "usa",
			Industry: "marketing",
			Tag:      "seo",
		}); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, testErr) {
			t.Fatalf("want: %v, got: %v", testErr, err)
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		errDecode := &api.DecodingError{}

		t.Run("ListRemoteJobs", func(t *testing.T) {
			// unknown status code
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{StatusCode: http.StatusTeapot}}

			if _, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
				Count:    20,
				Geo:      "usa",
				Industry: "marketing",
				Tag:      "seo",
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownStatusCode) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownStatusCode, err)
			}

			// unknown content type for 200 OK
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Header:     http.Header{"Content-Type": []string{"foo"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
				Count:    20,
				Geo:      "usa",
				Industry: "marketing",
				Tag:      "seo",
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownContentType) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownContentType, err)
			}

			// decoding error for known content type "application/json"
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Body:       io.NopCloser(strings.NewReader("{")),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
				Count:    20,
				Geo:      "usa",
				Industry: "marketing",
				Tag:      "seo",
			}); err == nil {
				t.Fatal("expected error")
			} else if !errors.As(err, &errDecode) {
				t.Fatalf("want: %v, got: %v", errDecode, err)
			}
		})
	})
}

func replay(t *testing.T, cassette string) {
	t.Helper()

	r, err := recorder.NewWithOptions(&recorder.Options{
		CassetteName:       cassette,
		Mode:               recorder.ModeReplayOnly,
		RealTransport:      http.DefaultTransport,
		SkipRequestLatency: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = r.Stop()
	})

	r.SetMatcher(matcher)
	http.DefaultClient.Transport = r
}

func matcher(r *http.Request, i cassette.Request) bool {
	if !cassette.DefaultMatcher(r, i) {
		return false
	}

	return getBody(r) == i.Body
}

func getBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}

	if r.GetBody == nil {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewReader(b))
		return string(b)
	}

	body, err := r.GetBody()
	if err != nil {
		panic(err)
	}
	defer body.Close()

	b, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func TestClient_VCR(t *testing.T) {
	ctx := context.Background()

	c, err := jobicy.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("2024-12-05", func(t *testing.T) {
		replay(t, "../../pkg/jobicy/vcr/2024-12-05")

		res, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
			Count:    20,
			Geo:      "usa",
			Industry: "marketing",
			Tag:      "seo",
		})
		if err != nil {
			t.Fatal(err)
		} else if res == nil {
			t.Fatal("result is nil")
		}
	})

	t.Run("2024-12-08", func(t *testing.T) {
		replay(t, "../../pkg/jobicy/vcr/2024-12-08")

		{
			res, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
				Count: 20,
				Tag:   "python",
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
				Count: 15,
				Geo:   "canada",
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
				Count:    30,
				Geo:      "usa",
				Industry: "copywriting",
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
				Count:    10,
				Industry: "supporting",
			})
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}
	})
}
