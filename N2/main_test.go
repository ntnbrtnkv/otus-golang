package n2

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithoutRepeatings(t *testing.T) {
	val, _ := Unpack("abc")
	require.Equal(t, "abc", val)

	val, _ = Unpack("")
	require.Equal(t, "", val)

	val, _ = Unpack("qwe\\4")
	require.Equal(t, "qwe4", val)
}

func TestWithRepeatings(t *testing.T) {
	val, _ := Unpack("a4bc2d5e")
	require.Equal(t, "aaaabccddddde", val)

	val, _ = Unpack("a4b3")
	require.Equal(t, "aaaabbb", val)

	val, _ = Unpack("qwe\\\\5")
	require.Equal(t, "qwe\\\\\\\\\\", val)
}

func TestWithZeroRepeatings(t *testing.T) {
	val, _ := Unpack("a0bc2d5e")
	require.Equal(t, "bccddddde", val)
}

func TestInvalidString(t *testing.T) {
	_, err := Unpack("34")
	require.Equal(t, errors.New("not correct string"), err)
}
