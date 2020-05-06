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

	TemplateId            string `mapstructure:"template_id"`
	SourceTemplateVersion string `mapstructure:"source_template_version"`
	VersionDescription    string `mapstructure:"version_description"`

	ctx interpolate.Context
}