package strtypes

import "testing"

func TestUUID(t *testing.T) {
	uid := NewUUID()
	if !uid.Valid() {
		t.Errorf("expected valid UUID, got invalid")
	}
}

func TestUUID_Valid(t *testing.T) {
	s := "123e4567-e89b-12d3-a456-426614174000"
	uid := UUID(s)
	if !uid.Valid() {
		t.Errorf("expected valid UUID, got invalid")
	}
}

func TestNewUUID(t *testing.T) {
	uid := NewUUID()
	if !uid.Valid() {
		t.Errorf("expected valid UUID, got invalid")
	}
}

func TestNewUUIDPointer(t *testing.T) {
	uid := NewUUIDPointer()
	if uid == nil {
		t.Errorf("expected non-nil UUID pointer, got nil")
	}
	if !uid.Valid() {
		t.Errorf("expected valid UUID, got invalid")
	}
}

func TestUUID_Value(t *testing.T) {
	s := "123e4567-e89b-12d3-a456-426614174000"
	uid := UUID(s)

	v, err := uid.Value()
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	if _, ok := v.(string); !ok {
		t.Errorf("expected string, got %T", v)
	}

	if v.(string) != s {
		t.Errorf("expected %s, got %s", s, v)
	}
}

func TestUUID_Scan(t *testing.T) {
	var uid UUID

	s := "123e4567-e89b-12d3-a456-426614174000"
	err := uid.Scan(s)
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	if uid.String() != s {
		t.Errorf("expected %s, got %s", s, uid.String())
	}
}

func TestUUID_JSON(t *testing.T) {
	s := "123e4567-e89b-12d3-a456-426614174000"

	uid := UUID(s)
	b, err := uid.MarshalJSON()
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	var newUID UUID
	err = newUID.UnmarshalJSON(b)
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	if newUID.String() != uid.String() {
		t.Errorf("expected %s, got %s", uid.String(), newUID.String())
	}
}
