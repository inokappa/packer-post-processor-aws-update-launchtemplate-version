package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2/hcldec"
	awscommon "github.com/hashicorp/packer/builder/amazon/common"
	_ "github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	_ "github.com/hashicorp/packer/template/interpolate"
	_ "os"
	_ "regexp"
	"strings"
)

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	p.config.ctx.Funcs = awscommon.TemplateFuncs
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
	}, raws...)

	if err != nil {
		return err
	}

	for _, tpl := range p.config.Templates {
		if tpl.Id == "" {
			return errors.New("empty `template_id` is not allowed. Please make sure that it is set correctly")
		}

		if tpl.SourceVersion == "" {
			return errors.New("empty `source_template_version` is not allowed. Please make sure that it is set correctly")
		}
	}

	return nil
}

func (p *PostProcessor) PostProcess(ctx context.Context, ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, bool, error) {
	session, err := p.config.Session()
	if err != nil {
		return nil, false, false, err
	}
	config := session.Config

	amiId := p.GetImageId(artifact)
	for _, tpl := range p.config.Templates {
		srcVer := GetLatestLaunchTemplateVersion(tpl.Id, tpl.SourceVersion)
		if tpl.VersionDescription == "" {
			tpl.VersionDescription = fmt.Sprintf("This launchtemplate uses source version %s.", srcVer)
		}

		message := fmt.Sprintf("Creating New Launch Template Version (Template ID: %s Source Version: %s, AMI ID: %s)", tpl.Id, srcVer, amiId)
		ui.Message(message)
		_, err = CreateLaunchTemplateVersion(amiId, tpl.Id, srcVer, tpl.VersionDescription)
		if err != nil {
			return nil, true, false, err
		}
	}

	artifact = &awscommon.Artifact{
		Amis: map[string]string{
			*config.Region: amiId,
		},
	}

	return artifact, true, false, nil
	// return artifact, true, nil
}

func (p *PostProcessor) GetImageId(artifact packer.Artifact) string {
	// example: ap-northeast-1:ami-xxxxxxxxxxxxxxxx
	splitedString := strings.Split(artifact.Id(), ":")
	amiId := splitedString[len(splitedString)-1]

	return amiId
}
