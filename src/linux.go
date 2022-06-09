package src

import (
	"path"
	"strings"
)

func Linux(verbose bool) {
	PkgManager := GetPackageManager(verbose)
	SysPackages := GetPackages(verbose, PkgManager)
	DirPath := CreateSyncFolder(verbose)
	_ = WriteFile(true, path.Join(DirPath, "pkgs.txt"), strings.Join(SysPackages, "\n"))
}
