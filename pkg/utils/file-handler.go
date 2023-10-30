package utils
import (
	"fmt"
	"os"
)

func ReadFile(Path string) (string, error) {
	file, err := os.ReadFile(Path)
	return string(file), err
}

func SaveFile(destion_path string, data string) bool {
	file, err := os.Create(destion_path)
	if err != nil {
		fmt.Println("Can't create the file", err)
		return false
	}
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Can't write to the file", err)
		return false
	}
	defer file.Close()
	return true

}
