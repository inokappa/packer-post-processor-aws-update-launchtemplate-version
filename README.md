# packer-post-processor-aws-update-launchtemplate-version

## Description

* Packer post-processor plugin for EC2 Launch Template create new version.

## Installation

```sh
$ mkdir -p ~/.packer.d/plugins
$ wget https://github.com/inokappa/packer-post-processor-aws-update-launchtemplate-version/releases/download/v0.0.1/packer-post-processor-aws-update-launchtemplate-version_darwin_amd64 -O ~/.packer.d/plugins/packer-post-processor-aws-update-launchtemplate-version
```

## Usage

The following example is a template to create New Launch template version.(Source template: `latest`)

```json
{
  "variables": {
    "aws_access_key": "{{env `AWS_ACCESS_KEY_ID`}}",
    "aws_secret_key": "{{env `AWS_SECRET_ACCESS_KEY`}}",
    "ssh_keypair_name": "{{env `KEYPAIR_NAME`}}",
    "ssh_private_key_file": "{{env `PRIVATE_KEY_PATH`}}",
    "ami_name": "packer-sample-{{timestamp}}"
  },
  "builders": [
    {
      "type": "amazon-ebs",
... snip ....
    }
  ],
  "provisioners": [
    {
      "type": "shell",
      "inline": [
        "echo 'shell'"
      ]
    }
  ],
  "post-processors": [
    [
      {
        "type": "aws-update-launchtemplate-version",
        "templates": [
          {
            "id": "lt-xxxxxxxxxxxxxxxxxxxx1",
            "source_version": "latest",
            "version_descripiton": "foo"
          },
          {
            "id": "lt-xxxxxxxxxxxxxxxxxxxx2",
            "source_version": "default",
            "version_descripiton": "bar"
          }
        ]
      }
    ]
  ]
}
```
