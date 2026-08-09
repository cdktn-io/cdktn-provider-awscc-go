// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretsmanagersecret

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/secretsmanagersecret/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsmanagerSecretGenerateSecretStringOutputReference interface {
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
	ExcludeCharacters() *string
	SetExcludeCharacters(val *string)
	ExcludeCharactersInput() *string
	ExcludeLowercase() interface{}
	SetExcludeLowercase(val interface{})
	ExcludeLowercaseInput() interface{}
	ExcludeNumbers() interface{}
	SetExcludeNumbers(val interface{})
	ExcludeNumbersInput() interface{}
	ExcludePunctuation() interface{}
	SetExcludePunctuation(val interface{})
	ExcludePunctuationInput() interface{}
	ExcludeUppercase() interface{}
	SetExcludeUppercase(val interface{})
	ExcludeUppercaseInput() interface{}
	// Experimental.
	Fqn() *string
	GenerateStringKey() *string
	SetGenerateStringKey(val *string)
	GenerateStringKeyInput() *string
	IncludeSpace() interface{}
	SetIncludeSpace(val interface{})
	IncludeSpaceInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PasswordLength() *float64
	SetPasswordLength(val *float64)
	PasswordLengthInput() *float64
	RequireEachIncludedType() interface{}
	SetRequireEachIncludedType(val interface{})
	RequireEachIncludedTypeInput() interface{}
	SecretStringTemplate() *string
	SetSecretStringTemplate(val *string)
	SecretStringTemplateInput() *string
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
	ResetExcludeCharacters()
	ResetExcludeLowercase()
	ResetExcludeNumbers()
	ResetExcludePunctuation()
	ResetExcludeUppercase()
	ResetGenerateStringKey()
	ResetIncludeSpace()
	ResetPasswordLength()
	ResetRequireEachIncludedType()
	ResetSecretStringTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecretsmanagerSecretGenerateSecretStringOutputReference
type jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeLowercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeLowercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeLowercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludePunctuation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludePunctuationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludePunctuationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeUppercase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ExcludeUppercaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeUppercaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GenerateStringKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generateStringKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GenerateStringKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generateStringKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) IncludeSpace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) IncludeSpaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) PasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) PasswordLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) RequireEachIncludedType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) RequireEachIncludedTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireEachIncludedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) SecretStringTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretStringTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) SecretStringTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretStringTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecretsmanagerSecretGenerateSecretStringOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecretsmanagerSecretGenerateSecretStringOutputReference {
	_init_.Initialize()

	if err := validateNewSecretsmanagerSecretGenerateSecretStringOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.secretsmanagerSecret.SecretsmanagerSecretGenerateSecretStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecretsmanagerSecretGenerateSecretStringOutputReference_Override(s SecretsmanagerSecretGenerateSecretStringOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.secretsmanagerSecret.SecretsmanagerSecretGenerateSecretStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetExcludeCharacters(val *string) {
	if err := j.validateSetExcludeCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCharacters",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetExcludeLowercase(val interface{}) {
	if err := j.validateSetExcludeLowercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeLowercase",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetExcludeNumbers(val interface{}) {
	if err := j.validateSetExcludeNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeNumbers",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetExcludePunctuation(val interface{}) {
	if err := j.validateSetExcludePunctuationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePunctuation",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetExcludeUppercase(val interface{}) {
	if err := j.validateSetExcludeUppercaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeUppercase",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetGenerateStringKey(val *string) {
	if err := j.validateSetGenerateStringKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generateStringKey",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetIncludeSpace(val interface{}) {
	if err := j.validateSetIncludeSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSpace",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetPasswordLength(val *float64) {
	if err := j.validateSetPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordLength",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetRequireEachIncludedType(val interface{}) {
	if err := j.validateSetRequireEachIncludedTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireEachIncludedType",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetSecretStringTemplate(val *string) {
	if err := j.validateSetSecretStringTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretStringTemplate",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetExcludeCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetExcludeLowercase() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeLowercase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetExcludeNumbers() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeNumbers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetExcludePunctuation() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludePunctuation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetExcludeUppercase() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeUppercase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetGenerateStringKey() {
	_jsii_.InvokeVoid(
		s,
		"resetGenerateStringKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetIncludeSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetPasswordLength() {
	_jsii_.InvokeVoid(
		s,
		"resetPasswordLength",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetRequireEachIncludedType() {
	_jsii_.InvokeVoid(
		s,
		"resetRequireEachIncludedType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ResetSecretStringTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretStringTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecretsmanagerSecretGenerateSecretStringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

