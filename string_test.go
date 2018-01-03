package string

import "testing"

func Test(t *testing.T) {
      var tests = []struct {
        s, want string
      }{
        {"Backwards", "sdrawkcaB"},
        {"Hello, Radar", "radaR ,olleH"},
        {"Hello, ステーキ", "キーテス ,olleH"},
        {"", ""},
      }
      for _, c := range tests {
          got := Reverse(c.s)
          if got != c.want {
            t.Error("Reverse (%q) == %q, want %q", c.s, got, c.want)
          }
      }
}
