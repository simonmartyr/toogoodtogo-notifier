package main

import (
	"testing"
	"time"
)

func TestShouldNotify_true_1day(t *testing.T) {
	yesterday := time.Now().Add(-(time.Hour * 24))
	if !ShouldNotify(yesterday.Format("2006-01-02")) {
		t.Error("ShouldNotify should be true")
	}
}

func TestShouldNotify_true_10days(t *testing.T) {
	tenDaysAgo := time.Now().Add(-(time.Hour * 24 * 10))
	if !ShouldNotify(tenDaysAgo.Format("2006-01-02")) {
		t.Error("ShouldNotify should be true")
	}
}

func TestShouldNotify_true_empty(t *testing.T) {
	if !ShouldNotify("") {
		t.Error("ShouldNotify should be true")
	}
}

func TestShouldNotify_false_same_day(t *testing.T) {
	yesterday := time.Now()
	if ShouldNotify(yesterday.Format("2006-01-02")) {
		t.Error("ShouldNotify should be false")
	}
}
