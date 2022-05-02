package service

import "testing"

func TestGetKey(t *testing.T) {
	s := NewService()
	s.Set("foo", "bar")

	val, err := s.Get("foo")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val != "bar" {
		t.Errorf("Expected 'bar', got %v", val)
	}
}

func TestSetKey(t *testing.T) {
	s := NewService()
	s.Set("foo", "bar")

	val, err := s.Get("foo")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if val != "bar" {
		t.Errorf("Expected 'bar', got %v", val)
	}
}
