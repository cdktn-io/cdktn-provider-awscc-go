// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3storagelens/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference interface {
	cdktn.ComplexObject
	ActivityMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelActivityMetricsOutputReference
	ActivityMetricsInput() interface{}
	AdvancedCostOptimizationMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsOutputReference
	AdvancedCostOptimizationMetricsInput() interface{}
	AdvancedDataProtectionMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsOutputReference
	AdvancedDataProtectionMetricsInput() interface{}
	AdvancedPerformanceMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference
	AdvancedPerformanceMetricsInput() interface{}
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
	DetailedStatusCodesMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodesMetricsOutputReference
	DetailedStatusCodesMetricsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrefixLevel() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelOutputReference
	PrefixLevelInput() interface{}
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
	PutActivityMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelActivityMetrics)
	PutAdvancedCostOptimizationMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics)
	PutAdvancedDataProtectionMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics)
	PutAdvancedPerformanceMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics)
	PutDetailedStatusCodesMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodesMetrics)
	PutPrefixLevel(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevel)
	ResetActivityMetrics()
	ResetAdvancedCostOptimizationMetrics()
	ResetAdvancedDataProtectionMetrics()
	ResetAdvancedPerformanceMetrics()
	ResetDetailedStatusCodesMetrics()
	ResetPrefixLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference
type jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ActivityMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelActivityMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelActivityMetricsOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ActivityMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedCostOptimizationMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedCostOptimizationMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedDataProtectionMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedDataProtectionMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedPerformanceMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) AdvancedPerformanceMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) DetailedStatusCodesMetrics() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodesMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodesMetricsOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodesMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) DetailedStatusCodesMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedStatusCodesMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PrefixLevel() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevelOutputReference
	_jsii_.Get(
		j,
		"prefixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PrefixLevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prefixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLens.S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference_Override(s S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLens.S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutActivityMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelActivityMetrics) {
	if err := s.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutAdvancedCostOptimizationMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics) {
	if err := s.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutAdvancedDataProtectionMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics) {
	if err := s.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutAdvancedPerformanceMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetrics) {
	if err := s.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutDetailedStatusCodesMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodesMetrics) {
	if err := s.validatePutDetailedStatusCodesMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDetailedStatusCodesMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) PutPrefixLevel(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevelPrefixLevel) {
	if err := s.validatePutPrefixLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPrefixLevel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetDetailedStatusCodesMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetDetailedStatusCodesMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ResetPrefixLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefixLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

