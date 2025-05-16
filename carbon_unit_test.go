package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type CarbonSuite struct {
	suite.Suite
}

func TestCarbonSuite(t *testing.T) {
	suite.Run(t, new(CarbonSuite))
}

func (s *CarbonSuite) TestNewCarbon() {
	loc, _ := time.LoadLocation(PRC)

	t1, _ := time.Parse(DateTimeLayout, "2020-08-05 13:14:15")
	t2, _ := time.ParseInLocation(DateTimeLayout, "2020-08-05 13:14:15", loc)

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.False(c.HasError())
		s.True(c.IsZero())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15 +0000 UTC", NewCarbon(t1).ToString())
		s.Equal(t1.String(), NewCarbon(t1).ToString())

		s.Equal("2020-08-05 13:14:15 +0800 CST", NewCarbon(t2).ToString())
		s.Equal(t2.String(), NewCarbon(t2).ToString())
	})
}

func (s *CarbonSuite) TestCarbon_Copy() {
	s.Run("copy nil", func() {
		var oldCarbon *Carbon
		oldCarbon = nil
		newCarbon := oldCarbon.Copy()

		s.Nil(oldCarbon)
		s.Nil(newCarbon)

		oldCarbon = oldCarbon.AddDay()
		s.Nil(oldCarbon)
		s.Nil(newCarbon)
	})

	s.Run("copy time", func() {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		s.Equal("2020-08-05 00:00:00 +0000 UTC", oldCarbon.ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", newCarbon.ToString())

		oldCarbon = oldCarbon.AddDay()
		s.Equal("2020-08-06 00:00:00 +0000 UTC", oldCarbon.ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", newCarbon.ToString())
	})

	s.Run("copy timezone", func() {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		s.Equal(UTC, oldCarbon.Timezone())
		s.Equal(UTC, newCarbon.Timezone())

		oldCarbon = oldCarbon.SetTimezone(PRC)
		s.Equal(PRC, oldCarbon.Timezone())
		s.Equal(UTC, newCarbon.Timezone())

		newCarbon = newCarbon.SetTimezone(Japan)
		s.Equal(PRC, oldCarbon.Timezone())
		s.Equal(Japan, newCarbon.Timezone())
	})

	s.Run("copy layout", func() {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		s.Equal("2006-01-02", oldCarbon.CurrentLayout())
		s.Equal("2006-01-02", newCarbon.CurrentLayout())

		oldCarbon = oldCarbon.SetLayout(DateTimeLayout)
		s.Equal(DateTimeLayout, oldCarbon.CurrentLayout())
		s.Equal(DateLayout, newCarbon.CurrentLayout())

		newCarbon = newCarbon.SetLayout(RFC1036Layout)
		s.Equal(DateTimeLayout, oldCarbon.CurrentLayout())
		s.Equal(RFC1036Layout, newCarbon.CurrentLayout())
	})

	s.Run("copy weekStartsAt", func() {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		s.Equal(DefaultWeekStartsAt, oldCarbon.WeekStartsAt())
		s.Equal(DefaultWeekStartsAt, newCarbon.WeekStartsAt())

		oldCarbon = oldCarbon.SetWeekStartsAt(Monday)
		s.Equal(Monday, oldCarbon.WeekStartsAt())
		s.Equal(DefaultWeekStartsAt, newCarbon.WeekStartsAt())

		newCarbon = newCarbon.SetWeekStartsAt(Sunday)
		s.Equal(Monday, oldCarbon.WeekStartsAt())
		s.Equal(Sunday, newCarbon.WeekStartsAt())
	})

	s.Run("copy lang", func() {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		s.Equal("August", oldCarbon.ToMonthString())
		s.Equal("August", newCarbon.ToMonthString())

		oldCarbon.SetLocale("zh-CN")
		s.Equal("八月", oldCarbon.ToMonthString())
		s.Equal("August", newCarbon.ToMonthString())

		newCarbon.SetLocale("jp")
		s.Equal("八月", oldCarbon.ToMonthString())
		s.Equal("8月", newCarbon.ToMonthString())
	})

	s.Run("copy error", func() {
		oldCarbon := Parse("xxx")
		newCarbon := oldCarbon.Copy()

		s.True(oldCarbon.HasError())
		s.True(newCarbon.HasError())

		newCarbon = newCarbon.SetLayout("xxx")
		s.True(oldCarbon.HasError())
		s.True(newCarbon.HasError())
	})
}
