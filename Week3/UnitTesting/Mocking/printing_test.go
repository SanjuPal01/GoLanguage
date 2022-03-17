package main

import "testing"

type SpySleeper struct {
	Calls int
}

// Here we are giving fake sleep (For Saving our time)
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestSum(t *testing.T) {
	fakeSleeper := &SpySleeper{}
	got := Sum(4, fakeSleeper)
	expected := 10
	if got != expected {
		t.Fatalf("FAILED!!!!, Expected %d, got %d\n", expected, got)
	} else if fakeSleeper.Calls != 4 {
		t.Fatalf("Failed!!!!, Expected %d calls, got %d", 4, fakeSleeper.Calls)
	} else {
		t.Logf("SUCCESS!!!!, Expected %d, Got %d\n", expected, got)
	}
}
