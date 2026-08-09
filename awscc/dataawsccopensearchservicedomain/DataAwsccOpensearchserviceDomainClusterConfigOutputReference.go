// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccopensearchservicedomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccopensearchservicedomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccOpensearchserviceDomainClusterConfigOutputReference interface {
	cdktn.ComplexObject
	ColdStorageOptions() DataAwsccOpensearchserviceDomainClusterConfigColdStorageOptionsOutputReference
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
	DedicatedMasterCount() *float64
	DedicatedMasterEnabled() cdktn.IResolvable
	DedicatedMasterType() *string
	// Experimental.
	Fqn() *string
	InstanceCount() *float64
	InstanceType() *string
	InternalValue() *DataAwsccOpensearchserviceDomainClusterConfig
	SetInternalValue(val *DataAwsccOpensearchserviceDomainClusterConfig)
	MultiAzWithStandbyEnabled() cdktn.IResolvable
	NodeOptions() DataAwsccOpensearchserviceDomainClusterConfigNodeOptionsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WarmCount() *float64
	WarmEnabled() cdktn.IResolvable
	WarmType() *string
	ZoneAwarenessConfig() DataAwsccOpensearchserviceDomainClusterConfigZoneAwarenessConfigOutputReference
	ZoneAwarenessEnabled() cdktn.IResolvable
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

// The jsii proxy struct for DataAwsccOpensearchserviceDomainClusterConfigOutputReference
type jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) ColdStorageOptions() DataAwsccOpensearchserviceDomainClusterConfigColdStorageOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainClusterConfigColdStorageOptionsOutputReference
	_jsii_.Get(
		j,
		"coldStorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) DedicatedMasterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) DedicatedMasterEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"dedicatedMasterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) DedicatedMasterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) InternalValue() *DataAwsccOpensearchserviceDomainClusterConfig {
	var returns *DataAwsccOpensearchserviceDomainClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) MultiAzWithStandbyEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"multiAzWithStandbyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) NodeOptions() DataAwsccOpensearchserviceDomainClusterConfigNodeOptionsList {
	var returns DataAwsccOpensearchserviceDomainClusterConfigNodeOptionsList
	_jsii_.Get(
		j,
		"nodeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) WarmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) WarmEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"warmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) WarmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) ZoneAwarenessConfig() DataAwsccOpensearchserviceDomainClusterConfigZoneAwarenessConfigOutputReference {
	var returns DataAwsccOpensearchserviceDomainClusterConfigZoneAwarenessConfigOutputReference
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) ZoneAwarenessEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}


func NewDataAwsccOpensearchserviceDomainClusterConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccOpensearchserviceDomainClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccOpensearchserviceDomainClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccOpensearchserviceDomainClusterConfigOutputReference_Override(d DataAwsccOpensearchserviceDomainClusterConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference)SetInternalValue(val *DataAwsccOpensearchserviceDomainClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomainClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

