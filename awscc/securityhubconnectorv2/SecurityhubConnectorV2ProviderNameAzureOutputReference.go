// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/securityhubconnectorv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubConnectorV2ProviderNameAzureOutputReference interface {
	cdktn.ComplexObject
	AwsConfigConnectorArn() *string
	SetAwsConfigConnectorArn(val *string)
	AwsConfigConnectorArnInput() *string
	AzureRegions() *[]*string
	SetAzureRegions(val *[]*string)
	AzureRegionsInput() *[]*string
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
	ScopeConfiguration() SecurityhubConnectorV2ProviderNameAzureScopeConfigurationOutputReference
	ScopeConfigurationInput() interface{}
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
	PutScopeConfiguration(value *SecurityhubConnectorV2ProviderNameAzureScopeConfiguration)
	ResetAwsConfigConnectorArn()
	ResetAzureRegions()
	ResetScopeConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubConnectorV2ProviderNameAzureOutputReference
type jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) AwsConfigConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsConfigConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) AwsConfigConnectorArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsConfigConnectorArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) AzureRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"azureRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) AzureRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"azureRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ScopeConfiguration() SecurityhubConnectorV2ProviderNameAzureScopeConfigurationOutputReference {
	var returns SecurityhubConnectorV2ProviderNameAzureScopeConfigurationOutputReference
	_jsii_.Get(
		j,
		"scopeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ScopeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityhubConnectorV2ProviderNameAzureOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecurityhubConnectorV2ProviderNameAzureOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubConnectorV2ProviderNameAzureOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.securityhubConnectorV2.SecurityhubConnectorV2ProviderNameAzureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityhubConnectorV2ProviderNameAzureOutputReference_Override(s SecurityhubConnectorV2ProviderNameAzureOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.securityhubConnectorV2.SecurityhubConnectorV2ProviderNameAzureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference)SetAwsConfigConnectorArn(val *string) {
	if err := j.validateSetAwsConfigConnectorArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsConfigConnectorArn",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference)SetAzureRegions(val *[]*string) {
	if err := j.validateSetAzureRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureRegions",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) PutScopeConfiguration(value *SecurityhubConnectorV2ProviderNameAzureScopeConfiguration) {
	if err := s.validatePutScopeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putScopeConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ResetAwsConfigConnectorArn() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsConfigConnectorArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ResetAzureRegions() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureRegions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ResetScopeConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetScopeConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityhubConnectorV2ProviderNameAzureOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

