// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccodbodbnetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccodbodbnetwork/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccOdbOdbNetworkManagedServicesOutputReference interface {
	cdktn.ComplexObject
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
	CrossRegionS3RestoreSourcesAccess() DataAwsccOdbOdbNetworkManagedServicesCrossRegionS3RestoreSourcesAccessList
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccOdbOdbNetworkManagedServices
	SetInternalValue(val *DataAwsccOdbOdbNetworkManagedServices)
	KmsAccess() DataAwsccOdbOdbNetworkManagedServicesKmsAccessOutputReference
	ManagedS3BackupAccess() DataAwsccOdbOdbNetworkManagedServicesManagedS3BackupAccessOutputReference
	ManagedServicesIpv4Cidrs() *[]*string
	ResourceGatewayArn() *string
	S3Access() DataAwsccOdbOdbNetworkManagedServicesS3AccessOutputReference
	ServiceNetworkArn() *string
	ServiceNetworkEndpoint() DataAwsccOdbOdbNetworkManagedServicesServiceNetworkEndpointOutputReference
	StsAccess() DataAwsccOdbOdbNetworkManagedServicesStsAccessOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ZeroEtlAccess() DataAwsccOdbOdbNetworkManagedServicesZeroEtlAccessOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccOdbOdbNetworkManagedServicesOutputReference
type jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) CrossRegionS3RestoreSourcesAccess() DataAwsccOdbOdbNetworkManagedServicesCrossRegionS3RestoreSourcesAccessList {
	var returns DataAwsccOdbOdbNetworkManagedServicesCrossRegionS3RestoreSourcesAccessList
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) InternalValue() *DataAwsccOdbOdbNetworkManagedServices {
	var returns *DataAwsccOdbOdbNetworkManagedServices
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) KmsAccess() DataAwsccOdbOdbNetworkManagedServicesKmsAccessOutputReference {
	var returns DataAwsccOdbOdbNetworkManagedServicesKmsAccessOutputReference
	_jsii_.Get(
		j,
		"kmsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ManagedS3BackupAccess() DataAwsccOdbOdbNetworkManagedServicesManagedS3BackupAccessOutputReference {
	var returns DataAwsccOdbOdbNetworkManagedServicesManagedS3BackupAccessOutputReference
	_jsii_.Get(
		j,
		"managedS3BackupAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ManagedServicesIpv4Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managedServicesIpv4Cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ResourceGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) S3Access() DataAwsccOdbOdbNetworkManagedServicesS3AccessOutputReference {
	var returns DataAwsccOdbOdbNetworkManagedServicesS3AccessOutputReference
	_jsii_.Get(
		j,
		"s3Access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ServiceNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ServiceNetworkEndpoint() DataAwsccOdbOdbNetworkManagedServicesServiceNetworkEndpointOutputReference {
	var returns DataAwsccOdbOdbNetworkManagedServicesServiceNetworkEndpointOutputReference
	_jsii_.Get(
		j,
		"serviceNetworkEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) StsAccess() DataAwsccOdbOdbNetworkManagedServicesStsAccessOutputReference {
	var returns DataAwsccOdbOdbNetworkManagedServicesStsAccessOutputReference
	_jsii_.Get(
		j,
		"stsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ZeroEtlAccess() DataAwsccOdbOdbNetworkManagedServicesZeroEtlAccessOutputReference {
	var returns DataAwsccOdbOdbNetworkManagedServicesZeroEtlAccessOutputReference
	_jsii_.Get(
		j,
		"zeroEtlAccess",
		&returns,
	)
	return returns
}


func NewDataAwsccOdbOdbNetworkManagedServicesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccOdbOdbNetworkManagedServicesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccOdbOdbNetworkManagedServicesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOdbOdbNetwork.DataAwsccOdbOdbNetworkManagedServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccOdbOdbNetworkManagedServicesOutputReference_Override(d DataAwsccOdbOdbNetworkManagedServicesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOdbOdbNetwork.DataAwsccOdbOdbNetworkManagedServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference)SetInternalValue(val *DataAwsccOdbOdbNetworkManagedServices) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbOdbNetworkManagedServicesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

