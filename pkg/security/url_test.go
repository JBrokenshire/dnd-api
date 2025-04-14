package security_test

import (
	"dnd-api/pkg/security"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsURLBad(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    string
		Expected bool
	}{
		{
			"Valid URL should pass",
			"https://purplemdm.com/owners",
			false,
		}, {
			"URL with filters should pass",
			"https://api.purplemdm.com/owners/activities?page_size=12&app=com.android.chrome",
			false,
		}, {
			// No idea how or why this would happen, but indicated bad config or bad load balancer activity
			"A blank URL should fail",
			"",
			true,
		}, {
			"Detect any use of single quote (') and fail ",
			"https://api.purplemdm.com/owners/activities?app=com.android.chrome '",
			true,
		},
		{
			`Detect use of double quotes (") and fail `,
			"https://api.purplemdm.com/owners/activities?app=com.android.chrome--",
			true,
		},
		{
			"Detect any use of backtick (') and fail ",
			"https://api.purplemdm.com/owners/activities?app=com.android.chrome `",
			true,
		},
		{
			"Detect use of comments (--) and fail ",
			"https://api.purplemdm.com/owners/activities?app=com.android.chrome--",
			true,
		},
		{
			//`userName = <script src="http://localhost:2001/Attackerscript.js"></script>`
			"Detect use of injecting script (script) and fail ",
			"https://api.purplemdm.com/owners/activities?app=<script",
			true,
		},
		{
			"Detect use of injecting script (src=) and fail ",
			"https://api.purplemdm.com/owners/activities?app=src=",
			true,
		},
		{
			"Detect use of oastify in URL. THis is a burpsuite callback",
			"https://api.purplemdm.com/roles/1?test.oastify.com%5c%5",
			true,
		},
		{
			Name:     "Can allow /scribe-transcriptions route through script check",
			Input:    "https://api.purplemdm.com/scribe-transcriptions/10",
			Expected: false,
		},
		{
			Name:     "Cannot allow /scribe-transcriptions route through with other bad word",
			Input:    "https://api.purplemdm.com/scribe-transcriptions/10?app=src=",
			Expected: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			got := security.IsURLBad(test.Input)
			assert.Equal(t, test.Expected, got)
		})
	}
}

func BenchmarkIsURLBad(b *testing.B) {
	// 85.21 ns/op
	for n := 0; n < b.N; n++ {
		security.IsURLBad("https://api.purplemdm.com")
	}
}
