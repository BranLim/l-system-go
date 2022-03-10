package main

import "testing"

func TestBuildPoint(t *testing.T) {

	var p = BuildPoint(1, 5, 5, N)
	var next = p.next(1)
	if !(next.direction == N && next.symbol == 0 && next.x == 5 && next.y == 4) {
		t.Errorf("Wrong point, got %+v", next)
	}
}
