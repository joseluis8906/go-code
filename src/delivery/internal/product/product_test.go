package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	testCases := map[string]struct {
		input   string
		want    Name
		wantErr bool
	}{
		"Valid": {
			input: "Product",
			want:  Name{value: "Product"},
		},
		"TooShort": {
			input:   "Pr",
			want:    Name{value: "<nil>"},
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewName(tc.input)

			assert.Equal(t, tc.want, got)
			if tc.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestReference(t *testing.T) {
	testCases := map[string]struct {
		input   string
		want    Ref
		wantErr bool
	}{
		"Valid": {
			input: "ABC-123",
			want:  Ref{value: "ABC-123"},
		},
		"TooShortLetters": {
			input:   "AB-123",
			want:    Ref{value: "<nil>"},
			wantErr: true,
		},
		"TooShortNumbers": {
			input:   "ABC-12",
			want:    Ref{value: "<nil>"},
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewRef(tc.input)

			assert.Equal(t, tc.want, got)
			if tc.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
