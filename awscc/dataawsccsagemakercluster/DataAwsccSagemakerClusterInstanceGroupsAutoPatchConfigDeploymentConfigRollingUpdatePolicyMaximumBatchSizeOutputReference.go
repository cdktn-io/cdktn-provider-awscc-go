// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccsagemakercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccsagemakercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSize
	SetInternalValue(val *DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSize)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	Value() *float64
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

// The jsii proxy struct for DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference
type jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) InternalValue() *DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSize {
	var returns *DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSize
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewDataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSagemakerCluster.DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference_Override(d DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSagemakerCluster.DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference)SetInternalValue(val *DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSize) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccSagemakerClusterInstanceGroupsAutoPatchConfigDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

