package constants

import (
	"os"
)

var SubDirInsideRepo = ".layers"
var ChainSubDir = ".chains"
var DirtyChainSubDir = ".dirty-chains"

var DirPermission = os.FileMode(0755)
var FilePermission = os.FileMode(0644)
