package api

import (
	"encoding/json"
	"testing"
)

// TestPPPDecode verifies the PPP struct extracts the PPPoE credentials from a
// real "NMC get" response. The "ppp" MIB strips the password, so NMC is used
// instead; this guards against a regression back to the password-less endpoint.
func TestPPPDecode(t *testing.T) {
	const nmcResponse = `{"status":{"WanMode":"Ethernet_PPP","Username":"user@example.com","Password":"s3cr3t","IPTVMode":"Dedicated"}}`

	var ppp PPP
	if err := json.Unmarshal([]byte(nmcResponse), &ppp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if got, want := ppp.Status.Username, "user@example.com"; got != want {
		t.Errorf("Username = %q, want %q", got, want)
	}
	if got, want := ppp.Status.Password, "s3cr3t"; got != want {
		t.Errorf("Password = %q, want %q", got, want)
	}
}

// TestPPPDecodeEmpty ensures an absent password decodes to the empty string
// rather than erroring, so the caller can fall back to the reset-URL hint.
func TestPPPDecodeEmpty(t *testing.T) {
	var ppp PPP
	if err := json.Unmarshal([]byte(`{"status":{"Username":"user"}}`), &ppp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if ppp.Status.Password != "" {
		t.Errorf("Password = %q, want empty", ppp.Status.Password)
	}
}
