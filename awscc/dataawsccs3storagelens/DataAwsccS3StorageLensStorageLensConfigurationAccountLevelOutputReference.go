// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccs3storagelens

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccs3storagelens/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference interface {
	cdktn.ComplexObject
	ActivityMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelActivityMetricsOutputReference
	AdvancedCostOptimizationMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference
	AdvancedDataProtectionMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference
	AdvancedPerformanceMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetricsOutputReference
	BucketLevel() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference
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
	DetailedStatusCodesMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetricsOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccS3StorageLensStorageLensConfigurationAccountLevel
	SetInternalValue(val *DataAwsccS3StorageLensStorageLensConfigurationAccountLevel)
	StorageLensGroupLevel() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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

// The jsii proxy struct for DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference
type jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) ActivityMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelActivityMetricsOutputReference {
	var returns DataAwsccS3StorageLensStorageLensConfigurationAccountLevelActivityMetricsOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedCostOptimizationMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference {
	var returns DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedDataProtectionMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference {
	var returns DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedPerformanceMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetricsOutputReference {
	var returns DataAwsccS3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) BucketLevel() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference {
	var returns DataAwsccS3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference
	_jsii_.Get(
		j,
		"bucketLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) DetailedStatusCodesMetrics() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetricsOutputReference {
	var returns DataAwsccS3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetricsOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodesMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) InternalValue() *DataAwsccS3StorageLensStorageLensConfigurationAccountLevel {
	var returns *DataAwsccS3StorageLensStorageLensConfigurationAccountLevel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) StorageLensGroupLevel() DataAwsccS3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelOutputReference {
	var returns DataAwsccS3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelOutputReference
	_jsii_.Get(
		j,
		"storageLensGroupLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccS3StorageLens.DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference_Override(d DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccS3StorageLens.DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetInternalValue(val *DataAwsccS3StorageLensStorageLensConfigurationAccountLevel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccS3StorageLensStorageLensConfigurationAccountLevelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

