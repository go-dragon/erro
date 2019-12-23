package erro

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	type T struct {
		Name string `json:"name"`
	}
	err := NewError(T{Name: "hi"})
	fmt.Println(err)
}
