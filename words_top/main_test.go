package words_top

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithAtLeast10Words(t *testing.T) {
	require.Equal(t, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, Top10("0 1 2 3 4 5 6 7 8 9"))
	require.Equal(t, []string{"0", "1", "10", "2", "3", "4", "5", "6", "7", "8"}, Top10("10 0 1 2 3 4 5 6 7 8 9"))
	require.Equal(t, []string{"the", "a", "and", "dummy", "industry", "ipsum", "lorem", "of", "text", "type"}, Top10("Lorem Ipsum is simply dummy, text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."))
}

func TestWithAtLessThan10Words(t *testing.T) {
	require.Equal(t, []string{"0", "1", "2", "3", "4"}, Top10("0 1 2 3 4"))

	require.Equal(t, []string{"1", "3"}, Top10("3 1 1 3 1"))

	require.Equal(t, []string{}, Top10(""))
}

func TestCaseInsencive(t *testing.T) {
	require.Equal(t, []string{"юникод", "test"}, Top10("test Test юникод ЮнИкОд ЮНИКОД"))
}
