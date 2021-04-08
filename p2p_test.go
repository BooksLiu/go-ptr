package ptr

import (
	"testing"
)

func TestPointerInt64ToTime(t *testing.T) {
	tests := []struct {
		name string
		args *int64
		want [2]string
	}{
		{
			name: "to-time-1617879700",
			args: Int64(1617879700),
			want: [2]string{"2021-04-08T19:01:40", "2021-04-08T11:01:40"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointerInt64ToTime(tt.args); got.Format("2006-01-02T15:04:05") != tt.want[0] && got.Format("2006-01-02T15:04:05") != tt.want[1] {
				t.Errorf("PointerInt64ToTime() = %v, want %v", got.Format("2006-01-02T15:04:05"), tt.want)
			}
		})
	}
}
