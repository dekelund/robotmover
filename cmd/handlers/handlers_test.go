package handlers_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dekelund/robotmover/cmd/handlers"
	"github.com/dekelund/robotmover/internal/robot/controllers"
)

func TestNewMux_control(t *testing.T) {
	cases := map[string]struct {
		body           string
		expectedBody   string
		expectedStatus int
	}{
		"example 1": {
			body: `5 5
1 2 N
RFRFFRFRF
`,
			expectedBody:   "Report: 1 3 N",
			expectedStatus: http.StatusOK,
		},
		"example 2": {
			body: `5 5
0 0 E
RFLFFLRF
`,
			expectedBody: "Report: 3 1 E",
			expectedStatus: http.StatusOK,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := controllers.New()
			mux := handlers.NewMux(c)

			srv := httptest.NewServer(mux)
			defer srv.Close()

			commands := bytes.NewBufferString(tc.body)
			req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, srv.URL+"/control", commands)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			if resp.StatusCode != tc.expectedStatus {
				t.Fatal("unexpected status:", resp.Status)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			if string(body) != tc.expectedBody {
				t.Fatal("unexpected body:", string(body))
			}
		})
	}
}

func TestNewMux_control_bad_request(t *testing.T) {
	cases := map[string]struct {
		body   string
		result int
	}{
		"invalid boundaries": {
			body: `5x 5
1 2 N
RFRFFRFRF
`,
			result: http.StatusBadRequest,
		},
		"invalid start position": {
			body: `5 5
0 0 EL
RFLFFLRF
`,
			result: http.StatusBadRequest,
		},
		"invalid instructions": {
			body: `5 5
0 0 EL
INVALID
`,
			result: http.StatusBadRequest,
		},
		"empty body": {
			body:   "",
			result: http.StatusBadRequest,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := controllers.New()
			mux := handlers.NewMux(c)

			srv := httptest.NewServer(mux)
			defer srv.Close()

			commands := bytes.NewBufferString(tc.body)
			req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, srv.URL+"/control", commands)

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			if resp.StatusCode != tc.result {
				t.Fatal("unexpected status:", resp.Status)
			}
		})
	}
}
