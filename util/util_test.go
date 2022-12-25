package util

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestGenerateId(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	fmt.Println(r)
	r = r * 10000.0
	fmt.Println(r)
	r = math.Round(r)
	fmt.Println(r)
	yy := int64(r * 10000)
	fmt.Println(yy)
}
