package wps_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"gitlab.com/fffd/wps"
)

func TestUserAgent(t *testing.T) {
	want := "testid"
	got := wps.NewUserAgent(want)
	equals(t, got.ID, want)
}

func TestAddOne(t *testing.T) {
	table := []struct {
		in, want int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}
	for _, test := range table {
		got := wps.AddOne(test.in)
		equals(t, got, test.want)
	}
}

// equals fails the test if got is not equal to want.
func equals(tb testing.TB, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\tgot: %#v\n\n\twant: %#v\033[39m\n\n", filepath.Base(file), line, got, want)
		tb.FailNow()
	}
}
