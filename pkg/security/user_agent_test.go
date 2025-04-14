package security_test

import (
	"dnd-api/pkg/security"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUserAgentBad(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    string
		Expected bool
	}{
		{
			"Valid Chrome should pass",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
			false,
		}, {
			"Valid Firefox should pass",
			"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:107.0) Gecko/20100101 Firefox/107.0",
			false,
		}, {
			"Edge on Win10 should pass",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.62",
			false,
		}, {
			"Brave on Win10 should pass",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			false,
		}, {
			// a user agent SHOULD be sent: https://www.rfc-editor.org/rfc/rfc7231#section-5.5.3
			"A blank user agent should fail",
			"",
			true,
		}, {
			"burp collaborator should be blocked",
			"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36 root@8f6gtjctw9glvnasfrq438tprig56wty20ghwkf9.burpcollaborator.net",
			true,
		}, {
			"block SSL Labs",
			"SSL Labs (https://www.ssllabs.com/about/assessment.html)",
			true,
		}, {
			"block openvas",
			"Mozilla/5.0 [en] (X11, U; OpenVAS-VT 9.0.3)",
			true,
		}, {
			"block nessus",
			"Mozilla/5.0 [en] (X11, U; OpenVAS-VT 9.0.3)",
			true,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			got := security.IsUserAgentBad(test.Input, false)
			assert.Equal(t, test.Expected, got)
		})
	}

}

func BenchmarkIsUserAgentBad(b *testing.B) {
	// with 700 bad agents this is 9549 ns/op
	for n := 0; n < b.N; n++ {
		security.IsUserAgentBad("A random string that shouldn't trigger. Example of GOOD", false)
	}
}
