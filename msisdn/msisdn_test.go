package msisdn

import "testing"

func TestNormalizeMsisdn(t *testing.T) {
	tests := []struct {
		name   string
		msisdn string
		want   string
	}{
		{"start_with_zero", "08123456789", "628123456789"},
		{"start_with_eighth", "8123456789", "628123456789"},
		{"start_with_plus", "+628123456789", "628123456789"},
		{"contain_space", " +628123456789 ", "628123456789"},
		{"contain_separator", " +6281-234-56789 ", "628123456789"},
		{"separate_by_space", " 6281 234 56789 ", "628123456789"},
		{"invalid", "8", ""},
		{"invalid", "0", ""},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeMsisdn(tt.msisdn); got != tt.want {
				t.Errorf("NormalizeMsisdn() = %v, want %v", got, tt.want)
			}
		})
	}
}
