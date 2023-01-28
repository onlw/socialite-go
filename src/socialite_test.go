package socialite

import "testing"

func TestUsage(t *testing.T) {
	New().Provider("github").RedirectTo()
}
