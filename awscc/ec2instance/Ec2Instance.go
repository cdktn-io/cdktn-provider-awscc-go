// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2instance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance awscc_ec2_instance}.
type Ec2Instance interface {
	cdktn.TerraformResource
	AdditionalInfo() *string
	SetAdditionalInfo(val *string)
	AdditionalInfoInput() *string
	Affinity() *string
	SetAffinity(val *string)
	AffinityInput() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BlockDeviceMappings() Ec2InstanceBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuOptions() Ec2InstanceCpuOptionsOutputReference
	CpuOptionsInput() interface{}
	CreditSpecification() Ec2InstanceCreditSpecificationOutputReference
	CreditSpecificationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableApiTermination() interface{}
	SetDisableApiTermination(val interface{})
	DisableApiTerminationInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	ElasticGpuSpecifications() Ec2InstanceElasticGpuSpecificationsList
	ElasticGpuSpecificationsInput() interface{}
	ElasticInferenceAccelerators() Ec2InstanceElasticInferenceAcceleratorsList
	ElasticInferenceAcceleratorsInput() interface{}
	EnclaveOptions() Ec2InstanceEnclaveOptionsOutputReference
	EnclaveOptionsInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HibernationOptions() Ec2InstanceHibernationOptionsOutputReference
	HibernationOptionsInput() interface{}
	HostId() *string
	SetHostId(val *string)
	HostIdInput() *string
	HostResourceGroupArn() *string
	SetHostResourceGroupArn(val *string)
	HostResourceGroupArnInput() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	Id() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceId() *string
	InstanceInitiatedShutdownBehavior() *string
	SetInstanceInitiatedShutdownBehavior(val *string)
	InstanceInitiatedShutdownBehaviorInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	Ipv6AddressCountInput() *float64
	Ipv6Addresses() Ec2InstanceIpv6AddressesList
	Ipv6AddressesInput() interface{}
	KernelId() *string
	SetKernelId(val *string)
	KernelIdInput() *string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	LaunchTemplate() Ec2InstanceLaunchTemplateOutputReference
	LaunchTemplateInput() interface{}
	LicenseSpecifications() Ec2InstanceLicenseSpecificationsList
	LicenseSpecificationsInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MetadataOptions() Ec2InstanceMetadataOptionsOutputReference
	MetadataOptionsInput() interface{}
	Monitoring() interface{}
	SetMonitoring(val interface{})
	MonitoringInput() interface{}
	NetworkInterfaces() Ec2InstanceNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	// The tree node.
	Node() constructs.Node
	PlacementGroupName() *string
	SetPlacementGroupName(val *string)
	PlacementGroupNameInput() *string
	PrivateDnsName() *string
	PrivateDnsNameOptions() Ec2InstancePrivateDnsNameOptionsOutputReference
	PrivateDnsNameOptionsInput() interface{}
	PrivateIp() *string
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddressInput() *string
	PropagateTagsToVolumeOnCreation() interface{}
	SetPropagateTagsToVolumeOnCreation(val interface{})
	PropagateTagsToVolumeOnCreationInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicDnsName() *string
	PublicIp() *string
	RamdiskId() *string
	SetRamdiskId(val *string)
	RamdiskIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SourceDestCheck() interface{}
	SetSourceDestCheck(val interface{})
	SourceDestCheckInput() interface{}
	SsmAssociations() Ec2InstanceSsmAssociationsList
	SsmAssociationsInput() interface{}
	State() Ec2InstanceStateOutputReference
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() Ec2InstanceTagsList
	TagsInput() interface{}
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	Volumes() Ec2InstanceVolumesList
	VolumesInput() interface{}
	VpcId() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBlockDeviceMappings(value interface{})
	PutCpuOptions(value *Ec2InstanceCpuOptions)
	PutCreditSpecification(value *Ec2InstanceCreditSpecification)
	PutElasticGpuSpecifications(value interface{})
	PutElasticInferenceAccelerators(value interface{})
	PutEnclaveOptions(value *Ec2InstanceEnclaveOptions)
	PutHibernationOptions(value *Ec2InstanceHibernationOptions)
	PutIpv6Addresses(value interface{})
	PutLaunchTemplate(value *Ec2InstanceLaunchTemplate)
	PutLicenseSpecifications(value interface{})
	PutMetadataOptions(value *Ec2InstanceMetadataOptions)
	PutNetworkInterfaces(value interface{})
	PutPrivateDnsNameOptions(value *Ec2InstancePrivateDnsNameOptions)
	PutSsmAssociations(value interface{})
	PutTags(value interface{})
	PutVolumes(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAdditionalInfo()
	ResetAffinity()
	ResetAvailabilityZone()
	ResetBlockDeviceMappings()
	ResetCpuOptions()
	ResetCreditSpecification()
	ResetDisableApiTermination()
	ResetEbsOptimized()
	ResetElasticGpuSpecifications()
	ResetElasticInferenceAccelerators()
	ResetEnclaveOptions()
	ResetHibernationOptions()
	ResetHostId()
	ResetHostResourceGroupArn()
	ResetIamInstanceProfile()
	ResetImageId()
	ResetInstanceInitiatedShutdownBehavior()
	ResetInstanceType()
	ResetIpv6AddressCount()
	ResetIpv6Addresses()
	ResetKernelId()
	ResetKeyName()
	ResetLaunchTemplate()
	ResetLicenseSpecifications()
	ResetMetadataOptions()
	ResetMonitoring()
	ResetNetworkInterfaces()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementGroupName()
	ResetPrivateDnsNameOptions()
	ResetPrivateIpAddress()
	ResetPropagateTagsToVolumeOnCreation()
	ResetRamdiskId()
	ResetSecurityGroupIds()
	ResetSecurityGroups()
	ResetSourceDestCheck()
	ResetSsmAssociations()
	ResetSubnetId()
	ResetTags()
	ResetTenancy()
	ResetUserData()
	ResetVolumes()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for Ec2Instance
type jsiiProxy_Ec2Instance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Ec2Instance) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) AdditionalInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Affinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) AffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) BlockDeviceMappings() Ec2InstanceBlockDeviceMappingsList {
	var returns Ec2InstanceBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) CpuOptions() Ec2InstanceCpuOptionsOutputReference {
	var returns Ec2InstanceCpuOptionsOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) CpuOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) CreditSpecification() Ec2InstanceCreditSpecificationOutputReference {
	var returns Ec2InstanceCreditSpecificationOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) CreditSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ElasticGpuSpecifications() Ec2InstanceElasticGpuSpecificationsList {
	var returns Ec2InstanceElasticGpuSpecificationsList
	_jsii_.Get(
		j,
		"elasticGpuSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ElasticGpuSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticGpuSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ElasticInferenceAccelerators() Ec2InstanceElasticInferenceAcceleratorsList {
	var returns Ec2InstanceElasticInferenceAcceleratorsList
	_jsii_.Get(
		j,
		"elasticInferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ElasticInferenceAcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticInferenceAcceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) EnclaveOptions() Ec2InstanceEnclaveOptionsOutputReference {
	var returns Ec2InstanceEnclaveOptionsOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) EnclaveOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) HibernationOptions() Ec2InstanceHibernationOptionsOutputReference {
	var returns Ec2InstanceHibernationOptionsOutputReference
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) HibernationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) HostResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Ipv6Addresses() Ec2InstanceIpv6AddressesList {
	var returns Ec2InstanceIpv6AddressesList
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Ipv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) LaunchTemplate() Ec2InstanceLaunchTemplateOutputReference {
	var returns Ec2InstanceLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) LaunchTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) LicenseSpecifications() Ec2InstanceLicenseSpecificationsList {
	var returns Ec2InstanceLicenseSpecificationsList
	_jsii_.Get(
		j,
		"licenseSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) LicenseSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) MetadataOptions() Ec2InstanceMetadataOptionsOutputReference {
	var returns Ec2InstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) MetadataOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) NetworkInterfaces() Ec2InstanceNetworkInterfacesList {
	var returns Ec2InstanceNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PlacementGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PlacementGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PrivateDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PrivateDnsNameOptions() Ec2InstancePrivateDnsNameOptionsOutputReference {
	var returns Ec2InstancePrivateDnsNameOptionsOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PrivateDnsNameOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PropagateTagsToVolumeOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTagsToVolumeOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PropagateTagsToVolumeOnCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTagsToVolumeOnCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PublicDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) RamdiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) RamdiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SourceDestCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SourceDestCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SsmAssociations() Ec2InstanceSsmAssociationsList {
	var returns Ec2InstanceSsmAssociationsList
	_jsii_.Get(
		j,
		"ssmAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SsmAssociationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssmAssociationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) State() Ec2InstanceStateOutputReference {
	var returns Ec2InstanceStateOutputReference
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Tags() Ec2InstanceTagsList {
	var returns Ec2InstanceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) Volumes() Ec2InstanceVolumesList {
	var returns Ec2InstanceVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Instance) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance awscc_ec2_instance} Resource.
func NewEc2Instance(scope constructs.Construct, id *string, config *Ec2InstanceConfig) Ec2Instance {
	_init_.Initialize()

	if err := validateNewEc2InstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Instance{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Instance.Ec2Instance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance awscc_ec2_instance} Resource.
func NewEc2Instance_Override(e Ec2Instance, scope constructs.Construct, id *string, config *Ec2InstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Instance.Ec2Instance",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2Instance)SetAdditionalInfo(val *string) {
	if err := j.validateSetAdditionalInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetAffinity(val *string) {
	if err := j.validateSetAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affinity",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetHostResourceGroupArn(val *string) {
	if err := j.validateSetHostResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetPlacementGroupName(val *string) {
	if err := j.validateSetPlacementGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroupName",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetPropagateTagsToVolumeOnCreation(val interface{}) {
	if err := j.validateSetPropagateTagsToVolumeOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTagsToVolumeOnCreation",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetRamdiskId(val *string) {
	if err := j.validateSetRamdiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramdiskId",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetSourceDestCheck(val interface{}) {
	if err := j.validateSetSourceDestCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDestCheck",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_Ec2Instance)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

// Generates CDKTN code for importing a Ec2Instance resource upon running "cdktn plan <stack-name>".
func Ec2Instance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEc2Instance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2Instance.Ec2Instance",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Ec2Instance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Instance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2Instance.Ec2Instance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Instance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Instance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2Instance.Ec2Instance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Instance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Instance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2Instance.Ec2Instance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2Instance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.ec2Instance.Ec2Instance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2Instance) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2Instance) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2Instance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2Instance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2Instance) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2Instance) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2Instance) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2Instance) PutBlockDeviceMappings(value interface{}) {
	if err := e.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutCpuOptions(value *Ec2InstanceCpuOptions) {
	if err := e.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutCreditSpecification(value *Ec2InstanceCreditSpecification) {
	if err := e.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutElasticGpuSpecifications(value interface{}) {
	if err := e.validatePutElasticGpuSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putElasticGpuSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutElasticInferenceAccelerators(value interface{}) {
	if err := e.validatePutElasticInferenceAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putElasticInferenceAccelerators",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutEnclaveOptions(value *Ec2InstanceEnclaveOptions) {
	if err := e.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutHibernationOptions(value *Ec2InstanceHibernationOptions) {
	if err := e.validatePutHibernationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHibernationOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutIpv6Addresses(value interface{}) {
	if err := e.validatePutIpv6AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv6Addresses",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutLaunchTemplate(value *Ec2InstanceLaunchTemplate) {
	if err := e.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutLicenseSpecifications(value interface{}) {
	if err := e.validatePutLicenseSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLicenseSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutMetadataOptions(value *Ec2InstanceMetadataOptions) {
	if err := e.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutNetworkInterfaces(value interface{}) {
	if err := e.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutPrivateDnsNameOptions(value *Ec2InstancePrivateDnsNameOptions) {
	if err := e.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutSsmAssociations(value interface{}) {
	if err := e.validatePutSsmAssociationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSsmAssociations",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) PutVolumes(value interface{}) {
	if err := e.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVolumes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Instance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_Ec2Instance) ResetAdditionalInfo() {
	_jsii_.InvokeVoid(
		e,
		"resetAdditionalInfo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetAffinity() {
	_jsii_.InvokeVoid(
		e,
		"resetAffinity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		e,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		e,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetElasticGpuSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticGpuSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetElasticInferenceAccelerators() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticInferenceAccelerators",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetHibernationOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetHibernationOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetHostId() {
	_jsii_.InvokeVoid(
		e,
		"resetHostId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetHostResourceGroupArn() {
	_jsii_.InvokeVoid(
		e,
		"resetHostResourceGroupArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		e,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetImageId() {
	_jsii_.InvokeVoid(
		e,
		"resetImageId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetKernelId() {
	_jsii_.InvokeVoid(
		e,
		"resetKernelId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetKeyName() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetLicenseSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetLicenseSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetMonitoring() {
	_jsii_.InvokeVoid(
		e,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetPlacementGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetPropagateTagsToVolumeOnCreation() {
	_jsii_.InvokeVoid(
		e,
		"resetPropagateTagsToVolumeOnCreation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetRamdiskId() {
	_jsii_.InvokeVoid(
		e,
		"resetRamdiskId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetSourceDestCheck() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceDestCheck",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetSsmAssociations() {
	_jsii_.InvokeVoid(
		e,
		"resetSsmAssociations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetSubnetId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetTenancy() {
	_jsii_.InvokeVoid(
		e,
		"resetTenancy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetUserData() {
	_jsii_.InvokeVoid(
		e,
		"resetUserData",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) ResetVolumes() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Instance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Instance) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

