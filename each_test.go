package gotil_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestEach(t *testing.T) {
	input := []int{10, 20, 30}
	expected := "102030"
	result := ""
	gotil.Each(input, func(v int) {
		result += fmt.Sprint(v)
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Each does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}

}

func BenchmarkEach(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gotil.Each([]int{10, 20, 30}, func(v int) {
			fmt.Println(v)
		})
	}
}

func ExampleEach() {
	gotil.Each([]string{"gotilty", "gotil"}, func(v string) {
		fmt.Fprint(os.Stdout, v)
	})
	// Output: gotiltygotil
}
