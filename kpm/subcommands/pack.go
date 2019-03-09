package subcommands

import (
	"fmt"
	"os"

	"./common"
	"./utils/constants"
	"./utils/files"
	"./utils/log"
	"./utils/types"
)

// PackCmd packs a local template package so it is available for use in the given local KPM repository.
func PackCmd(packageDirPathArg *string, kpmHomeDirPathArg *string) error {
	var err error

	// Package directory
	var packageDirPath string
	packageDirPath, err = files.GetAbsolutePathOrDefaultFunc(packageDirPathArg, files.GetWorkingDir)
	if err != nil {
		return err
	}

	// Get KPM home directory
	var kpmHomeDirPath string
	kpmHomeDirPath, err = files.GetAbsolutePathOrDefaultFunc(kpmHomeDirPathArg, constants.GetDefaultKpmHomeDirPath)
	if err != nil {
		return err
	}

	// Get local package repository directory
	var localPackageRepositoryDirPath = constants.GetPackageRepositoryDirPath(kpmHomeDirPath)

	// Log resolved paths
	log.Info("====")
	log.Info(fmt.Sprintf("Package directory:             %s", packageDirPath))
	log.Info(fmt.Sprintf("Package repository directory:  %s", localPackageRepositoryDirPath))
	log.Info("====")

	// Validate package and get package info
	log.Verbose("Getting package info")
	var packageInfo *types.PackageInfo
	packageInfo, err = common.GetPackageInfo(packageDirPath)
	if err != nil {
		return err
	}

	// Get package name with version and output path
	var packageNameWithVersion = constants.GetPackageFullName(packageInfo.Name, packageInfo.Version)
	var outputDirPath = constants.GetPackageDirPath(localPackageRepositoryDirPath, packageNameWithVersion)

	// Delete the output directory in case it isn't empty
	os.RemoveAll(outputDirPath)

	// Copy package to output directory
	log.Verbose(fmt.Sprintf("Copying package to: %s", outputDirPath))
	files.CopyDir(packageDirPath, outputDirPath)

	log.Info(fmt.Sprintf("Template package name:    %s", packageInfo.Name))
	log.Info(fmt.Sprintf("Template package version: %s", packageInfo.Version))

	log.Verbose(fmt.Sprintf("Repository directory:     %s", outputDirPath))

	return nil
}
