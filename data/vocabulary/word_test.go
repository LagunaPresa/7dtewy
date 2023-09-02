package vocabulary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/LagunaPresa/7dtewy/codec"
)

func TestDict(t *testing.T) {
	t.Run("no duplication", func(t *testing.T) {
		used := make(map[string]interface{})
		for _, tt := range Dict {
			_, found := used[tt.Plain]
			assert.False(t, found, "`%s` is duplicated", tt.Plain)
			used[tt.Plain] = nil
		}
	})

	for _, tt := range Dict {
		t.Run(fmt.Sprintf("decodable %s", tt.Plain), func(t *testing.T) {
			got, _ := codec.DecodeCandidates(tt.Encoded)
			assert.Contains(t, got, tt.Plain)
		})
	}

	for _, tt := range Dict {
		t.Run(fmt.Sprintf("encodable %s", tt.Plain), func(t *testing.T) {
			got, _ := codec.EncodeCandidates(tt.Plain)
			assert.Contains(t, got, tt.Encoded)
		})
	}
}
