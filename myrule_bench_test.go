package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkParseRule(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree, got, err := Parse(`((app-id in ["com.github.app.alpha.develop"]) and (app-version newer-than-or-equal-to "4.63") and (platform in ["Android"]))`)
		assert.NotNil(b, tree)
		assert.Equal(b, true, got)
		assert.NoError(b, err)
	}
}
