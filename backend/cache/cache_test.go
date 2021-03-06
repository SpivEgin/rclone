// Test Cache filesystem interface

// +build !plan9

package cache_test

import (
	"testing"

	"github.com/artpar/rclone/backend/cache"
	_ "github.com/artpar/rclone/backend/local"
	"github.com/artpar/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestCache:",
		NilObject:  (*cache.Object)(nil),
	})
}
