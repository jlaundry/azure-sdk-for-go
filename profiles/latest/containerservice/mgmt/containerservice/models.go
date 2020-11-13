// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package containerservice

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2020-11-01/containerservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AgentPoolMode = original.AgentPoolMode

const (
	System AgentPoolMode = original.System
	User   AgentPoolMode = original.User
)

type AgentPoolType = original.AgentPoolType

const (
	AvailabilitySet         AgentPoolType = original.AvailabilitySet
	VirtualMachineScaleSets AgentPoolType = original.VirtualMachineScaleSets
)

type Code = original.Code

const (
	Running Code = original.Running
	Stopped Code = original.Stopped
)

type ConnectionStatus = original.ConnectionStatus

const (
	Approved     ConnectionStatus = original.Approved
	Disconnected ConnectionStatus = original.Disconnected
	Pending      ConnectionStatus = original.Pending
	Rejected     ConnectionStatus = original.Rejected
)

type Expander = original.Expander

const (
	LeastWaste Expander = original.LeastWaste
	MostPods   Expander = original.MostPods
	Random     Expander = original.Random
)

type LicenseType = original.LicenseType

const (
	None          LicenseType = original.None
	WindowsServer LicenseType = original.WindowsServer
)

type LoadBalancerSku = original.LoadBalancerSku

const (
	Basic    LoadBalancerSku = original.Basic
	Standard LoadBalancerSku = original.Standard
)

type ManagedClusterPodIdentityProvisioningState = original.ManagedClusterPodIdentityProvisioningState

const (
	Assigned ManagedClusterPodIdentityProvisioningState = original.Assigned
	Deleting ManagedClusterPodIdentityProvisioningState = original.Deleting
	Failed   ManagedClusterPodIdentityProvisioningState = original.Failed
	Updating ManagedClusterPodIdentityProvisioningState = original.Updating
)

type ManagedClusterSKUName = original.ManagedClusterSKUName

const (
	ManagedClusterSKUNameBasic ManagedClusterSKUName = original.ManagedClusterSKUNameBasic
)

type ManagedClusterSKUTier = original.ManagedClusterSKUTier

const (
	Free ManagedClusterSKUTier = original.Free
	Paid ManagedClusterSKUTier = original.Paid
)

type NetworkMode = original.NetworkMode

const (
	Bridge      NetworkMode = original.Bridge
	Transparent NetworkMode = original.Transparent
)

type NetworkPlugin = original.NetworkPlugin

const (
	Azure   NetworkPlugin = original.Azure
	Kubenet NetworkPlugin = original.Kubenet
)

type NetworkPolicy = original.NetworkPolicy

const (
	NetworkPolicyAzure  NetworkPolicy = original.NetworkPolicyAzure
	NetworkPolicyCalico NetworkPolicy = original.NetworkPolicyCalico
)

type OSDiskType = original.OSDiskType

const (
	Ephemeral OSDiskType = original.Ephemeral
	Managed   OSDiskType = original.Managed
)

type OSType = original.OSType

const (
	Linux   OSType = original.Linux
	Windows OSType = original.Windows
)

type OutboundType = original.OutboundType

const (
	LoadBalancer       OutboundType = original.LoadBalancer
	UserDefinedRouting OutboundType = original.UserDefinedRouting
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicy

const (
	Deallocate ScaleSetEvictionPolicy = original.Deallocate
	Delete     ScaleSetEvictionPolicy = original.Delete
)

type ScaleSetPriority = original.ScaleSetPriority

const (
	Regular ScaleSetPriority = original.Regular
	Spot    ScaleSetPriority = original.Spot
)

type StorageProfileTypes = original.StorageProfileTypes

const (
	ManagedDisks   StorageProfileTypes = original.ManagedDisks
	StorageAccount StorageProfileTypes = original.StorageAccount
)

type UpgradeChannel = original.UpgradeChannel

const (
	UpgradeChannelNone   UpgradeChannel = original.UpgradeChannelNone
	UpgradeChannelPatch  UpgradeChannel = original.UpgradeChannelPatch
	UpgradeChannelRapid  UpgradeChannel = original.UpgradeChannelRapid
	UpgradeChannelStable UpgradeChannel = original.UpgradeChannelStable
)

type VMSizeTypes = original.VMSizeTypes

const (
	StandardA1          VMSizeTypes = original.StandardA1
	StandardA10         VMSizeTypes = original.StandardA10
	StandardA11         VMSizeTypes = original.StandardA11
	StandardA1V2        VMSizeTypes = original.StandardA1V2
	StandardA2          VMSizeTypes = original.StandardA2
	StandardA2mV2       VMSizeTypes = original.StandardA2mV2
	StandardA2V2        VMSizeTypes = original.StandardA2V2
	StandardA3          VMSizeTypes = original.StandardA3
	StandardA4          VMSizeTypes = original.StandardA4
	StandardA4mV2       VMSizeTypes = original.StandardA4mV2
	StandardA4V2        VMSizeTypes = original.StandardA4V2
	StandardA5          VMSizeTypes = original.StandardA5
	StandardA6          VMSizeTypes = original.StandardA6
	StandardA7          VMSizeTypes = original.StandardA7
	StandardA8          VMSizeTypes = original.StandardA8
	StandardA8mV2       VMSizeTypes = original.StandardA8mV2
	StandardA8V2        VMSizeTypes = original.StandardA8V2
	StandardA9          VMSizeTypes = original.StandardA9
	StandardB2ms        VMSizeTypes = original.StandardB2ms
	StandardB2s         VMSizeTypes = original.StandardB2s
	StandardB4ms        VMSizeTypes = original.StandardB4ms
	StandardB8ms        VMSizeTypes = original.StandardB8ms
	StandardD1          VMSizeTypes = original.StandardD1
	StandardD11         VMSizeTypes = original.StandardD11
	StandardD11V2       VMSizeTypes = original.StandardD11V2
	StandardD11V2Promo  VMSizeTypes = original.StandardD11V2Promo
	StandardD12         VMSizeTypes = original.StandardD12
	StandardD12V2       VMSizeTypes = original.StandardD12V2
	StandardD12V2Promo  VMSizeTypes = original.StandardD12V2Promo
	StandardD13         VMSizeTypes = original.StandardD13
	StandardD13V2       VMSizeTypes = original.StandardD13V2
	StandardD13V2Promo  VMSizeTypes = original.StandardD13V2Promo
	StandardD14         VMSizeTypes = original.StandardD14
	StandardD14V2       VMSizeTypes = original.StandardD14V2
	StandardD14V2Promo  VMSizeTypes = original.StandardD14V2Promo
	StandardD15V2       VMSizeTypes = original.StandardD15V2
	StandardD16sV3      VMSizeTypes = original.StandardD16sV3
	StandardD16V3       VMSizeTypes = original.StandardD16V3
	StandardD1V2        VMSizeTypes = original.StandardD1V2
	StandardD2          VMSizeTypes = original.StandardD2
	StandardD2sV3       VMSizeTypes = original.StandardD2sV3
	StandardD2V2        VMSizeTypes = original.StandardD2V2
	StandardD2V2Promo   VMSizeTypes = original.StandardD2V2Promo
	StandardD2V3        VMSizeTypes = original.StandardD2V3
	StandardD3          VMSizeTypes = original.StandardD3
	StandardD32sV3      VMSizeTypes = original.StandardD32sV3
	StandardD32V3       VMSizeTypes = original.StandardD32V3
	StandardD3V2        VMSizeTypes = original.StandardD3V2
	StandardD3V2Promo   VMSizeTypes = original.StandardD3V2Promo
	StandardD4          VMSizeTypes = original.StandardD4
	StandardD4sV3       VMSizeTypes = original.StandardD4sV3
	StandardD4V2        VMSizeTypes = original.StandardD4V2
	StandardD4V2Promo   VMSizeTypes = original.StandardD4V2Promo
	StandardD4V3        VMSizeTypes = original.StandardD4V3
	StandardD5V2        VMSizeTypes = original.StandardD5V2
	StandardD5V2Promo   VMSizeTypes = original.StandardD5V2Promo
	StandardD64sV3      VMSizeTypes = original.StandardD64sV3
	StandardD64V3       VMSizeTypes = original.StandardD64V3
	StandardD8sV3       VMSizeTypes = original.StandardD8sV3
	StandardD8V3        VMSizeTypes = original.StandardD8V3
	StandardDS1         VMSizeTypes = original.StandardDS1
	StandardDS11        VMSizeTypes = original.StandardDS11
	StandardDS11V2      VMSizeTypes = original.StandardDS11V2
	StandardDS11V2Promo VMSizeTypes = original.StandardDS11V2Promo
	StandardDS12        VMSizeTypes = original.StandardDS12
	StandardDS12V2      VMSizeTypes = original.StandardDS12V2
	StandardDS12V2Promo VMSizeTypes = original.StandardDS12V2Promo
	StandardDS13        VMSizeTypes = original.StandardDS13
	StandardDS132V2     VMSizeTypes = original.StandardDS132V2
	StandardDS134V2     VMSizeTypes = original.StandardDS134V2
	StandardDS13V2      VMSizeTypes = original.StandardDS13V2
	StandardDS13V2Promo VMSizeTypes = original.StandardDS13V2Promo
	StandardDS14        VMSizeTypes = original.StandardDS14
	StandardDS144V2     VMSizeTypes = original.StandardDS144V2
	StandardDS148V2     VMSizeTypes = original.StandardDS148V2
	StandardDS14V2      VMSizeTypes = original.StandardDS14V2
	StandardDS14V2Promo VMSizeTypes = original.StandardDS14V2Promo
	StandardDS15V2      VMSizeTypes = original.StandardDS15V2
	StandardDS1V2       VMSizeTypes = original.StandardDS1V2
	StandardDS2         VMSizeTypes = original.StandardDS2
	StandardDS2V2       VMSizeTypes = original.StandardDS2V2
	StandardDS2V2Promo  VMSizeTypes = original.StandardDS2V2Promo
	StandardDS3         VMSizeTypes = original.StandardDS3
	StandardDS3V2       VMSizeTypes = original.StandardDS3V2
	StandardDS3V2Promo  VMSizeTypes = original.StandardDS3V2Promo
	StandardDS4         VMSizeTypes = original.StandardDS4
	StandardDS4V2       VMSizeTypes = original.StandardDS4V2
	StandardDS4V2Promo  VMSizeTypes = original.StandardDS4V2Promo
	StandardDS5V2       VMSizeTypes = original.StandardDS5V2
	StandardDS5V2Promo  VMSizeTypes = original.StandardDS5V2Promo
	StandardE16sV3      VMSizeTypes = original.StandardE16sV3
	StandardE16V3       VMSizeTypes = original.StandardE16V3
	StandardE2sV3       VMSizeTypes = original.StandardE2sV3
	StandardE2V3        VMSizeTypes = original.StandardE2V3
	StandardE3216sV3    VMSizeTypes = original.StandardE3216sV3
	StandardE328sV3     VMSizeTypes = original.StandardE328sV3
	StandardE32sV3      VMSizeTypes = original.StandardE32sV3
	StandardE32V3       VMSizeTypes = original.StandardE32V3
	StandardE4sV3       VMSizeTypes = original.StandardE4sV3
	StandardE4V3        VMSizeTypes = original.StandardE4V3
	StandardE6416sV3    VMSizeTypes = original.StandardE6416sV3
	StandardE6432sV3    VMSizeTypes = original.StandardE6432sV3
	StandardE64sV3      VMSizeTypes = original.StandardE64sV3
	StandardE64V3       VMSizeTypes = original.StandardE64V3
	StandardE8sV3       VMSizeTypes = original.StandardE8sV3
	StandardE8V3        VMSizeTypes = original.StandardE8V3
	StandardF1          VMSizeTypes = original.StandardF1
	StandardF16         VMSizeTypes = original.StandardF16
	StandardF16s        VMSizeTypes = original.StandardF16s
	StandardF16sV2      VMSizeTypes = original.StandardF16sV2
	StandardF1s         VMSizeTypes = original.StandardF1s
	StandardF2          VMSizeTypes = original.StandardF2
	StandardF2s         VMSizeTypes = original.StandardF2s
	StandardF2sV2       VMSizeTypes = original.StandardF2sV2
	StandardF32sV2      VMSizeTypes = original.StandardF32sV2
	StandardF4          VMSizeTypes = original.StandardF4
	StandardF4s         VMSizeTypes = original.StandardF4s
	StandardF4sV2       VMSizeTypes = original.StandardF4sV2
	StandardF64sV2      VMSizeTypes = original.StandardF64sV2
	StandardF72sV2      VMSizeTypes = original.StandardF72sV2
	StandardF8          VMSizeTypes = original.StandardF8
	StandardF8s         VMSizeTypes = original.StandardF8s
	StandardF8sV2       VMSizeTypes = original.StandardF8sV2
	StandardG1          VMSizeTypes = original.StandardG1
	StandardG2          VMSizeTypes = original.StandardG2
	StandardG3          VMSizeTypes = original.StandardG3
	StandardG4          VMSizeTypes = original.StandardG4
	StandardG5          VMSizeTypes = original.StandardG5
	StandardGS1         VMSizeTypes = original.StandardGS1
	StandardGS2         VMSizeTypes = original.StandardGS2
	StandardGS3         VMSizeTypes = original.StandardGS3
	StandardGS4         VMSizeTypes = original.StandardGS4
	StandardGS44        VMSizeTypes = original.StandardGS44
	StandardGS48        VMSizeTypes = original.StandardGS48
	StandardGS5         VMSizeTypes = original.StandardGS5
	StandardGS516       VMSizeTypes = original.StandardGS516
	StandardGS58        VMSizeTypes = original.StandardGS58
	StandardH16         VMSizeTypes = original.StandardH16
	StandardH16m        VMSizeTypes = original.StandardH16m
	StandardH16mr       VMSizeTypes = original.StandardH16mr
	StandardH16r        VMSizeTypes = original.StandardH16r
	StandardH8          VMSizeTypes = original.StandardH8
	StandardH8m         VMSizeTypes = original.StandardH8m
	StandardL16s        VMSizeTypes = original.StandardL16s
	StandardL32s        VMSizeTypes = original.StandardL32s
	StandardL4s         VMSizeTypes = original.StandardL4s
	StandardL8s         VMSizeTypes = original.StandardL8s
	StandardM12832ms    VMSizeTypes = original.StandardM12832ms
	StandardM12864ms    VMSizeTypes = original.StandardM12864ms
	StandardM128ms      VMSizeTypes = original.StandardM128ms
	StandardM128s       VMSizeTypes = original.StandardM128s
	StandardM6416ms     VMSizeTypes = original.StandardM6416ms
	StandardM6432ms     VMSizeTypes = original.StandardM6432ms
	StandardM64ms       VMSizeTypes = original.StandardM64ms
	StandardM64s        VMSizeTypes = original.StandardM64s
	StandardNC12        VMSizeTypes = original.StandardNC12
	StandardNC12sV2     VMSizeTypes = original.StandardNC12sV2
	StandardNC12sV3     VMSizeTypes = original.StandardNC12sV3
	StandardNC24        VMSizeTypes = original.StandardNC24
	StandardNC24r       VMSizeTypes = original.StandardNC24r
	StandardNC24rsV2    VMSizeTypes = original.StandardNC24rsV2
	StandardNC24rsV3    VMSizeTypes = original.StandardNC24rsV3
	StandardNC24sV2     VMSizeTypes = original.StandardNC24sV2
	StandardNC24sV3     VMSizeTypes = original.StandardNC24sV3
	StandardNC6         VMSizeTypes = original.StandardNC6
	StandardNC6sV2      VMSizeTypes = original.StandardNC6sV2
	StandardNC6sV3      VMSizeTypes = original.StandardNC6sV3
	StandardND12s       VMSizeTypes = original.StandardND12s
	StandardND24rs      VMSizeTypes = original.StandardND24rs
	StandardND24s       VMSizeTypes = original.StandardND24s
	StandardND6s        VMSizeTypes = original.StandardND6s
	StandardNV12        VMSizeTypes = original.StandardNV12
	StandardNV24        VMSizeTypes = original.StandardNV24
	StandardNV6         VMSizeTypes = original.StandardNV6
)

type AccessProfile = original.AccessProfile
type AgentPool = original.AgentPool
type AgentPoolAvailableVersions = original.AgentPoolAvailableVersions
type AgentPoolAvailableVersionsProperties = original.AgentPoolAvailableVersionsProperties
type AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem = original.AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem
type AgentPoolListResult = original.AgentPoolListResult
type AgentPoolListResultIterator = original.AgentPoolListResultIterator
type AgentPoolListResultPage = original.AgentPoolListResultPage
type AgentPoolUpgradeProfile = original.AgentPoolUpgradeProfile
type AgentPoolUpgradeProfileProperties = original.AgentPoolUpgradeProfileProperties
type AgentPoolUpgradeProfilePropertiesUpgradesItem = original.AgentPoolUpgradeProfilePropertiesUpgradesItem
type AgentPoolUpgradeSettings = original.AgentPoolUpgradeSettings
type AgentPoolsClient = original.AgentPoolsClient
type AgentPoolsCreateOrUpdateFuture = original.AgentPoolsCreateOrUpdateFuture
type AgentPoolsDeleteFuture = original.AgentPoolsDeleteFuture
type AgentPoolsUpgradeNodeImageVersionFuture = original.AgentPoolsUpgradeNodeImageVersionFuture
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CredentialResult = original.CredentialResult
type CredentialResults = original.CredentialResults
type DiagnosticsProfile = original.DiagnosticsProfile
type KubeletConfig = original.KubeletConfig
type LinuxOSConfig = original.LinuxOSConfig
type LinuxProfile = original.LinuxProfile
type ManagedCluster = original.ManagedCluster
type ManagedClusterAADProfile = original.ManagedClusterAADProfile
type ManagedClusterAPIServerAccessProfile = original.ManagedClusterAPIServerAccessProfile
type ManagedClusterAccessProfile = original.ManagedClusterAccessProfile
type ManagedClusterAddonProfile = original.ManagedClusterAddonProfile
type ManagedClusterAddonProfileIdentity = original.ManagedClusterAddonProfileIdentity
type ManagedClusterAgentPoolProfile = original.ManagedClusterAgentPoolProfile
type ManagedClusterAgentPoolProfileProperties = original.ManagedClusterAgentPoolProfileProperties
type ManagedClusterAutoUpgradeProfile = original.ManagedClusterAutoUpgradeProfile
type ManagedClusterIdentity = original.ManagedClusterIdentity
type ManagedClusterIdentityUserAssignedIdentitiesValue = original.ManagedClusterIdentityUserAssignedIdentitiesValue
type ManagedClusterListResult = original.ManagedClusterListResult
type ManagedClusterListResultIterator = original.ManagedClusterListResultIterator
type ManagedClusterListResultPage = original.ManagedClusterListResultPage
type ManagedClusterLoadBalancerProfile = original.ManagedClusterLoadBalancerProfile
type ManagedClusterLoadBalancerProfileManagedOutboundIPs = original.ManagedClusterLoadBalancerProfileManagedOutboundIPs
type ManagedClusterLoadBalancerProfileOutboundIPPrefixes = original.ManagedClusterLoadBalancerProfileOutboundIPPrefixes
type ManagedClusterLoadBalancerProfileOutboundIPs = original.ManagedClusterLoadBalancerProfileOutboundIPs
type ManagedClusterPodIdentity = original.ManagedClusterPodIdentity
type ManagedClusterPodIdentityException = original.ManagedClusterPodIdentityException
type ManagedClusterPodIdentityProfile = original.ManagedClusterPodIdentityProfile
type ManagedClusterPodIdentityProvisioningInfo = original.ManagedClusterPodIdentityProvisioningInfo
type ManagedClusterPoolUpgradeProfile = original.ManagedClusterPoolUpgradeProfile
type ManagedClusterPoolUpgradeProfileUpgradesItem = original.ManagedClusterPoolUpgradeProfileUpgradesItem
type ManagedClusterProperties = original.ManagedClusterProperties
type ManagedClusterPropertiesAutoScalerProfile = original.ManagedClusterPropertiesAutoScalerProfile
type ManagedClusterPropertiesIdentityProfileValue = original.ManagedClusterPropertiesIdentityProfileValue
type ManagedClusterSKU = original.ManagedClusterSKU
type ManagedClusterServicePrincipalProfile = original.ManagedClusterServicePrincipalProfile
type ManagedClusterUpgradeProfile = original.ManagedClusterUpgradeProfile
type ManagedClusterUpgradeProfileProperties = original.ManagedClusterUpgradeProfileProperties
type ManagedClusterWindowsProfile = original.ManagedClusterWindowsProfile
type ManagedClustersClient = original.ManagedClustersClient
type ManagedClustersCreateOrUpdateFuture = original.ManagedClustersCreateOrUpdateFuture
type ManagedClustersDeleteFuture = original.ManagedClustersDeleteFuture
type ManagedClustersResetAADProfileFuture = original.ManagedClustersResetAADProfileFuture
type ManagedClustersResetServicePrincipalProfileFuture = original.ManagedClustersResetServicePrincipalProfileFuture
type ManagedClustersRotateClusterCertificatesFuture = original.ManagedClustersRotateClusterCertificatesFuture
type ManagedClustersStartFuture = original.ManagedClustersStartFuture
type ManagedClustersStopFuture = original.ManagedClustersStopFuture
type ManagedClustersUpdateTagsFuture = original.ManagedClustersUpdateTagsFuture
type MasterProfile = original.MasterProfile
type NetworkProfile = original.NetworkProfile
type OperationListResult = original.OperationListResult
type OperationValue = original.OperationValue
type OperationValueDisplay = original.OperationValueDisplay
type OperationsClient = original.OperationsClient
type PowerState = original.PowerState
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkResourcesListResult = original.PrivateLinkResourcesListResult
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ResolvePrivateLinkServiceIDClient = original.ResolvePrivateLinkServiceIDClient
type Resource = original.Resource
type ResourceReference = original.ResourceReference
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type SubResource = original.SubResource
type SysctlConfig = original.SysctlConfig
type TagsObject = original.TagsObject
type UserAssignedIdentity = original.UserAssignedIdentity
type VMDiagnostics = original.VMDiagnostics

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAgentPoolListResultIterator(page AgentPoolListResultPage) AgentPoolListResultIterator {
	return original.NewAgentPoolListResultIterator(page)
}
func NewAgentPoolListResultPage(getNextPage func(context.Context, AgentPoolListResult) (AgentPoolListResult, error)) AgentPoolListResultPage {
	return original.NewAgentPoolListResultPage(getNextPage)
}
func NewAgentPoolsClient(subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClient(subscriptionID)
}
func NewAgentPoolsClientWithBaseURI(baseURI string, subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedClusterListResultIterator(page ManagedClusterListResultPage) ManagedClusterListResultIterator {
	return original.NewManagedClusterListResultIterator(page)
}
func NewManagedClusterListResultPage(getNextPage func(context.Context, ManagedClusterListResult) (ManagedClusterListResult, error)) ManagedClusterListResultPage {
	return original.NewManagedClusterListResultPage(getNextPage)
}
func NewManagedClustersClient(subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClient(subscriptionID)
}
func NewManagedClustersClientWithBaseURI(baseURI string, subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResolvePrivateLinkServiceIDClient(subscriptionID string) ResolvePrivateLinkServiceIDClient {
	return original.NewResolvePrivateLinkServiceIDClient(subscriptionID)
}
func NewResolvePrivateLinkServiceIDClientWithBaseURI(baseURI string, subscriptionID string) ResolvePrivateLinkServiceIDClient {
	return original.NewResolvePrivateLinkServiceIDClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAgentPoolModeValues() []AgentPoolMode {
	return original.PossibleAgentPoolModeValues()
}
func PossibleAgentPoolTypeValues() []AgentPoolType {
	return original.PossibleAgentPoolTypeValues()
}
func PossibleCodeValues() []Code {
	return original.PossibleCodeValues()
}
func PossibleConnectionStatusValues() []ConnectionStatus {
	return original.PossibleConnectionStatusValues()
}
func PossibleExpanderValues() []Expander {
	return original.PossibleExpanderValues()
}
func PossibleLicenseTypeValues() []LicenseType {
	return original.PossibleLicenseTypeValues()
}
func PossibleLoadBalancerSkuValues() []LoadBalancerSku {
	return original.PossibleLoadBalancerSkuValues()
}
func PossibleManagedClusterPodIdentityProvisioningStateValues() []ManagedClusterPodIdentityProvisioningState {
	return original.PossibleManagedClusterPodIdentityProvisioningStateValues()
}
func PossibleManagedClusterSKUNameValues() []ManagedClusterSKUName {
	return original.PossibleManagedClusterSKUNameValues()
}
func PossibleManagedClusterSKUTierValues() []ManagedClusterSKUTier {
	return original.PossibleManagedClusterSKUTierValues()
}
func PossibleNetworkModeValues() []NetworkMode {
	return original.PossibleNetworkModeValues()
}
func PossibleNetworkPluginValues() []NetworkPlugin {
	return original.PossibleNetworkPluginValues()
}
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return original.PossibleNetworkPolicyValues()
}
func PossibleOSDiskTypeValues() []OSDiskType {
	return original.PossibleOSDiskTypeValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossibleOutboundTypeValues() []OutboundType {
	return original.PossibleOutboundTypeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleScaleSetEvictionPolicyValues() []ScaleSetEvictionPolicy {
	return original.PossibleScaleSetEvictionPolicyValues()
}
func PossibleScaleSetPriorityValues() []ScaleSetPriority {
	return original.PossibleScaleSetPriorityValues()
}
func PossibleStorageProfileTypesValues() []StorageProfileTypes {
	return original.PossibleStorageProfileTypesValues()
}
func PossibleUpgradeChannelValues() []UpgradeChannel {
	return original.PossibleUpgradeChannelValues()
}
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return original.PossibleVMSizeTypesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
