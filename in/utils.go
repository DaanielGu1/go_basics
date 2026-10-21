package in

import (
	"fmt"
	"errors"
	"strconv"
)

func IOGetString(msg string) (string,error){
	_,err:= fmt.Println(msg)
	if err!= nil {
		return "",err
	}

	var str string

	_,err=fmt.Scan(&str)
	if err!= nil {
		return "",err
	}

	return str,err
}

func IOGetInt(msg string) (int,error) {
	integerStr,err := IOGetString(msg)
	if err != nil {
		return 0,err
	}
	for true{
		integer,err:= strconv.Atoi(integerStr)
		if errors.Is(err, strconv.ErrRange){
			fmt.Println("Index out of range, type a lower id number")
			continue
		}
		if errors.Is(err, strconv.ErrSyntax){
			fmt.Println("This string can't be parsed into a integer")
			continue
		}
		return integer,err
		}
	return 0,err
}

