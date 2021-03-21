package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ryoh827/sorting-file-script/pkg/file"
)

func Init() {
	fmt.Print("Input Dir: ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	rootDir := stdin.Text()

	paths := file.Dirwalk(rootDir)
	extList := file.GetExtList(paths)

	file.Sort(paths, rootDir, extList)
}
