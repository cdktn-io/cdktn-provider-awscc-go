// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelensgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3storagelensgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3StorageLensGroupFilterOutputReference interface {
	cdktn.ComplexObject
	And() S3StorageLensGroupFilterAndOutputReference
	AndInput() interface{}
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
	MatchAnyTag() S3StorageLensGroupFilterMatchAnyTagList
	MatchAnyTagInput() interface{}
	MatchObjectAge() S3StorageLensGroupFilterMatchObjectAgeOutputReference
	MatchObjectAgeInput() interface{}
	MatchObjectSize() S3StorageLensGroupFilterMatchObjectSizeOutputReference
	MatchObjectSizeInput() interface{}
	Or() S3StorageLensGroupFilterOrOutputReference
	OrInput() interface{}
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
	PutAnd(value *S3StorageLensGroupFilterAnd)
	PutMatchAnyTag(value interface{})
	PutMatchObjectAge(value *S3StorageLensGroupFilterMatchObjectAge)
	PutMatchObjectSize(value *S3StorageLensGroupFilterMatchObjectSize)
	PutOr(value *S3StorageLensGroupFilterOr)
	ResetAnd()
	ResetMatchAnyPrefix()
	ResetMatchAnySuffix()
	ResetMatchAnyTag()
	ResetMatchObjectAge()
	ResetMatchObjectSize()
	ResetOr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensGroupFilterOutputReference
type jsiiProxy_S3StorageLensGroupFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) And() S3StorageLensGroupFilterAndOutputReference {
	var returns S3StorageLensGroupFilterAndOutputReference
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchAnyPrefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchAnyPrefixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchAnySuffix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnySuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchAnySuffixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchAnySuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchAnyTag() S3StorageLensGroupFilterMatchAnyTagList {
	var returns S3StorageLensGroupFilterMatchAnyTagList
	_jsii_.Get(
		j,
		"matchAnyTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchAnyTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchAnyTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchObjectAge() S3StorageLensGroupFilterMatchObjectAgeOutputReference {
	var returns S3StorageLensGroupFilterMatchObjectAgeOutputReference
	_jsii_.Get(
		j,
		"matchObjectAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchObjectAgeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchObjectAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchObjectSize() S3StorageLensGroupFilterMatchObjectSizeOutputReference {
	var returns S3StorageLensGroupFilterMatchObjectSizeOutputReference
	_jsii_.Get(
		j,
		"matchObjectSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) MatchObjectSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchObjectSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) Or() S3StorageLensGroupFilterOrOutputReference {
	var returns S3StorageLensGroupFilterOrOutputReference
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensGroupFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) S3StorageLensGroupFilterOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensGroupFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensGroupFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLensGroup.S3StorageLensGroupFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensGroupFilterOutputReference_Override(s S3StorageLensGroupFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3StorageLensGroup.S3StorageLensGroupFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference)SetMatchAnyPrefix(val *[]*string) {
	if err := j.validateSetMatchAnyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAnyPrefix",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference)SetMatchAnySuffix(val *[]*string) {
	if err := j.validateSetMatchAnySuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchAnySuffix",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensGroupFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) PutAnd(value *S3StorageLensGroupFilterAnd) {
	if err := s.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAnd",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) PutMatchAnyTag(value interface{}) {
	if err := s.validatePutMatchAnyTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchAnyTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) PutMatchObjectAge(value *S3StorageLensGroupFilterMatchObjectAge) {
	if err := s.validatePutMatchObjectAgeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchObjectAge",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) PutMatchObjectSize(value *S3StorageLensGroupFilterMatchObjectSize) {
	if err := s.validatePutMatchObjectSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMatchObjectSize",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) PutOr(value *S3StorageLensGroupFilterOr) {
	if err := s.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOr",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		s,
		"resetAnd",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ResetMatchAnyPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnyPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ResetMatchAnySuffix() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnySuffix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ResetMatchAnyTag() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchAnyTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ResetMatchObjectAge() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchObjectAge",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ResetMatchObjectSize() {
	_jsii_.InvokeVoid(
		s,
		"resetMatchObjectSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		s,
		"resetOr",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3StorageLensGroupFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

