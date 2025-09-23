package main
import (
	"fmt"
	"os"
	// "log"
	"path"
)
func main(){
	//TempDir:returns the default directory used on the current OS for temporary files,
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	fmt.Println(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Println(err)
		return
		}
	defer f.Close()
}
