package random_test

import (
	"dnd-api/pkg/random"
	"fmt"
	"testing"
)

func TestClientSecret(t *testing.T) {
	secret, err := random.ClientSecret()
	if err != nil {
		t.Errorf("Error generating secret: %v", err)
	}
	if len(secret) != 72 {
		t.Errorf("Expected length of 128, got %v", len(secret))
	}
	fmt.Println("ClientSecret:", secret)
}

func TestUidV4(t *testing.T) {
	Uid := random.UidV4()

	if len(Uid) != 36 {
		t.Errorf("Expected length of 36, got %v", len(Uid))
	}
	fmt.Println("UidV4:", Uid)
}

func TestStringFixed(t *testing.T) {
	str := random.StringFixed(32)
	if len(str) != 32 {
		t.Errorf("Expected length of 10, got %v", len(str))
	}
	fmt.Println("StringFixed:", str)

	str64 := random.StringFixed(64)
	if len(str64) != 64 {
		t.Errorf("Expected length of 64, got %v", len(str64))
	}
	fmt.Println("StringFixed64:", str64)
}
