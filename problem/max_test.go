package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumber(t *testing.T) {
	t.Run("Test Max Number Array", func(t *testing.T) {
		tc1 := MaxNumberArray([]int{1, 2, 3, 8, 9, 3, 2, 1})
		tc2 := MaxNumberArray([]int{1, 2, 1, 2, 2, 1})
		tc3 := MaxNumberArray([]int{7, 1, 2, 9, 7, 2, 1})

		assert.Equal(t, 3, tc1)
		assert.Equal(t, 2, tc2)
		assert.Equal(t, 2, tc3)
	})
}
