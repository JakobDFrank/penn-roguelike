package model

import (
	"testing"
)

func TestNewDirection(t *testing.T) {

	t.Run("test_valid_direction", func(t *testing.T) {
		for _, dir := range DirectionNameToEnum {

			if _, err := NewDirection(int(dir)); err != nil {
				t.Fatalf("NewDirection on a valid Direction returned an error: %v", err)
			}

		}
	})

	t.Run("test_invalid_direction", func(t *testing.T) {
		if _, err := NewDirection(int(Down) + 1); err == nil {
			t.Fatal("NewDirection on an invalid Direction id not return an error")
		}
	})
}

func TestDirection_IsValid(t *testing.T) {

	t.Run("test_valid_direction", func(t *testing.T) {
		for _, dir := range DirectionNameToEnum {
			if !dir.IsValid() {
				t.Fatalf("IsValid on a valid Direction returned false: %d", dir)
			}
		}
	})

	t.Run("test_invalid_direction", func(t *testing.T) {
		dir := Direction(int(Down) + 1)

		if dir.IsValid() {
			t.Fatalf("IsValid on an invalid Direction returned true: %d", dir)
		}
	})
}

func TestPlayer_ResetHitpoints(t *testing.T) {
	plyer := NewPlayer(0, 0, 0)

	plyer.Hitpoints = 0

	plyer.ResetHitpoints()

	if plyer.Hitpoints != _startingHitpoints {
		t.Fatalf("player hitpoints did not reset. hitpoints: %d", plyer.Hitpoints)
	}
}
