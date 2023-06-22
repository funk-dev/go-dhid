package dhid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDHID_New(t *testing.T) {
	dhid := New("de:01234:5678")
	assert.Equal(t, DHID("de:01234:5678"), dhid)

	dhid = New("de:01234:5678:9:10")
	assert.Equal(t, DHID("de:01234:5678:9:10"), dhid)
}

func TestDHID_String(t *testing.T) {
	dhid := DHID("de:01234:5678")
	assert.Equal(t, "de:01234:5678", dhid.String())

	dhid = "de:01234:5678:9:10"
	assert.Equal(t, "de:01234:5678:9:10", dhid.String())
}

func TestDHID_ToSlug(t *testing.T) {
	dhid := DHID("de:01234:5678")
	assert.Equal(t, "de-01234-5678", dhid.ToSlug())

	dhid = "de:01234:5678:9:10"
	assert.Equal(t, "de-01234-5678-9-10", dhid.ToSlug())
}

func TestDHID_IsValid(t *testing.T) {
	dhid := DHID("de:01234:5678")
	assert.Equal(t, true, dhid.IsValid())

	dhid = "de:01234:5678:9:10"
	assert.Equal(t, true, dhid.IsValid())

	dhid = ""
	assert.Equal(t, false, dhid.IsValid())
}

func TestDHID_GetStopPart(t *testing.T) {
	dhid := DHID("de:01234:5678:9:10")
	assert.Equal(t, DHID("de:01234:5678"), dhid.GetStopPart())

	dhid = "de:01234:5678:9"
	assert.Equal(t, DHID("de:01234:5678"), dhid.GetStopPart())

	dhid = "de:01234:5678"
	assert.Equal(t, DHID("de:01234:5678"), dhid.GetStopPart())
}
