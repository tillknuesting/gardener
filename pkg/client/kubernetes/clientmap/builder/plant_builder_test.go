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

package builder

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	fakeclientmap "github.com/gardener/gardener/pkg/client/kubernetes/clientmap/fake"
	"github.com/gardener/gardener/pkg/client/kubernetes/clientmap/keys"
	fakeclientset "github.com/gardener/gardener/pkg/client/kubernetes/fake"
)

var _ = Describe("PlantClientMapBuilder", func() {

	var (
		ctx context.Context

		fakeGardenClientMap *fakeclientmap.ClientMap
		fakeGardenClientSet *fakeclientset.ClientSet
	)

	BeforeEach(func() {
		ctx = context.TODO()

		fakeGardenClientSet = fakeclientset.NewClientSet()
		fakeGardenClientMap = fakeclientmap.NewClientMapBuilder().WithClientSetForKey(keys.ForGarden(), fakeGardenClientSet).Build()
	})

	Context("#gardenClientFunc", func() {
		It("should be correctly set by WithGardenClientMap", func() {
			builder := NewPlantClientMapBuilder().WithGardenClientMap(fakeGardenClientMap)
			Expect(builder.gardenClientFunc(ctx)).To(BeEquivalentTo(fakeGardenClientSet))
		})

		It("should be correctly set by WithGardenClientSet", func() {
			builder := NewPlantClientMapBuilder().WithGardenClientSet(fakeGardenClientSet)
			Expect(builder.gardenClientFunc(ctx)).To(BeEquivalentTo(fakeGardenClientSet))
		})
	})

	Context("#Build", func() {
		It("should fail if garden ClientMap was not set", func() {
			clientMap, err := NewPlantClientMapBuilder().Build()
			Expect(err).To(MatchError("garden client is required but not set"))
			Expect(clientMap).To(BeNil())
		})

		It("should succeed to build ClientMap", func() {
			clientSet, err := NewPlantClientMapBuilder().
				WithGardenClientMap(fakeGardenClientMap).
				Build()
			Expect(err).NotTo(HaveOccurred())
			Expect(clientSet).NotTo(BeNil())
		})
	})

})
