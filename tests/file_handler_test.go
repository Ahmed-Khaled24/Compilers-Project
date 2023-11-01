package tests

import (
	"os"
	"testing"

	"github.com/ahmedelsayed968/Compilers-Project/pkg/utils"
)

func TestSaveFile(t *testing.T) {
	tempFileName := "testfile.txt"

	// Test case 1: Successful file creation and write
	data := "Test data"
	if !utils.SaveFile(tempFileName, data) {
		t.Errorf("SaveFile failed for a valid case")
	}

	// Test case 2: File not created
	invalidFileName := "/this/is/an/invalid/path/testfile.txt"
	if utils.SaveFile(invalidFileName, data) {
		t.Errorf("SaveFile succeeded for an invalid file path")
	}

	// Test case 3: Failed write
	readOnlyFileName := tempFileName
	if !utils.SaveFile(readOnlyFileName, data) {
		t.Errorf("SaveFile failed for a read-only file")
	}
}
func TestReadFile(t *testing.T) {

	// Test case 1:
	// Invalid Path
	PATH1 := "this/is/an/invalid/path/testfile.txt"
	_, err := utils.ReadFile(PATH1)
	if err == nil {
		t.Errorf("Read From invalid Path")
	}

	// Test case 2:
	// valid Path
	PATH2 := "/test.txt"
	_, err = os.Create(PATH2)
	_, err = utils.ReadFile(PATH2)
	if err != nil {
		t.Errorf("ReadFile filed from Reading a valid file")
	}
}
