package main

import (
	_ "fmt"
	_ "os"
	"sort"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/awserr"
	_ "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetLatestLaunchTemplateVersion(launchTemplateId string, verSpec string) string {
	svc := ec2.New(session.New())
	input := &ec2.DescribeLaunchTemplateVersionsInput{
		LaunchTemplateId: aws.String(launchTemplateId),
	}

	result, err := svc.DescribeLaunchTemplateVersions(input)
	if err != nil {
		// Error の原因を返すようにしたい
		return ""
	}

	var versions []int64
	for _, r := range result.LaunchTemplateVersions {
		if verSpec == "default" {
			if *r.DefaultVersion {
				return strconv.FormatInt(*r.VersionNumber, 10)
			}
		} else {
			versions = append(versions, *r.VersionNumber)
		}
	}

	var versionNumber string
	// sortedVersion := sort.Sort(sort.Reverse(sort.IntSlice(versions)))
	sort.Slice(versions, func(i, j int) bool {
		return versions[i] < versions[j]
	})
	if verSpec == "latest" {
		versionNumber = strconv.FormatInt(versions[len(versions)-1], 10)
	} else {
		versionNumber = verSpec
	}

	return versionNumber
}
