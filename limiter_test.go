package concurrencylimiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	count := 0
	l := NewLimiter(1)
	for i := 0; i < 2; i++ {
		l.Submit(func() {
			count += 1
		})
	}
	l.Wait()
	assert.Equal(t, 2, count)
}
