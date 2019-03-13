package goutils

/*
//usage: install https://nuwen.net/mingw.html, and add C:\MinGW\bin to PATH
const char* build_date(void)
{
	static const char* psz_build_date = __DATE__;
	return psz_build_date;
}
const char* build_time(void)
{
	static const char* psz_build_time = __TIME__;
	return psz_build_time;
}
*/
import "C"

import (
	"strings"
)

var (
	_buildDate = C.GoString(C.build_date())
	_buildTime = C.GoString(C.build_time())
)


func GetCompileTime() string {
	date_array := strings.FieldsFunc(_buildDate, func(c rune) bool { return c == ' ' })
	time_array := strings.FieldsFunc(_buildTime, func(c rune) bool { return c == ':' })
	//date_array := strings.Split(_buildDate, " ")
	//time_array := strings.Split(_buildTime, ":")
	if(3 != len(date_array) || 3 != len(time_array)) {
		return "can't support get compile time"
	}

	var month string = ""
	switch date_array[0] {
	case "Jan":
		month = "1"
	case "Feb":
		month = "2"
	case "Mar":
		month = "3"
	case "Apr":
		month = "4"
	case "May":
		month = "5"
	case "Jun":
		month = "6"
	case "Jul":
		month = "7"
	case "Aug":
		month = "8"
	case "Sep":
		month = "9"
	case "Oct":
		month = "10"
	case "Nov":
		month = "11"
	case "Dec":
		month = "12"
	}

	result := date_array[2] + "-" + month + "-" + date_array[1] + " " + time_array[0] + ":" + time_array[1] + ":" + time_array[2]
	return result
}
