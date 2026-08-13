// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/batchcomputeenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchComputeEnvironmentComputeResourcesOutputReference interface {
	cdktn.ComplexObject
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
	BidPercentage() *float64
	SetBidPercentage(val *float64)
	BidPercentageInput() *float64
	CapacityTags() *map[string]*string
	SetCapacityTags(val *map[string]*string)
	CapacityTagsInput() *map[string]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DesiredvCpus() *float64
	SetDesiredvCpus(val *float64)
	DesiredvCpusInput() *float64
	Ec2Configuration() BatchComputeEnvironmentComputeResourcesEc2ConfigurationList
	Ec2ConfigurationInput() interface{}
	Ec2KeyPair() *string
	SetEc2KeyPair(val *string)
	Ec2KeyPairInput() *string
	// Experimental.
	Fqn() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceRole() *string
	SetInstanceRole(val *string)
	InstanceRoleInput() *string
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	InstanceTypesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchTemplate() BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference
	LaunchTemplateInput() interface{}
	ManagedInstancesProvider() BatchComputeEnvironmentComputeResourcesManagedInstancesProviderOutputReference
	ManagedInstancesProviderInput() interface{}
	MaxvCpus() *float64
	SetMaxvCpus(val *float64)
	MaxvCpusInput() *float64
	MinvCpus() *float64
	SetMinvCpus(val *float64)
	MinvCpusInput() *float64
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	PlacementGroupInput() *string
	ScalingPolicy() BatchComputeEnvironmentComputeResourcesScalingPolicyOutputReference
	ScalingPolicyInput() interface{}
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SpotIamFleetRole() *string
	SetSpotIamFleetRole(val *string)
	SpotIamFleetRoleInput() *string
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	SubnetsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdateToLatestImageVersion() interface{}
	SetUpdateToLatestImageVersion(val interface{})
	UpdateToLatestImageVersionInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutEc2Configuration(value interface{})
	PutLaunchTemplate(value *BatchComputeEnvironmentComputeResourcesLaunchTemplate)
	PutManagedInstancesProvider(value *BatchComputeEnvironmentComputeResourcesManagedInstancesProvider)
	PutScalingPolicy(value *BatchComputeEnvironmentComputeResourcesScalingPolicy)
	ResetAllocationStrategy()
	ResetBidPercentage()
	ResetCapacityTags()
	ResetDesiredvCpus()
	ResetEc2Configuration()
	ResetEc2KeyPair()
	ResetImageId()
	ResetInstanceRole()
	ResetInstanceTypes()
	ResetLaunchTemplate()
	ResetManagedInstancesProvider()
	ResetMaxvCpus()
	ResetMinvCpus()
	ResetPlacementGroup()
	ResetScalingPolicy()
	ResetSecurityGroupIds()
	ResetSpotIamFleetRole()
	ResetSubnets()
	ResetTags()
	ResetType()
	ResetUpdateToLatestImageVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchComputeEnvironmentComputeResourcesOutputReference
type jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) BidPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) BidPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) CapacityTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacityTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) CapacityTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacityTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) DesiredvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) DesiredvCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredvCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2Configuration() BatchComputeEnvironmentComputeResourcesEc2ConfigurationList {
	var returns BatchComputeEnvironmentComputeResourcesEc2ConfigurationList
	_jsii_.Get(
		j,
		"ec2Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) LaunchTemplate() BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference {
	var returns BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) LaunchTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ManagedInstancesProvider() BatchComputeEnvironmentComputeResourcesManagedInstancesProviderOutputReference {
	var returns BatchComputeEnvironmentComputeResourcesManagedInstancesProviderOutputReference
	_jsii_.Get(
		j,
		"managedInstancesProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ManagedInstancesProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInstancesProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MaxvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MaxvCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxvCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MinvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MinvCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minvCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ScalingPolicy() BatchComputeEnvironmentComputeResourcesScalingPolicyOutputReference {
	var returns BatchComputeEnvironmentComputeResourcesScalingPolicyOutputReference
	_jsii_.Get(
		j,
		"scalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ScalingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SpotIamFleetRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotIamFleetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SpotIamFleetRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotIamFleetRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) UpdateToLatestImageVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateToLatestImageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) UpdateToLatestImageVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateToLatestImageVersionInput",
		&returns,
	)
	return returns
}


func NewBatchComputeEnvironmentComputeResourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BatchComputeEnvironmentComputeResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewBatchComputeEnvironmentComputeResourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.batchComputeEnvironment.BatchComputeEnvironmentComputeResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesOutputReference_Override(b BatchComputeEnvironmentComputeResourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.batchComputeEnvironment.BatchComputeEnvironmentComputeResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetBidPercentage(val *float64) {
	if err := j.validateSetBidPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPercentage",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetCapacityTags(val *map[string]*string) {
	if err := j.validateSetCapacityTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityTags",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetDesiredvCpus(val *float64) {
	if err := j.validateSetDesiredvCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredvCpus",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetEc2KeyPair(val *string) {
	if err := j.validateSetEc2KeyPairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2KeyPair",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetInstanceRole(val *string) {
	if err := j.validateSetInstanceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetMaxvCpus(val *float64) {
	if err := j.validateSetMaxvCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxvCpus",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetMinvCpus(val *float64) {
	if err := j.validateSetMinvCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minvCpus",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetSpotIamFleetRole(val *string) {
	if err := j.validateSetSpotIamFleetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotIamFleetRole",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference)SetUpdateToLatestImageVersion(val interface{}) {
	if err := j.validateSetUpdateToLatestImageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateToLatestImageVersion",
		val,
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PutEc2Configuration(value interface{}) {
	if err := b.validatePutEc2ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEc2Configuration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PutLaunchTemplate(value *BatchComputeEnvironmentComputeResourcesLaunchTemplate) {
	if err := b.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PutManagedInstancesProvider(value *BatchComputeEnvironmentComputeResourcesManagedInstancesProvider) {
	if err := b.validatePutManagedInstancesProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putManagedInstancesProvider",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PutScalingPolicy(value *BatchComputeEnvironmentComputeResourcesScalingPolicy) {
	if err := b.validatePutScalingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putScalingPolicy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetBidPercentage() {
	_jsii_.InvokeVoid(
		b,
		"resetBidPercentage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetCapacityTags() {
	_jsii_.InvokeVoid(
		b,
		"resetCapacityTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetDesiredvCpus() {
	_jsii_.InvokeVoid(
		b,
		"resetDesiredvCpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetEc2Configuration() {
	_jsii_.InvokeVoid(
		b,
		"resetEc2Configuration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetEc2KeyPair() {
	_jsii_.InvokeVoid(
		b,
		"resetEc2KeyPair",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetImageId() {
	_jsii_.InvokeVoid(
		b,
		"resetImageId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetInstanceRole() {
	_jsii_.InvokeVoid(
		b,
		"resetInstanceRole",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		b,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetManagedInstancesProvider() {
	_jsii_.InvokeVoid(
		b,
		"resetManagedInstancesProvider",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetMaxvCpus() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxvCpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetMinvCpus() {
	_jsii_.InvokeVoid(
		b,
		"resetMinvCpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		b,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetScalingPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetScalingPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		b,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetSpotIamFleetRole() {
	_jsii_.InvokeVoid(
		b,
		"resetSpotIamFleetRole",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetSubnets() {
	_jsii_.InvokeVoid(
		b,
		"resetSubnets",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetUpdateToLatestImageVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetUpdateToLatestImageVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

