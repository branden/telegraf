package dcos_metrics

import (
	"testing"
)

func TestSplitHostPort(t *testing.T) {
	type testCase struct {
		hostPort string
		host     string
		port     int
	}

	testCases := []testCase{
		testCase{
			hostPort: "localhost:8000",
			host:     "localhost",
			port:     8000,
		},
		testCase{
			hostPort: "10.10.0.1:10",
			host:     "10.10.0.1",
			port:     10,
		},
		testCase{
			hostPort: "host.name.com:60000",
			host:     "host.name.com",
			port:     60000,
		},
		testCase{
			hostPort: ":80",
			host:     "",
			port:     80,
		},
	}

	for _, tc := range testCases {
		host, port, err := splitHostPort(tc.hostPort)
		if err != nil {
			t.Fatal(err)
		}

		if host != tc.host {
			t.Fatalf("expected host %s for hostport %s, got %s", tc.host, tc.hostPort, host)
		}
		if port != tc.port {
			t.Fatalf("expected port %d for hostport %s, got %d", tc.port, tc.hostPort, port)
		}
	}

	badPortCases := []string{
		"localhost:foo",
		"localhost:",
	}
	for _, hostPort := range badPortCases {
		_, _, err := splitHostPort(hostPort)
		if err == nil {
			t.Fatalf("expected error for hostport %s", hostPort)
		}
	}
}
