package carbon_test

import (
	"fmt"

	"github.com/dromara/carbon/v2"
)

func ExampleZeroValue() {
	fmt.Println(carbon.ZeroValue().ToString())
	// Output:
	// 0001-01-01 00:00:00 +0000 UTC
}

func ExampleEpochValue() {
	fmt.Println(carbon.EpochValue().ToString())
	// Output:
	// 1970-01-01 00:00:00 +0000 UTC
}

func ExampleMaxValue() {
	fmt.Println(carbon.MaxValue().ToString())
	// Output:
	// 9999-12-31 23:59:59.999999999 +0000 UTC
}

func ExampleMinValue() {
	fmt.Println(carbon.MinValue().ToString())
	// Output:
	// 0001-01-01 00:00:00 +0000 UTC
}

func ExampleMaxDuration() {
	fmt.Println(carbon.MaxDuration().Seconds())
	// Output:
	// 9.223372036854776e+09
}

func ExampleMinDuration() {
	fmt.Println(carbon.MinDuration().Seconds())
	// Output:
	// -9.223372036854776e+09
}

func ExampleMax() {
	fmt.Println(carbon.Max(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2021-08-05 00:00:00 +0000 UTC
}

func ExampleMin() {
	fmt.Println(carbon.Min(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2020-08-06 00:00:00 +0000 UTC
}

func ExampleCarbon_Closest() {
	fmt.Println(carbon.Parse("2020-08-05").Closest(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2020-08-06 00:00:00 +0000 UTC
}

func ExampleCarbon_Farthest() {
	fmt.Println(carbon.Parse("2020-08-05").Farthest(carbon.Parse("2020-08-06"), carbon.Parse("2021-08-05")).ToString())
	// Output:
	// 2021-08-05 00:00:00 +0000 UTC
}
