package main

import (
	"strings"
	"time"
)

const (
	FormatLongMonth             = "January"
	FormatMonth                 = "Jan"
	FormatNumMonth              = "1"
	FormatZeroMonth             = "01"
	FormatLongWeekDay           = "Monday"
	FormatWeekDay               = "Mon"
	FormatDay                   = "2"
	FormatUnderDay              = "_2"
	FormatZeroDay               = "02"
	FormatHour                  = "15"
	FormatHour12                = "3"
	FormatZeroHour12            = "03"
	FormatMinute                = "4"
	FormatZeroMinute            = "04"
	FormatSecond                = "5"
	FormatZeroSecond            = "05"
	FormatLongYear              = "2006"
	FormatYear                  = "06"
	FormatPM                    = "PM"
	Formatpm                    = "pm"
	FormatTZ                    = "MST"
	FormatISO8601TZ             = "Z0700" // prints Z for UTC
	FormatISO8601SecondsTZ      = "Z070000"
	FormatISO8601ColonTZ        = "Z07:00" // prints Z for UTC
	FormatISO8601ColonSecondsTZ = "Z07:00:00"
	FormatNumTZ                 = "-0700" // always numeric
	FormatNumSecondsTz          = "-070000"
	FormatNumShortTZ            = "-07"    // always numeric
	FormatNumColonTZ            = "-07:00" // always numeric
	FormatNumColonSecondsTZ     = "-07:00:00"
)

var PhpFormatMap = map[string]string{
	"d": FormatZeroDay,
	"D": FormatWeekDay,
	"j": FormatDay,
	"l": FormatLongWeekDay,
	"F": FormatLongMonth,
	"m": FormatZeroMonth,
	"M": FormatMonth,
	"n": FormatNumMonth,
	"Y": FormatLongYear,
	"y": FormatYear,
	"a": Formatpm,
	"A": FormatPM,
	"g": FormatHour12,
	"G": FormatHour,
	"h": FormatZeroHour12,
	"i": FormatZeroMinute,
	"s": FormatZeroSecond,
	"O": FormatNumTZ,
	"P": FormatNumColonTZ,
	"T": FormatTZ,
	"c": time.RFC3339,
	"r": time.RFC1123Z,
}

func ConvertFromPhp(s string) string {
	for phpString, goStrig := range PhpFormatMap {
		s = strings.Replace(s, phpString, goStrig, -1)
	}
	return s
}
