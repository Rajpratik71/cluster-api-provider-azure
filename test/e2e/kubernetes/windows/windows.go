// +build e2e

/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package windows

import (
	"fmt"
	"log"
)

type OSVersion string
type WindowsTestImages string

const (
	Unknown  = OSVersion("")
	LTSC2019 = OSVersion("2019")
)

const (
	IIS   = WindowsTestImages("IIS")
	Httpd = WindowsTestImages("httpd")
)

type WindowsImage struct {
	BaseImage string
	Tags      map[OSVersion]string
}

func (i *WindowsImage) GetImage(version OSVersion) string {
	tag, ok := i.Tags[version]
	if !ok {
		log.Printf("Warning: Tag for version %s not found for image %s", version, i.BaseImage)
	}
	return fmt.Sprintf("%s:%s", i.BaseImage, tag)
}

func GetWindowsImage(testImage WindowsTestImages, version OSVersion) string {
	iisImage := WindowsImage{
		BaseImage: "mcr.microsoft.com/windows/servercore/iis",
		Tags: map[OSVersion]string{
			LTSC2019: "windowsservercore-ltsc2019",
		},
	}

	httpd := WindowsImage{
		BaseImage: "k8sprow.azurecr.io/kubernetes-e2e-test-images/httpd",
		Tags: map[OSVersion]string{
			LTSC2019: "2.4.39-alpine",
		},
	}

	switch testImage {
	case IIS:
		return iisImage.GetImage(version)
	case Httpd:
		return httpd.GetImage(version)
	default:
		return iisImage.GetImage(version)
	}
}
