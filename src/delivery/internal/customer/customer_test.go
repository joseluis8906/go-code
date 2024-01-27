package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	testCases := map[string]struct {
		input   string
		want    Email
		wantErr bool
	}{
		"Valid": {
			input: "john.doe@example.com",
			want:  Email{value: "john.doe@example.com"},
		},

		"MissingDomain": {
			input:   "john.doe",
			want:    Email{value: "<nil>"},
			wantErr: true,
		},

		"MissingName": {
			input:   "@example.com",
			want:    Email{value: "<nil>"},
			wantErr: true,
		},

		"NotAllowedChars": {
			input:   "john#doe@example.com",
			want:    Email{value: "<nil>"},
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewEmail(tc.input)

			assert.Equal(t, tc.want, got)
			if tc.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
