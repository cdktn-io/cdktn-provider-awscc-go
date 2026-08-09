// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3storagelens/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3StorageLensStorageLensConfigurationAccountLevelOutputReference interface {
	cdktn.ComplexObject
	ActivityMetrics() S3StorageLensStorageLensConfigurationAccountLevelActivityMetricsOutputReference
	ActivityMetricsInput() interface{}
	AdvancedCostOptimizationMetrics() S3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference
	AdvancedCostOptimizationMetricsInput() interface{}
	AdvancedDataProtectionMetrics() S3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference
	AdvancedDataProtectionMetricsInput() interface{}
	AdvancedPerformanceMetrics() S3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetricsOutputReference
	AdvancedPerformanceMetricsInput() interface{}
	BucketLevel() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference
	BucketLevelInput() interface{}
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
	DetailedStatusCodesMetrics() S3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetricsOutputReference
	DetailedStatusCodesMetricsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageLensGroupLevel() S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelOutputReference
	StorageLensGroupLevelInput() interface{}
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
	PutActivityMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelActivityMetrics)
	PutAdvancedCostOptimizationMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics)
	PutAdvancedDataProtectionMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics)
	PutAdvancedPerformanceMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetrics)
	PutBucketLevel(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevel)
	PutDetailedStatusCodesMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetrics)
	PutStorageLensGroupLevel(value *S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevel)
	ResetActivityMetrics()
	ResetAdvancedCostOptimizationMetrics()
	ResetAdvancedDataProtectionMetrics()
	ResetAdvancedPerformanceMetrics()
	ResetDetailedStatusCodesMetrics()
	ResetStorageLensGroupLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensStorageLensConfigurationAccountLevelOutputReference
type jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ActivityMetrics() S3StorageLensStorageLensConfigurationAccountLevelActivityMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelActivityMetricsOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ActivityMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedCostOptimizationMetrics() S3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedCostOptimizationMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedDataProtectionMetrics() S3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedDataProtectionMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedPerformanceMetrics() S3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetricsOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) AdvancedPerformanceMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) BucketLevel() S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelBucketLevelOutputReference
	_jsii_.Get(
		j,
		"bucketLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) BucketLevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) DetailedStatusCodesMetrics() S3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetricsOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetricsOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodesMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) DetailedStatusCodesMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedStatusCodesMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) StorageLensGroupLevel() S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelOutputReference
	_jsii_.Get(
		j,
		"storageLensGroupLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) StorageLensGroupLevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageLensGroupLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensStorageLensConfigurationAccountLevelOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3StorageLensStorageLensConfigurationAccountLevelOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensStorageLensConfigurationAccountLevelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLens.S3StorageLensStorageLensConfigurationAccountLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensStorageLensConfigurationAccountLevelOutputReference_Override(s S3StorageLensStorageLensConfigurationAccountLevelOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLens.S3StorageLensStorageLensConfigurationAccountLevelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) PutActivityMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelActivityMetrics) {
	if err := s.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) PutAdvancedCostOptimizationMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics) {
	if err := s.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) PutAdvancedDataProtectionMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics) {
	if err := s.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) PutAdvancedPerformanceMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelAdvancedPerformanceMetrics) {
	if err := s.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) PutBucketLevel(value *S3StorageLensStorageLensConfigurationAccountLevelBucketLevel) {
	if err := s.validatePutBucketLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBucketLevel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) PutDetailedStatusCodesMetrics(value *S3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetrics) {
	if err := s.validatePutDetailedStatusCodesMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDetailedStatusCodesMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) PutStorageLensGroupLevel(value *S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevel) {
	if err := s.validatePutStorageLensGroupLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStorageLensGroupLevel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ResetDetailedStatusCodesMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetDetailedStatusCodesMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ResetStorageLensGroupLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageLensGroupLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationAccountLevelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

