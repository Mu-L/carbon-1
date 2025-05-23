package carbon

import (
	"testing"
)

func BenchmarkZeroValue(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ZeroValue()
	}
}

func BenchmarkEpochValue(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		EpochValue()
	}
}

func BenchmarkMaxValue(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		MaxValue()
	}
}

func BenchmarkMinValue(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		MinValue()
	}
}

func BenchmarkMaxDuration(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		MaxDuration()
	}
}

func BenchmarkMinDuration(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		MinDuration()
	}
}

func BenchmarkMax(b *testing.B) {
	c1 := Parse("2020-08-06")
	c2 := Parse("2021-08-05")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Max(c1, c2)
	}
}

func BenchmarkMin(b *testing.B) {
	c1 := Parse("2020-08-06")
	c2 := Parse("2021-08-05")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Min(c1, c2)
	}
}

func BenchmarkCarbon_Closest(b *testing.B) {
	c1 := Parse("2020-08-04")
	c2 := Parse("2020-08-05")
	c3 := Parse("2021-08-06")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c1.Closest(c2, c3)
	}
}

func BenchmarkCarbon_Farthest(b *testing.B) {
	c1 := Parse("2020-08-04")
	c2 := Parse("2020-08-05")
	c3 := Parse("2021-08-06")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c1.Farthest(c2, c3)
	}
}
