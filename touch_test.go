package main

import (
	"os"
	"testing"
	"time"
)

func TestTouch(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs } ()
	testFileName := "test.txt"
	os.Args = []string{"cmd", testFileName}
	touch()
	fileInfo, err := os.Stat(testFileName)
	if err != nil {
		t.Errorf("createFile(\"test.txt\"): Unexpected error: %v", err)
	}
	actualName := fileInfo.Name()
	if actualName != testFileName {
		t.Errorf("createFile(\"test.txt\"): file name expected %s got %s", testFileName, actualName)
	}
	touch()
	fileInfo, err = os.Stat(testFileName)
	now := time.Now().Local().UnixNano() / int64(time.Second)
	actualTime := fileInfo.ModTime().UnixNano() / int64(time.Second)
	if now != actualTime {
		t.Errorf("changeFileTime(\"test.txt\"): mod time expected %v got %v", now, actualTime)
	}

}
