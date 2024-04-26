package readfilehelper

import (
	//"errors"
	"fmt"
	"testing"
)

var test1 = []struct {
	arg1   string
	result string
}{
	{
		"testfile", "https://google.com",
	},
	{
		"testfile", "https://www.microsoft.com",
	},
}

var test2 = []struct {
	arg1   string
	result string
}{
	{
		"testfile", "nil",
	},
	{
		"testfile2", "error opening Url's file",
	},
}

//error opening Url's file

func TestReadfilehelper(t *testing.T) {
	for _, t := range test1 {
		fmt.Println(t)

	}
	for x, st := range test1 {
		if s, _ := ReadLineOfFile(st.arg1); s[x] != st.result {
			t.Errorf("Wanted %v, got %v \n", st.result, s)

		}

	}
	for _, sv := range test2 {
		if _, err := ReadLineOfFile(sv.arg1); err != sv.result {
			t.Errorf("Wanted %v, got %v \n", sv.result, err)

		}
	}

}
