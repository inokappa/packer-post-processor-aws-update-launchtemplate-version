package main

import (
	_ "fmt"
	_ "os"

	"github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/awserr"
	_ "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"strconv"
)

func CreateLaunchTemplateVersion(amiId string, launchTemplateId string, srcVersion string, versionDesc string) ([]string, error) {
	svc := ec2.New(session.New())
	input := &ec2.CreateLaunchTemplateVersionInput{
		LaunchTemplateData: &ec2.RequestLaunchTemplateData{
			ImageId: aws.String(amiId),
		},
		LaunchTemplateId:   aws.String(launchTemplateId),
		SourceVersion:      aws.String(srcVersion),
		VersionDescription: aws.String(versionDesc),
	}

	result, err := svc.CreateLaunchTemplateVersion(input)
	if err != nil {
		// fmt.Println(err.Error())
		return nil, err
	}

	res := []string{
		*result.LaunchTemplateVersion.LaunchTemplateData.ImageId,
		strconv.FormatInt(*result.LaunchTemplateVersion.VersionNumber, 10),
	}

	return res, nil
}
