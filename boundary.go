package carbon

// StartOfCentury returns a Carbon instance for start of the century.
// 本世纪开始时间
func (c *Carbon) StartOfCentury() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(c.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0).SetLocation(c.loc)
}

// EndOfCentury returns a Carbon instance for end of the century.
// 本世纪结束时间
func (c *Carbon) EndOfCentury() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(c.Year()/YearsPerCentury*YearsPerCentury+99, 12, 31, 23, 59, 59, 999999999).SetLocation(c.loc)
}

// StartOfDecade returns a Carbon instance for start of the decade.
// 本年代开始时间
func (c *Carbon) StartOfDecade() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(c.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0).SetLocation(c.loc)
}

// EndOfDecade returns a Carbon instance for end of the decade.
// 本年代结束时间
func (c *Carbon) EndOfDecade() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(c.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999).SetLocation(c.loc)
}

// StartOfYear returns a Carbon instance for start of the year.
// 本年开始时间
func (c *Carbon) StartOfYear() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(c.Year(), 1, 1, 0, 0, 0, 0).SetLocation(c.loc)
}

// EndOfYear returns a Carbon instance for end of the year.
// 本年结束时间
func (c *Carbon) EndOfYear() *Carbon {
	if c.IsInvalid() {
		return c
	}
	return create(c.Year(), 12, 31, 23, 59, 59, 999999999).SetLocation(c.loc)
}

// StartOfQuarter returns a Carbon instance for start of the quarter.
// 本季度开始时间
func (c *Carbon) StartOfQuarter() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter, day := c.Year(), c.Quarter(), 1
	return create(year, 3*quarter-2, day, 0, 0, 0, 0).SetLocation(c.loc)
}

// EndOfQuarter returns a Carbon instance for end of the quarter.
// 本季度结束时间
func (c *Carbon) EndOfQuarter() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, quarter, day := c.Year(), c.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	return create(year, 3*quarter, day, 23, 59, 59, 999999999).SetLocation(c.loc)
}

// StartOfMonth returns a Carbon instance for start of the month.
// 本月开始时间
func (c *Carbon) StartOfMonth() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return create(year, month, 1, 0, 0, 0, 0).SetLocation(c.loc)
}

// EndOfMonth returns a Carbon instance for end of the month.
// 本月结束时间
func (c *Carbon) EndOfMonth() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, _ := c.Date()
	return create(year, month+1, 0, 23, 59, 59, 999999999).SetLocation(c.loc)
}

// StartOfWeek returns a Carbon instance for start of the week.
// 本周开始时间
func (c *Carbon) StartOfWeek() *Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekStartsAt := c.StdTime().Weekday(), c.WeekStartsAt()
	if dayOfWeek == weekStartsAt {
		return c.StartOfDay()
	}
	return c.SubDays(int(DaysPerWeek+dayOfWeek-weekStartsAt) % DaysPerWeek).StartOfDay()
}

// EndOfWeek returns a Carbon instance for end of the week.
// 本周结束时间
func (c *Carbon) EndOfWeek() *Carbon {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekEndsAt := c.StdTime().Weekday(), c.WeekEndsAt()
	if dayOfWeek == weekEndsAt {
		return c.EndOfDay()
	}
	return c.AddDays(int(DaysPerWeek-dayOfWeek+weekEndsAt) % DaysPerWeek).EndOfDay()
}

// StartOfDay returns a Carbon instance for start of the day.
// 本日开始时间
func (c *Carbon) StartOfDay() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, 0, 0, 0, 0).SetLocation(c.loc)
}

// EndOfDay returns a Carbon instance for end of the day.
// 本日结束时间
func (c *Carbon) EndOfDay() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, 23, 59, 59, 999999999).SetLocation(c.loc)
}

// StartOfHour returns a Carbon instance for start of the hour.
// 小时开始时间
func (c *Carbon) StartOfHour() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, c.Hour(), 0, 0, 0).SetLocation(c.loc)
}

// EndOfHour returns a Carbon instance for end of the hour.
// 小时结束时间
func (c *Carbon) EndOfHour() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day := c.Date()
	return create(year, month, day, c.Hour(), 59, 59, 999999999).SetLocation(c.loc)
}

// StartOfMinute returns a Carbon instance for start of the minute.
// 分钟开始时间
func (c *Carbon) StartOfMinute() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return create(year, month, day, hour, minute, 0, 0).SetLocation(c.loc)
}

// EndOfMinute returns a Carbon instance for end of the minute.
// 分钟结束时间
func (c *Carbon) EndOfMinute() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, _ := c.DateTime()
	return create(year, month, day, hour, minute, 59, 999999999).SetLocation(c.loc)
}

// StartOfSecond returns a Carbon instance for start of the second.
// 秒开始时间
func (c *Carbon) StartOfSecond() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, 0).SetLocation(c.loc)
}

// EndOfSecond returns a Carbon instance for end of the second.
// 秒结束时间
func (c *Carbon) EndOfSecond() *Carbon {
	if c.IsInvalid() {
		return c
	}
	year, month, day, hour, minute, second := c.DateTime()
	return create(year, month, day, hour, minute, second, 999999999).SetLocation(c.loc)
}
