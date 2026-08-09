// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelensgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3storagelensgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3StorageLensGroupFilterAndOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchAnyPrefix() *[]*string
	SetMatchAnyPrefix(val *[]*string)
	MatchAnyPrefixInput() *[]*string
	MatchAnySuffix() *[]*string
	SetMatchAnySuffix(val *[]*string)
	MatchAnySuffixInput() *[]*string
	MatchAnyTag() S3StorageLensGroupFilterAndMatchAnyTagList
	MatchAnyTagInput() interface{}
	MatchObjectAge() S3StorageLensGroupFilterAndMatchObjectAgeOutputReference
	MatchObjectAgeInput() interface{}
	MatchObjectSize() S3StorageLensGroupFilterAndMatchObjectSizeOutputReference
	MatchObjectSizeInput() interface{}
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
	PutMatchAnyTag(value interface{})
	PutMatchObjectAge(value *S3StorageLensGroupFilterAndMatchObjectAge)
	PutMatchObjectSize(value *S3StorageLensGroupFilterAndMatchObjectSize)
	ResetMatchAnyPrefix()
	ResetMatchAnySuffix()
	ResetMatchAnyTag()
	ResetMatchObjectAge()
	ResetMatchObjectSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensGroupFilterAndOutputReference
type jsiiProxy_S3StorageLensGroupFilterAndOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchAnyPrefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchAnyPrefixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchAnySuffix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnySuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchAnySuffixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnySuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchAnyTag() S3StorageLensGroupFilterAndMatchAnyTagList {
	var returns S3StorageLensGroupFilterAndMatchAnyTagList
	_jsii_.Get(
		j,
		"matchAnyTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchAnyTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchAnyTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchObjectAge() S3StorageLensGroupFilterAndMatchObjectAgeOutputReference {
	var returns S3StorageLensGroupFilterAndMatchObjectAgeOutputReference
	_jsii_.Get(
		j,
		"matchObjectAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchObjectAgeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchObjectAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchObjectSize() S3StorageLensGroupFilterAndMatchObjectSizeOutputReference {
	var returns S3StorageLensGroupFilterAndMatchObjectSizeOutputReference
	_jsii_.Get(
		j,
		"matchObjectSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) MatchObjectSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchObjectSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensGroupFilterAndOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3StorageLensGroupFilterAndOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensGroupFilterAndOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensGroupFilterAndOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLensGroup.S3StorageLensGroupFilterAndOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensGroupFilterAndOutputReference_Override(s S3StorageLensGroupFilterAndOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLensGroup.S3StorageLensGroupFilterAndOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference)SetMatchAnyPrefix(val *[]*string) {
	if err := j.validateSetMatchAnyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAnyPrefix",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference)SetMatchAnySuffix(val *[]*string) {
	if err := j.validateSetMatchAnySuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAnySuffix",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterAndOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) PutMatchAnyTag(value interface{}) {
	if err := s.validatePutMatchAnyTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchAnyTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) PutMatchObjectAge(value *S3StorageLensGroupFilterAndMatchObjectAge) {
	if err := s.validatePutMatchObjectAgeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchObjectAge",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) PutMatchObjectSize(value *S3StorageLensGroupFilterAndMatchObjectSize) {
	if err := s.validatePutMatchObjectSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchObjectSize",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ResetMatchAnyPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnyPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ResetMatchAnySuffix() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnySuffix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ResetMatchAnyTag() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnyTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ResetMatchObjectAge() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchObjectAge",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ResetMatchObjectSize() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchObjectSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3StorageLensGroupFilterAndOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

