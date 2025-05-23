package carbon

import "time"

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

// MaxValue returns a Carbon instance for the greatest supported date.
// 返回 Carbon 的最大值
func MaxValue() *Carbon {
	return NewCarbon(time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC))
}

// MinValue returns a Carbon instance for the lowest supported date.
// 返回 Carbon 的最小值
func MinValue() *Carbon {
	return NewCarbon(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
}

// MaxDuration returns the maximum duration value.
// 返回 Duration 的最大值
func MaxDuration() Duration {
	return maxDuration
}

// MinDuration returns the minimum duration value.
// 返回 Duration 的最小值
func MinDuration() Duration {
	return minDuration
}

// Max returns the maximum Carbon instance from the given Carbon instance.
// 返回最大的 Carbon 实例
func Max(c1 *Carbon, c2 ...*Carbon) (c *Carbon) {
	c = c1
	if c.IsInvalid() {
		return
	}
	if len(c2) == 0 {
		return
	}
	for i := range c2 {
		if c2[i].IsInvalid() {
			return c2[i]
		}
		if c2[i].Gte(c) {
			c = c2[i]
		}
	}
	return
}

// Min returns the minimum Carbon instance from the given Carbon instance.
// 返回最小的 Carbon 实例
func Min(c1 *Carbon, c2 ...*Carbon) (c *Carbon) {
	c = c1
	if c.IsInvalid() {
		return
	}
	if len(c2) == 0 {
		return
	}
	for i := range c2 {
		if c2[i].IsInvalid() {
			return c2[i]
		}
		if c2[i].Lte(c) {
			c = c2[i]
		}
	}
	return
}

// Closest returns the closest Carbon instance from the given Carbon instance.
// 返回离给定 carbon 实例最近的 Carbon 实例
func (c *Carbon) Closest(c1 *Carbon, c2 *Carbon) *Carbon {
	if c.IsInvalid() {
		return c
	}
	if c1.IsInvalid() {
		return c1
	}
	if c2.IsInvalid() {
		return c2
	}
	if c.DiffAbsInSeconds(c1) < c.DiffAbsInSeconds(c2) {
		return c1
	}
	return c2
}

// Farthest returns the farthest Carbon instance from the given Carbon instance.
// 返回离给定 carbon 实例最远的 Carbon 实例
func (c *Carbon) Farthest(c1 *Carbon, c2 *Carbon) *Carbon {
	if c.IsInvalid() {
		return c
	}
	if c1.IsInvalid() {
		return c1
	}
	if c2.IsInvalid() {
		return c2
	}
	if c.DiffAbsInSeconds(c1) > c.DiffAbsInSeconds(c2) {
		return c1
	}
	return c2
}
