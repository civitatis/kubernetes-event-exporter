package metrics

import (
	"testing"
	"time"
)

func TestNewMetricsStore(t *testing.T) {
	// Test creating a new metrics store
	store := NewMetricsStore("test_")
	defer DestroyMetricsStore(store)

	// Verify that the store was created
	if store == nil {
		t.Fatal("NewMetricsStore returned nil")
	}

	// Verify all metrics are initialized
	if store.EventsProcessed == nil {
		t.Error("EventsProcessed not initialized")
	}
	if store.EventsDiscarded == nil {
		t.Error("EventsDiscarded not initialized")
	}
	if store.WatchErrors == nil {
		t.Error("WatchErrors not initialized")
	}
	if store.SendErrors == nil {
		t.Error("SendErrors not initialized")
	}
	if store.BuildInfo == nil {
		t.Error("BuildInfo not initialized")
	}
	if store.KubeApiReadCacheHits == nil {
		t.Error("KubeApiReadCacheHits not initialized")
	}
	if store.KubeApiReadRequests == nil {
		t.Error("KubeApiReadRequests not initialized")
	}
	if store.LastProcessedEventTimestamp == nil {
		t.Error("LastProcessedEventTimestamp not initialized")
	}
}

func TestSetLastEventProcessedTime(t *testing.T) {
	// Save original state
	originalTime := lastEventProcessedTime
	defer func() { lastEventProcessedTime = originalTime }()

	// Initially should be zero time
	lastEventProcessedTime = time.Time{}
	if !lastEventProcessedTime.IsZero() {
		t.Error("Expected lastEventProcessedTime to be zero initially")
	}

	// Set the time
	before := time.Now()
	SetLastEventProcessedTime()
	after := time.Now()

	// Verify the time was set to something reasonable
	if lastEventProcessedTime.IsZero() {
		t.Error("Expected lastEventProcessedTime to be set after calling SetLastEventProcessedTime")
	}
	if lastEventProcessedTime.Before(before) {
		t.Error("lastEventProcessedTime should not be before the call")
	}
	if lastEventProcessedTime.After(after) {
		t.Error("lastEventProcessedTime should not be after the call")
	}
}

func TestDestroyMetricsStore(t *testing.T) {
	// Create a store
	store := NewMetricsStore("test_destroy_")
	if store == nil {
		t.Fatal("Failed to create metrics store")
	}
	if globalMetricsStore != store {
		t.Error("globalMetricsStore should be set to the new store")
	}

	// Destroy it - should not panic
	DestroyMetricsStore(store)
}

func TestMetricsStoreBasicFunctionality(t *testing.T) {
	// Integration test for basic functionality
	store := NewMetricsStore("integration_test_")
	defer DestroyMetricsStore(store)

	// Simulate event processing
	store.EventsProcessed.Inc()
	store.LastProcessedEventTimestamp.SetToCurrentTime()
	SetLastEventProcessedTime()

	// Test that the metrics are accessible
	if store.EventsProcessed == nil {
		t.Error("EventsProcessed should not be nil")
	}
	if store.LastProcessedEventTimestamp == nil {
		t.Error("LastProcessedEventTimestamp should not be nil")
	}
}
