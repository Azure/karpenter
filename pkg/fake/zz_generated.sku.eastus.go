//go:build !ignore_autogenerated

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

package fake

import (
	"github.com/samber/lo"
	// nolint SA1019 - deprecated package
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2022-08-01/compute"
)

// generated at 2023-12-18T15:28:53Z

func init() {
	// ResourceSkus is a list of selected VM SKUs for a given region
	ResourceSkus["eastus"] = []compute.ResourceSku{
		{
			Name:         lo.ToPtr("Standard_A0"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("A0"),
			Family:       lo.ToPtr("standardA0_A7Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{
				{
					Type:   compute.ResourceSkuRestrictionsType("Location"),
					Values: &[]string{"eastus"},
					RestrictionInfo: &compute.ResourceSkuRestrictionInfo{
						Locations: &[]string{
							"eastus",
						},
						Zones: &[]string{},
					},
					ReasonCode: "NotAvailableForSubscription",
				},
				{
					Type:   compute.ResourceSkuRestrictionsType("Zone"),
					Values: &[]string{"eastus"},
					RestrictionInfo: &compute.ResourceSkuRestrictionInfo{
						Locations: &[]string{
							"eastus",
						},
						Zones: &[]string{
							"1",
							"2",
							"3",
						},
					},
					ReasonCode: "NotAvailableForSubscription",
				},
			},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("20480")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("0.75")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS,PaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("50")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations:    &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{}}},
		},
		{
			Name:         lo.ToPtr("Standard_B1s"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("B1s"),
			Family:       lo.ToPtr("standardBSFamily"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("4096")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("400")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("22528000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("22528000")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("32212254720")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("320")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("22500000")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_D2s_v3"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("D2s_v3"),
			Family:       lo.ToPtr("standardDSv3Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("16384")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("4")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("160")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("4000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("33554432")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("33554432")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("53687091200")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("3200")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("50331648")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_D2_v2"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("D2_v2"),
			Family:       lo.ToPtr("standardDv2Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("102400")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("7")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS,PaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("210")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("6000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("97517568")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("48234496")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_D2_v3"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("D2_v3"),
			Family:       lo.ToPtr("standardDv3Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("51200")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("4")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS,PaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("160")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("3000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("48234496")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("24117248")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_D2_v5"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("D2_v5"),
			Family:       lo.ToPtr("standardDv5Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("0")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("4")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("19000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("131072000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("131072000")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("3750")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("89128960")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_D4s_v3"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("D4s_v3"),
			Family:       lo.ToPtr("standardDSv3Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("32768")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("4")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("16")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("4")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("160")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("8000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("67108864")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("67108864")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("107374182400")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("6400")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("100663296")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_D64s_v3"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("D64s_v3"),
			Family:       lo.ToPtr("standardDSv3Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("524288")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("64")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("256")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("32")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("64")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("160")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("128000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("1073741824")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("1073741824")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("1717986918400")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("80000")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("1258291200")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("8")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_DC8s_v3"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("DC8s_v3"),
			Family:       lo.ToPtr("standardDCSv3Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("0")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("64")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("32")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("77000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("77000000000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("77000000000")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("25600")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("384000000")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("8")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_DS2_v2"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("DS2_v2"),
			Family:       lo.ToPtr("standardDSv2Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("14336")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("7")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("210")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("8000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("65536000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("65536000")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("92341796864")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("6400")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("96000000")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_F16s_v2"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("F16s_v2"),
			Family:       lo.ToPtr("standardFSv2Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("131072")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("16")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("32")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("32")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("16")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("195")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("32000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("262144000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("262144000")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("274877906944")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("25600")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("384000000")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("4")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_M8-2ms"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("M8-2ms"),
			Family:       lo.ToPtr("standardMSFamily"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{
				{
					Type:   compute.ResourceSkuRestrictionsType("Zone"),
					Values: &[]string{"eastus"},
					RestrictionInfo: &compute.ResourceSkuRestrictionInfo{
						Locations: &[]string{
							"eastus",
						},
						Zones: &[]string{
							"3",
						},
					},
					ReasonCode: "NotAvailableForSubscription",
				},
			},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("256000")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("218.75")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("MaxWriteAcceleratorDisksAllowed"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("ACUs"), Value: lo.ToPtr("160")},
				{Name: lo.ToPtr("ParentSize"), Value: lo.ToPtr("Standard_M8ms")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("2")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("10000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("104857600")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("104857600")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("851477266432")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("5000")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("131072000")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("TrustedLaunchDisabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_NC16as_T4_v3"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("NC16as_T4_v3"),
			Family:       lo.ToPtr("Standard NCASv3_T4 Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("360448")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("16")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("110")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("32")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("16")},
				{Name: lo.ToPtr("GPUs"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("16320")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("251658240")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("251658240")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("154619000000")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("24480")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("368640000")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("8")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_NC24ads_A100_v4"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("NC24ads_A100_v4"),
			Family:       lo.ToPtr("StandardNCADSA100v4Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("65536")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("24")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("220")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("8")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("24")},
				{Name: lo.ToPtr("GPUs"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedIOPS"), Value: lo.ToPtr("75000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedReadBytesPerSecond"), Value: lo.ToPtr("1000000000")},
				{Name: lo.ToPtr("CombinedTempDiskAndCachedWriteBytesPerSecond"), Value: lo.ToPtr("1000000000")},
				{Name: lo.ToPtr("CachedDiskBytes"), Value: lo.ToPtr("274877906944")},
				{Name: lo.ToPtr("UncachedDiskIOPS"), Value: lo.ToPtr("25600")},
				{Name: lo.ToPtr("UncachedDiskBytesPerSecond"), Value: lo.ToPtr("384000000")},
				{Name: lo.ToPtr("NvmeDiskSizeInMiB"), Value: lo.ToPtr("915456")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("2")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"3",
			},
			},
			},
		},
		{
			Name:         lo.ToPtr("Standard_NC6s_v3"),
			Tier:         lo.ToPtr("Standard"),
			Kind:         lo.ToPtr(""),
			Size:         lo.ToPtr("NC6s_v3"),
			Family:       lo.ToPtr("standardNCSv3Family"),
			ResourceType: lo.ToPtr("virtualMachines"),
			APIVersions:  &[]string{},
			Costs:        &[]compute.ResourceSkuCosts{},
			Restrictions: &[]compute.ResourceSkuRestrictions{
				{
					Type:   compute.ResourceSkuRestrictionsType("Zone"),
					Values: &[]string{"eastus"},
					RestrictionInfo: &compute.ResourceSkuRestrictionInfo{
						Locations: &[]string{
							"eastus",
						},
						Zones: &[]string{
							"1",
							"2",
							"3",
						},
					},
					ReasonCode: "NotAvailableForSubscription",
				},
			},
			Capabilities: &[]compute.ResourceSkuCapabilities{
				{Name: lo.ToPtr("MaxResourceVolumeMB"), Value: lo.ToPtr("344064")},
				{Name: lo.ToPtr("OSVhdSizeMB"), Value: lo.ToPtr("1047552")},
				{Name: lo.ToPtr("vCPUs"), Value: lo.ToPtr("6")},
				{Name: lo.ToPtr("MemoryPreservingMaintenanceSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("HyperVGenerations"), Value: lo.ToPtr("V1,V2")},
				{Name: lo.ToPtr("MemoryGB"), Value: lo.ToPtr("112")},
				{Name: lo.ToPtr("MaxDataDiskCount"), Value: lo.ToPtr("12")},
				{Name: lo.ToPtr("CpuArchitectureType"), Value: lo.ToPtr("x64")},
				{Name: lo.ToPtr("LowPriorityCapable"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("PremiumIO"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("VMDeploymentTypes"), Value: lo.ToPtr("IaaS")},
				{Name: lo.ToPtr("vCPUsAvailable"), Value: lo.ToPtr("6")},
				{Name: lo.ToPtr("GPUs"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("vCPUsPerCore"), Value: lo.ToPtr("1")},
				{Name: lo.ToPtr("EphemeralOSDiskSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("EncryptionAtHostSupported"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("CapacityReservationSupported"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("AcceleratedNetworkingEnabled"), Value: lo.ToPtr("True")},
				{Name: lo.ToPtr("RdmaEnabled"), Value: lo.ToPtr("False")},
				{Name: lo.ToPtr("MaxNetworkInterfaces"), Value: lo.ToPtr("4")},
			},
			Locations: &[]string{"eastus"},
			LocationInfo: &[]compute.ResourceSkuLocationInfo{{Location: lo.ToPtr("eastus"), Zones: &[]string{
				"1",
				"2",
				"3",
			},
			},
			},
		},
	}
}
