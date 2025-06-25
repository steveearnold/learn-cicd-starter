package auth

import "testing"

func TestReturnZero(t *testing.T) {
    result := ReturnZero() // ✅ no `auth.` needed
    if result != 0 {
        t.Errorf("Expected 0, got %d", result)
    }
}
