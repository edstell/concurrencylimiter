package concurrencylimiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	count := 0
	cl := New(1)
	for i := 0; i < 2; i++ {
		cl.Do(func() {
			count += 1
		})
	}
	cl.Wait()
	assert.Equal(t, 2, count)
}
