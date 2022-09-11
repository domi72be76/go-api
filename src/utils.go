package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"regexp"
)

// Validate the date of birth value is 
// on format YYYY-MM-DD
// and before the today date
func isDateOfBirthValid(dateOfBirth string) bool 	{

	obj, err := regexp.MatchString("^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$", dateOfBirth)
	if err != nil {
		fmt.Println("Unable to valifate dateOfBirth.")
	}
	
	if obj {
		// check if birth Date is before today
		var dateIntArray = getIntSplitDate(dateOfBirth)
		today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
		timeOfBirth := time.Date(dateIntArray[0], time.Month(dateIntArray[1]), dateIntArray[2], 0, 0, 0, 0, time.UTC) // Date(2017, 1, 1
		
		return  timeOfBirth.Before(today)
	} else {
		return false
	}
}

//Validate the username value contain only letters
func isUsermameValid(user string) bool 		{
	obj, err := regexp.MatchString("^[a-zA-Z]+$", user)
	if err != nil {
		fmt.Println("Unable to valifate username.")
	}
	
	return obj
}

//Validate the username value contain only letters
func daysBeforeBirthDay(dateOfBirth string) int { 
	
	var dateIntArray = getIntSplitDate(dateOfBirth)
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	// if birthday already append this yea
	nextBirthDay := time.Date(time.Now().Year() , time.Month(dateIntArray[1]), dateIntArray[2], 0, 0, 0, 0, time.UTC) 
	if nextBirthDay.Before(today)  {
		// then it will be next year
		nextBirthDay = time.Date(time.Now().Year()+1 , time.Month(dateIntArray[1]), dateIntArray[2], 0, 0, 0, 0, time.UTC) 			
	}

	// return the number of day
	return int(nextBirthDay.Sub(today).Hours() / 24)
}

func getIntSplitDate(dateOfBirth string) []int { 
	dateStrArray := strings.Split(dateOfBirth, "-")
	
	//cast date string array to int array
	var dateIntArray = []int{}
	for _, i := range dateStrArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		dateIntArray = append(dateIntArray, j)
	}

	return dateIntArray
}