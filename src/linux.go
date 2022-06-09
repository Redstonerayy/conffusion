package src

func Linux(verbose bool) {
	PkgManager := GetPackageManager(verbose)
	SysPackages := GetPackages(verbose, PkgManager)
	DirPath := CreateSyncFolder(false)
}
