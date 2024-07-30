package reference

import "testing"

func TestHasElements(t *testing.T) {
	value := uint(36)
	want := []string{"Lightning", "Holy"}
	sut := ContainsElements(Elements(value))

	if sut[0] != want[0] {
		t.Fatalf("wanted %s got %s", want[0], sut[0])
	}

	if sut[1] != want[1] {
		t.Fatalf("wanted %s got %s", want[1], sut[1])
	}
}

func TestHasConditions(t *testing.T) {
	value := uint32(16717824)
	want := []string{"Silence", "Berserk", "Flying", "Regen", "Slow", "Haste", "Stop", "Shell", "Safe", "Reflect"}
	sut := ContainsConditions(Conditions(value))

	if sut[0] != want[0] {
		t.Fatalf("wanted %s got %s", want[0], sut[0])
	}

	if sut[2] != want[2] {
		t.Fatalf("wanted %s got %s", want[2], sut[2])
	}

	if sut[6] != want[6] {
		t.Fatalf("wanted %s got %s", want[6], sut[6])
	}
}

func TestHasFlags(t *testing.T) {
	value := uint32(17537)
	want := []string{"Dies at 0MP", "Undead", "Can't Suplex", "Special Event"}
	sut := ContainsFlags(MiscFlags(value))

	if sut[0] != want[0] {
		t.Fatalf("wanted %s got %s", want[0], sut[0])
	}

	if sut[1] != want[1] {
		t.Fatalf("wanted %s got %s", want[1], sut[1])
	}

	if sut[2] != want[2] {
		t.Fatalf("wanted %s got %s", want[2], sut[2])
	}

	if sut[3] != want[3] {
		t.Fatalf("wanted %s got %s", want[3], sut[3])
	}
}
