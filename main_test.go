package main

import "testing"

func TestGreetingWithName(t *testing.T) {
	if got := Greeting("zadig"); got != "hello, zadig" {
		t.Fatalf("Greeting() = %q, want %q", got, "hello, zadig")
	}
}

func TestGreetingWithoutName(t *testing.T) {
	if got := Greeting(""); got != "hello, anonymous" {
		t.Fatalf("Greeting() = %q, want %q", got, "hello, anonymous")
	}
}

func TestRiskyDiscountNormalCase(t *testing.T) {
	if got := RiskyDiscount(100, 20); got != 80 {
		t.Fatalf("RiskyDiscount() = %d, want %d", got, 80)
	}
}

func TestReleaseBlockerScanBugs(t *testing.T) {
	scanBugs := 6
	if IsSafeToRelease(scanBugs, 95) {
		t.Fatalf("release should be blocked when scan bugs = %d, threshold is 5", scanBugs)
	}
}

func TestReleaseBlockerTestPassRate(t *testing.T) {
	testPassRate := 88.5
	if IsSafeToRelease(1, testPassRate) {
		t.Fatalf("release should be blocked when test pass rate = %.1f%%, threshold is 90%%", testPassRate)
	}
}

func TestExpectedFailureForDemo(t *testing.T) {
	t.Fatalf("intentional failure: simulate a release test failure for AI release specialist demo")
}
