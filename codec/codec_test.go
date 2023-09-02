package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCandidates(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		want    []string
		wantErr bool
	}{
		{
			name:    "usual",
			encoded: "errn",
			want:    []string{"book", "ymmi"},
		},
		{
			name:    "outside the range",
			encoded: "ERRN",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeCandidates(tt.encoded)
			if assert.Equal(t, tt.wantErr, err != nil) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEncodeCandidates(t *testing.T) {
	tests := []struct {
		name    string
		plain   string
		want    []string
		wantErr bool
	}{
		{
			name:  "usual",
			plain: "stagnant",
			want:  []string{"abhnuhub", "zbhnuhub"},
		},
		{
			name:    "outside the range",
			plain:   "zoo",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeCandidates(tt.plain)
			if assert.Equal(t, tt.wantErr, err != nil) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
