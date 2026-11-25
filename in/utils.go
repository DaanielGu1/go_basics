package in

import (
	"fmt"
	"errors"
	"strconv"
	"io"
)

func IOGetString(msg string, r io.Reader, w io.Writer) (string, error){
	_, err := fmt.Fprint(w, msg)
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
	s, err := IOGetString(msg, r, w)
	if err != nil {
		return 999, err
	}
	n, err := strconv.Atoi(s)
	if errors.Is(err, strconv.ErrRange) {
		fmt.Fprint(w, "Number out of range, type a lower number")
		return 9999, err
	}
	if errors.Is(err, strconv.ErrSyntax) {
		fmt.Fprint(w, "This string can't be parsed into a integer")
		return 99999, err
	}
	
	return n, err
}


