package painter

import "testing"

func TestNextNorthPoint(t *testing.T) {

	var p = BuildPoint('1', 5, 5, N)
	var next = p.Next('1')
	if !(next.Direction == N && next.Symbol == '1' && next.X == 5 && next.Y == 4) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextSouthPoint(t *testing.T) {

	var p = BuildPoint('1', 5, 5, S)
	var next = p.Next('1')
	if !(next.Direction == S && next.Symbol == '1' && next.X == 5 && next.Y == 6) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextNEPoint(t *testing.T) {

	var p = BuildPoint('1', 5, 5, NE)
	var next = p.Next('1')
	if !(next.Direction == NE && next.Symbol == '1' && next.X == 6 && next.Y == 4) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextSEPoint(t *testing.T) {

	var p = BuildPoint('1', 5, 5, SE)
	var next = p.Next('1')
	if !(next.Direction == SE && next.Symbol == '1' && next.X == 6 && next.Y == 6) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextNWPoint(t *testing.T) {

	var p = BuildPoint('1', 5, 5, NW)
	var next = p.Next('1')
	if !(next.Direction == NW && next.Symbol == '1' && next.X == 4 && next.Y == 4) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextSWPoint(t *testing.T) {

	var p = BuildPoint('1', 5, 5, SW)
	var next = p.Next('1')
	if !(next.Direction == SW && next.Symbol == '1' && next.X == 4 && next.Y == 6) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextForZero(t *testing.T) {

	var p = BuildPoint('1', 5, 5, SW)
	var next = p.Next('0')
	if !(next.Direction == SW && next.Symbol == '0' && next.X == 5 && next.Y == 5) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextNWithRotateLeft(t *testing.T) {
	var p = BuildPoint('1', 5, 5, N)
	var next = p.Next('[')
	if !(next.Direction == NW && next.Symbol == '1' && next.X == 4 && next.Y == 5) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextSWithRotateLeft(t *testing.T) {
	var p = BuildPoint('1', 5, 5, S)
	var next = p.Next('[')
	if !(next.Direction == SE && next.Symbol == '1' && next.X == 4 && next.Y == 5) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextSWithRotateRight(t *testing.T) {
	var p = BuildPoint('1', 5, 5, S)
	var next = p.Next(']')
	if !(next.Direction == SW && next.Symbol == '1' && next.X == 4 && next.Y == 5) {
		t.Errorf("Wrong point, got %+v", next)
	}
}

func TestNextNWithRotateRight(t *testing.T) {
	var p = BuildPoint('1', 5, 5, NW)
	var next = p.Next(']')
	if !(next.Direction == N && next.Symbol == '1' && next.X == 4 && next.Y == 5) {
		t.Errorf("Wrong point, got %+v", next)
	}
}
