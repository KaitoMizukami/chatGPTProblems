package passwordValidation

import (
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "Correct Password",
			password: "Thisiscorrect1@",
			want:     true,
		},
		{
			name:     "Password is less than 8",
			password: "Shor1*",
			want:     false,
		},
		{
			name:     "No uppercase in passsword",
			password: "nouppercase1%",
			want:     false,
		},
		{
			name:     "No lowercase in password",
			password: "NOLOWERCSE1&",
			want:     false,
		},
		{
			name:     "No special character in password",
			password: "NoSpecialChar",
			want:     false,
		},
	}

	for _, tc := range tests {
		got := validatePassword(tc.password)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
