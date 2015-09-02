package comm
import (
	"os"
	"strings"
	"io/ioutil"
)

func ToString(filepath string) (string,error) {
	by, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(by),nil
}

func ToTrimString(filepath string) (string,error) {
	str,err := ToString(filepath)
	if err != nil {
		return "",err
	}
	return strings.TrimSpace(str),nil
}

func IsFile(fp string) bool {
	f,err := os.Stat(fp)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}