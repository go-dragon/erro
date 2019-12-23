package erro

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	  err := NewError("hello")
	  fmt.Println(err)
}

