package x
import (
    "os/exec"
    "bytes"
)



func RunCmd(keys string) (string, bool,error) {
    cmd := exec.Command("sh","-c",keys)
    var out bytes.Buffer
	  cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
		return "Error in Run()",false,err
	}
	return out.String(),true, err
}
