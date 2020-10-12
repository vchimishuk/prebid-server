package adoppler

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	bidder := NewAdopplerBidder("http://{{.AccountID}}.trustedmarketplace.com")
	adapterstest.RunJSONBidderTest(t, "adopplertest", bidder)
}
