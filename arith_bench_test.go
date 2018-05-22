// +build !sse,!avx

package vecf32

import (
	"testing"

	"github.com/chewxy/math32"
)

/* BENCHMARKS */

func _vanillaVecAdd(a, b []float32) {
	for i := range a {
		a[i] += b[i]
	}
}

func BenchmarkVecAdd(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Add(x, y)
	}
}

func BenchmarkVanillaVecAdd(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecAdd(x, y)
	}
}

func _vanillaVecSub(a, b []float32) {
	for i := range a {
		a[i] -= b[i]
	}
}

func BenchmarkVecSub(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Sub(x, y)
	}
}

func BenchmarkVanillaVecSub(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecSub(x, y)
	}
}

func _vanillaVecMul(a, b []float32) {
	for i := range a {
		a[i] *= b[i]
	}
}

func BenchmarkVecMul(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Mul(x, y)
	}
}

func BenchmarkVanillaVecMul(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecMul(x, y)
	}
}

func _vanillaVecDiv(a, b []float32) {
	for i := range a {
		a[i] /= b[i]
	}
}

func BenchmarkVecDiv(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		Div(x, y)
	}
}

func BenchmarkVanillaVecDiv(b *testing.B) {
	x := Range(0, niceprime)
	y := Range(niceprime, 2*niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecDiv(x, y)
	}
}

func _vanillaVecSqrt(a []float32) {
	for i, v := range a {
		a[i] = math32.MobileSqrt(v)
	}
}

func BenchmarkVecSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		Sqrt(x)
	}
}

func BenchmarkVanillaVecSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecSqrt(x)
	}
}

func _vanillaVecInverseSqrt(a []float32) {
	for i, v := range a {
		a[i] = 1.0 / math32.MobileSqrt(v)
	}
}

func BenchmarkVecInvSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		InvSqrt(x)
	}
}

func BenchmarkVanillaVecInvSqrt(b *testing.B) {
	x := Range(0, niceprime)

	for n := 0; n < b.N; n++ {
		_vanillaVecInverseSqrt(x)
	}
}
