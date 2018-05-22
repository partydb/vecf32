// +build !sse
// +build !avx

package vecf32

/*

IMPORTANT NOTE:

Currently Div does not handle division by zero correctly. It returns a NaN instead of +Inf

*/

import (
	"testing"
	"unsafe"

	"github.com/chewxy/math32"
	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	assert := assert.New(t)

	a := Range(0, niceprime-1)

	correct := Range(0, niceprime-1)
	for i := range correct {
		correct[i] = 1
	}
	Div(a, a)
	assert.Equal(correct[1:], a[1:])
	assert.Equal(true, math32.IsInf(a[0], 0), "a[0] is: %v", a[0])

	b := Range(niceprime, 2*niceprime-1)
	for i := range correct {
		correct[i] = a[i] / b[i]
	}

	Div(a, b)
	assert.Equal(correct[1:], a[1:])
	assert.Equal(true, math32.IsInf(a[0], 0), "a[0] is: %v", a[0])

	/* Weird Corner Cases*/

	for i := 1; i < 65; i++ {
		a = Range(0, i)
		var testAlign bool
		addr := &a[0]
		u := uint(uintptr(unsafe.Pointer(addr)))
		if u&uint(32) != 0 {
			testAlign = true
		}

		if testAlign {
			b = Range(i, 2*i)
			correct = make([]float32, i)
			for j := range correct {
				correct[j] = a[j] / b[j]
			}
			Div(a, b)
			assert.Equal(correct[1:], a[1:])
		}
	}

}

func TestSqrt(t *testing.T) {
	assert := assert.New(t)

	a := Range(0, niceprime-1)

	correct := Range(0, niceprime-1)
	for i, v := range correct {
		correct[i] = math32.MobileSqrt(v)
	}
	Sqrt(a)
	assert.Equal(correct, a)

	// negatives
	a = []float32{-1, -2, -3, -4}
	Sqrt(a)

	for _, v := range a {
		if !math32.IsNaN(v) {
			t.Error("Expected NaN")
		}
	}

	/* Weird Corner Cases*/
	for i := 1; i < 65; i++ {
		a = Range(0, i)
		var testAlign bool
		addr := &a[0]
		u := uint(uintptr(unsafe.Pointer(addr)))
		if u&uint(32) != 0 {
			testAlign = true
		}

		if testAlign {
			correct = make([]float32, i)
			for j := range correct {
				correct[j] = math32.MobileSqrt(a[j])
			}
			Sqrt(a)
			assert.Equal(correct, a)
		}
	}
}

func TestInvSqrt(t *testing.T) {

	assert := assert.New(t)
	a := Range(0, niceprime-1)

	correct := Range(0, niceprime-1)
	for i, v := range correct {
		correct[i] = 1.0 / math32.MobileSqrt(v)
	}
	InvSqrt(a)
	assert.Equal(correct[1:], a[1:])
	if !math32.IsInf(a[0], 0) {
		t.Error("1/0 should be +Inf or -Inf")
	}

	// Weird Corner Cases

	for i := 1; i < 65; i++ {
		a = Range(0, i)
		var testAlign bool
		addr := &a[0]
		u := uint(uintptr(unsafe.Pointer(addr)))
		if u&uint(32) != 0 {
			testAlign = true
		}

		if testAlign {
			correct = make([]float32, i)
			for j := range correct {
				correct[j] = 1.0 / math32.MobileSqrt(a[j])
			}
			InvSqrt(a)
			assert.Equal(correct[1:], a[1:], "i = %d, %v", i, Range(0, i))
			if !math32.IsInf(a[0], 0) {
				t.Error("1/0 should be +Inf or -Inf")
			}
		}
	}
}
