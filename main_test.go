package aldoutil

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	setupTest()
	code := m.Run()
	shutdownTest()

	os.Exit(code)
}

func setupTest() {
}

func shutdownTest() {
}

func TestStatus(t *testing.T) {
	now := time.Now()
	validDate := now.Add(-time.Hour * 24 * 30)

	// New.
	product := Product{
		Availability: true,
		CreatedAt:    now,
		ChangedAt:    now,
	}
	result := product.Status(validDate)
	want := "new"
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}

	// Include on Zunka site.
	product.MongodbId = "123456789012345678901234"
	product.StatusCleanedAt = product.ChangedAt.Add(time.Hour * 1)
	result = product.Status(validDate)
	want = ""
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}

	// Modified.
	product.ChangedAt = product.StatusCleanedAt.Add(time.Hour * 1)
	result = product.Status(validDate)
	want = "changed"
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}

	// Cleaned.
	product.StatusCleanedAt = product.ChangedAt.Add(time.Minute * 1)
	result = product.Status(validDate)
	want = ""
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}

	// Unavailable.
	product.Availability = false
	result = product.Status(validDate)
	want = "unavailable"
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}

	// Removed.
	product.RemovedAt = now
	result = product.Status(validDate)
	want = "removed"
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}

	// Expired status.
	expiredStatus := now.Add(-time.Hour * 24 * 31)
	product = Product{
		Availability: true,
		CreatedAt:    expiredStatus,
		ChangedAt:    expiredStatus,
	}
	result = product.Status(validDate)
	want = ""
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}

	// changed.
	product.CreatedAt = now
	product.StatusCleanedAt = now.Add(time.Hour * 1)
	product.ChangedAt = now.Add(time.Hour * 2)
	result = product.Status(validDate)
	want = "changed"
	if result != want {
		t.Errorf("result = %q, want %q", result, want)
	}
}
