/*
Portions Copyright (c) Microsoft Corporation.

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

package utilization_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/aws/karpenter-core/pkg/test"

	"github.com/Azure/karpenter/pkg/apis/v1alpha2"
	"github.com/Azure/karpenter/test/pkg/environment/azure"
)

var env *azure.Environment

func TestUtilization(t *testing.T) {
	RegisterFailHandler(Fail)
	BeforeSuite(func() {
		env = azure.NewEnvironment(t)
	})
	RunSpecs(t, "Utilization")
}

var _ = BeforeEach(func() { env.BeforeEach() })
var _ = AfterEach(func() { env.Cleanup() })
var _ = AfterEach(func() { env.AfterEach() })

// TODO (charliedmcb): add back in referenced to debug lib. For now I'm commenting out all references for this test, to avoid vetting the lib
// var _ = Describe("Utilization", Label(debug.NoWatch), Label(debug.NoEvents), func() {
var _ = Describe("Utilization", func() {
	It("should provision one pod per node", func() {
		nodeClass := env.DefaultAKSNodeClass()
		nodePool := env.DefaultNodePool(nodeClass)
		test.ReplaceRequirements(nodePool, v1.NodeSelectorRequirement{
			Key:      v1alpha2.LabelSKUCPU,
			Operator: v1.NodeSelectorOpLt,
			Values:   []string{"3"},
		})
		deployment := test.Deployment(test.DeploymentOptions{
			Replicas:   10,
			PodOptions: test.PodOptions{ResourceRequirements: v1.ResourceRequirements{Requests: v1.ResourceList{v1.ResourceCPU: resource.MustParse("1.1")}}, Image: "mcr.microsoft.com/oss/kubernetes/pause:3.6"}}) // See above comment on the required adding of the azure PodOptions.Image reference

		env.ExpectCreated(nodePool, nodeClass, deployment)
		env.EventuallyExpectHealthyPodCount(labels.SelectorFromSet(deployment.Spec.Selector.MatchLabels), int(*deployment.Spec.Replicas))
		env.ExpectCreatedNodeCount("==", int(*deployment.Spec.Replicas)) // One pod per node enforced by instance size
	})
})
