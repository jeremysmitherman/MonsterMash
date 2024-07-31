package retroarch

import "testing"

func TestMonsterRAMString(t *testing.T) {
	monsters := ParseRetroArchReturn("READ_CORE_RAM 2001 B9 00 F1 05 B9 00 FF FF 15 00 03 00")

	if monsters[0] != 185 {
		t.Fatalf("Expected 185, got %d", monsters[0])
	}

	if monsters[1] != 21 {
		t.Fatalf("Expected 21, got %d", monsters[4])
	}

	if monsters[2] != 3 {
		t.Fatalf("Expected 3, got %d", monsters[5])
	}

}
