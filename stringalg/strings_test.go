package stringalg

import "testing"

func TestStringContainsHash(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s:      "ABCD",
				substr: "BAD",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s:      "ABCD",
				substr: "BCE",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s:      "ABCD",
				substr: "AA",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringContainsHash(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("StringContainsHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
