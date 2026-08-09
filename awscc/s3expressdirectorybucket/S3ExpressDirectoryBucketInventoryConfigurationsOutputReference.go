// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressdirectorybucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3expressdirectorybucket/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3ExpressDirectoryBucketInventoryConfigurationsOutputReference interface {
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
	Destination() S3ExpressDirectoryBucketInventoryConfigurationsDestinationOutputReference
	DestinationInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludedObjectVersions() *string
	SetIncludedObjectVersions(val *string)
	IncludedObjectVersionsInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OptionalFields() *[]*string
	SetOptionalFields(val *[]*string)
	OptionalFieldsInput() *[]*string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	ScheduleFrequency() *string
	SetScheduleFrequency(val *string)
	ScheduleFrequencyInput() *string
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
	PutDestination(value *S3ExpressDirectoryBucketInventoryConfigurationsDestination)
	ResetDestination()
	ResetEnabled()
	ResetId()
	ResetIncludedObjectVersions()
	ResetOptionalFields()
	ResetPrefix()
	ResetScheduleFrequency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ExpressDirectoryBucketInventoryConfigurationsOutputReference
type jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) Destination() S3ExpressDirectoryBucketInventoryConfigurationsDestinationOutputReference {
	var returns S3ExpressDirectoryBucketInventoryConfigurationsDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) DestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) IncludedObjectVersions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includedObjectVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) IncludedObjectVersionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includedObjectVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) OptionalFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optionalFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) OptionalFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optionalFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ScheduleFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ScheduleFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ExpressDirectoryBucketInventoryConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3ExpressDirectoryBucketInventoryConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewS3ExpressDirectoryBucketInventoryConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3ExpressDirectoryBucket.S3ExpressDirectoryBucketInventoryConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3ExpressDirectoryBucketInventoryConfigurationsOutputReference_Override(s S3ExpressDirectoryBucketInventoryConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3ExpressDirectoryBucket.S3ExpressDirectoryBucketInventoryConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetIncludedObjectVersions(val *string) {
	if err := j.validateSetIncludedObjectVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedObjectVersions",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetOptionalFields(val *[]*string) {
	if err := j.validateSetOptionalFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalFields",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetScheduleFrequency(val *string) {
	if err := j.validateSetScheduleFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleFrequency",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) PutDestination(value *S3ExpressDirectoryBucketInventoryConfigurationsDestination) {
	if err := s.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ResetDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ResetIncludedObjectVersions() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludedObjectVersions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ResetOptionalFields() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionalFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ResetScheduleFrequency() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduleFrequency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketInventoryConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

