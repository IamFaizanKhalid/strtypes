package time

import "testing"

const sampleTime = "2024-04-07T10:35:00Z"

func TestTime(t *testing.T) {
	uid := NewTime()
	if !uid.Valid() {
		t.Errorf("expected valid Time, got invalid")
	}
}

func TestTime_Valid(t *testing.T) {
	uid := Time(sampleTime)
	if !uid.Valid() {
		t.Errorf("expected valid Time, got invalid")
	}
}

func TestNewTime(t *testing.T) {
	uid := NewTime()
	if !uid.Valid() {
		t.Errorf("expected valid Time, got invalid")
	}
}

func TestNewTimePointer(t *testing.T) {
	uid := NewTimePointer()
	if uid == nil {
		t.Errorf("expected non-nil Time pointer, got nil")
	}
	if !uid.Valid() {
		t.Errorf("expected valid Time, got invalid")
	}
}

func TestTime_Value(t *testing.T) {
	uid := Time(sampleTime)

	v, err := uid.Value()
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	if _, ok := v.(string); !ok {
		t.Errorf("expected string, got %T", v)
	}

	if v.(string) != sampleTime {
		t.Errorf("expected %s, got %s", sampleTime, v)
	}
}

func TestTime_Scan(t *testing.T) {
	var uid Time

	err := uid.Scan(sampleTime)
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	if uid.String() != sampleTime {
		t.Errorf("expected %s, got %s", sampleTime, uid.String())
	}
}

func TestTime_JSON(t *testing.T) {
	uid := Time(sampleTime)
	b, err := uid.MarshalJSON()
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	var newUID Time
	err = newUID.UnmarshalJSON(b)
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	if newUID.String() != uid.String() {
		t.Errorf("expected %s, got %s", uid.String(), newUID.String())
	}
}
