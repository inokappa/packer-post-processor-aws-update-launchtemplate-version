//go:generate mapstructure-to-hcl2 -type Config
package main

import (
	awscommon "github.com/hashicorp/packer/builder/amazon/common"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig    `mapstructure:",squash"`
	awscommon.AccessConfig `mapstructure:",squash"`

	Templates []struct {
		Id                 string `mapstructure:"id"`
		SourceVersion      string `mapstructure:"source_version"`
		VersionDescription string `mapstructure:"version_description"`
	} `mapstructure:"templates"`

	ctx interpolate.Context
}
