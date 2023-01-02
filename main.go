package main

import "tfupdate/pkg/utils"

func main() {
	futils := utils.NewFolderUtils("./pkg")
	folders := futils.ListDir()

	for _, f := range folders {
		println(f)
	}
}
