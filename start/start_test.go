package start

import "testing"

func TestStar(t *testing.T) {
	cases := []struct {
		want string
	}{
		{"Hello, Shig-Neru"},
	}
	for _, c := range cases {
		got := HelloWorld()
		if got != c.want {
			t.Errorf("%s, is not %s", got, c.want)
		}
	}

}
