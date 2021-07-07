package models

import (
	"errors"

	"gorm.io/gorm"
)

// Image is what generates a OSTree Commit.
type Image struct {
	gorm.Model
	Name         string
	Account      string
	Distribution string
	Description  string
	Status       string
	Version      int `gorm:"default:1"`
	CommitID     int
	Commit       *Commit
	InstallerID  int
	Installer    *Installer
}

const (
	// Errors
	// DistributionCantBeNilMessage is the error message when a distribution is nil
	DistributionCantBeNilMessage = "distribution can't be empty"
	// ArchitectureCantBeEmptyMessage is the error message when the architecture is empty
	ArchitectureCantBeEmptyMessage = "architecture can't be empty"

	// ImageTypes
	// ImageTypeInstaller is the installer image type on Image Builder
	ImageTypeInstaller = "rhel-edge-installer"
	// ImageTypeCommit is the installer image type on Image Builder
	ImageTypeCommit = "rhel-edge-commit"

	// Status
	// ImageStatusCreated is for when a image is created
	ImageStatusCreated = "CREATED"
	// ImageStatusBuilding is for when a image is building
	ImageStatusBuilding = "BUILDING"
	// ImageStatusError is for when a image is on a error state
	ImageStatusError = "ERROR"
	// ImageStatusSuccess is for when a image is available to the user
	ImageStatusSuccess = "SUCCESS"
)

// ValidateRequest validates an Image Request
func (i *Image) ValidateRequest() error {
	if i.Distribution == "" {
		return errors.New(DistributionCantBeNilMessage)
	}
	if i.Commit == nil || i.Commit.Arch == "" {
		return errors.New(ArchitectureCantBeEmptyMessage)
	}
	return nil
}
