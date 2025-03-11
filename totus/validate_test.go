package totus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_Email(t *testing.T) {
	client, err := NewTotus("", "", "")
	assert.NoError(t, err)
	val := client.Validate()

	resp, err := val.Email("test@example.com", CheckLevelL4Dbs)
	if err != nil {
		t.Fatalf("Email() error = %v", err)
	}
	if resp.IsValid() || resp.Score() != 5 {
		t.Errorf("Email() = %v, want valid with score 5", resp)
	}
}
