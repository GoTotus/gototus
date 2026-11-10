package totus

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTotus_Balance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/balance", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer my-secret-key", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"credits": 1234.56,
		})
	}))
	defer ts.Close()

	client, err := NewTotus("my-secret-key", ts.URL, "")
	require.NoError(t, err)

	// Direct call on client
	bal, err := client.Balance()
	require.NoError(t, err)
	assert.Equal(t, 1234.56, bal.Credits)
	assert.Contains(t, bal.String(), "1234.56")

	// Call via Utility()
	utilBal, err := client.Utility().Balance()
	require.NoError(t, err)
	assert.Equal(t, 1234.56, utilBal.Credits)
}

func TestTotus_Balance_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error": "invalid token",
		})
	}))
	defer ts.Close()

	client, err := NewTotus("invalid-key", ts.URL, "")
	require.NoError(t, err)

	bal, err := client.Balance()
	assert.Error(t, err)
	assert.Nil(t, bal)
	assert.ErrorIs(t, err, ErrAuthentication)
}

func TestTotus_Authenticated(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/authenticated", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "Bearer valid-key", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"authenticated": true,
			"description":   "your api_key is working",
			"token_sm":      "TTT_123...456",
		})
	}))
	defer ts.Close()

	client, err := NewTotus("valid-key", ts.URL, "")
	require.NoError(t, err)

	// Direct call on client
	auth, err := client.Authenticated()
	require.NoError(t, err)
	assert.True(t, auth.Authenticated)
	assert.Equal(t, "your api_key is working", auth.Description)
	assert.Equal(t, "TTT_123...456", auth.TokenSM)
	assert.Contains(t, auth.String(), "your api_key is working")

	// Call via Utility()
	utilAuth, err := client.Utility().Authenticated()
	require.NoError(t, err)
	assert.True(t, utilAuth.Authenticated)
}
