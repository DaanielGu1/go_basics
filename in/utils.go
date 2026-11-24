package in

import (
	"fmt"
	"errors"
	"strconv"
	"io"
)

func IOGetString(msg string, r io.Reader, w io.Writer) (string, error){
	_, err := fmt.Println(msg)
	if err != nil {
		return "", err
	}

	var str string

	_, err = fmt.Fscan(r, &str)
	if err != nil {
		return "", err
	}

	return str,err
}

func IOGetInt(msg string, r io.Reader, w io.Writer) (int, error) {
	for {
		s, err := IOGetString(msg, r, w)
		if err != nil {
			return 0, err
		}
		n, err := strconv.Atoi(s)
		if errors.Is(err, strconv.ErrRange) {
			fmt.Println("Index out of range, type a lower id number")
			continue
		}
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Println("This string can't be parsed into a integer")
			continue
		}
		return n, err
	}
}

