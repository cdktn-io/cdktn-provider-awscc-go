// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccsagemakercomputequota

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawsccsagemakercomputequota/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference interface {
	cdktn.ComplexObject
	AcceleratorPartition() DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsAcceleratorPartitionOutputReference
	Accelerators() *float64
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
	Count() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InstanceType() *string
	InternalValue() *DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimits
	SetInternalValue(val *DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimits)
	MemoryInGiB() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VCpu() *float64
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

// The jsii proxy struct for DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference
type jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) AcceleratorPartition() DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsAcceleratorPartitionOutputReference {
	var returns DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsAcceleratorPartitionOutputReference
	_jsii_.Get(
		j,
		"acceleratorPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) Accelerators() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) InternalValue() *DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimits {
	var returns *DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) MemoryInGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) VCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vCpu",
		&returns,
	)
	return returns
}


func NewDataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSagemakerComputeQuota.DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference_Override(d DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSagemakerComputeQuota.DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference)SetInternalValue(val *DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccSagemakerComputeQuotaComputeQuotaConfigResourceSharingConfigAbsoluteBorrowLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

