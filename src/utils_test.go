package main

import (
	"testing"
	"strconv"
)

func TestIsDateOfBirthValid(t *testing.T) {
	cases := []struct { dob string; expected bool} {
        {"2000-01-02", true },
		{"1939-12-23", true },
		{"2099-12-23", false },
		{"20000102", false },
		{"2000-01-22-", false },
		{"22-01-02", false },
		{"yyyy-mm-dd", false },
    }
 
    for _, c := range cases {
        res := isDateOfBirthValid(c.dob)
        if res != c.expected {
            t.Log("error for ("+ c.dob +")  expecting:"+ strconv.FormatBool(c.expected) +", but got:"+ strconv.FormatBool(res))
            t.Fail()
        }
    }
}

func TestIsUsermameValid(t *testing.T) {
	cases := []struct { useername string; expected bool} {
        {"abc", true },
		{"AbC", true },
		{"123", false },
		{"i2E", false },
		{"_1?", false },
    }
 
    for _, c := range cases {
        res := isUsermameValid(c.useername)
        if res != c.expected {
            t.Log("error for ("+ c.useername +") expecting:"+ strconv.FormatBool(c.expected) +", but got:"+ strconv.FormatBool(res))
            t.Fail()
        }
    }
} 

func TestDaysBeforeBirthDay(t *testing.T){ 
	cases := []struct { dob string } {
        {"2022-01-01"},
		{"2000-12-31"},
		{"1911-11-11"},
    }
 
    for _, c := range cases {
        var res int = daysBeforeBirthDay(c.dob)
		if !(res < 366) {
            t.Log("error for ("+ c.dob +")  expecting < 366, but got:"+ strconv.Itoa(res))
            t.Fail()
        }
	}
}