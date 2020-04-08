// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator_test

import (
	gardenlinux_generator "github.com/gardener/gardener-extension-os-gardenlinux/pkg/generator"

	"github.com/gardener/gardener/extensions/pkg/controller/operatingsystemconfig/oscommon/generator/test"
	"github.com/gobuffalo/packr"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Garden Linux OS Generator Test", func() {

	Describe("Conformance Tests", func() {
		var box = packr.NewBox("./testfiles")
		test.DescribeTest(gardenlinux_generator.CloudInitGenerator(), box)()
	})
})
