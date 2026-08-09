// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccbedrockagentcoreoauth2credentialprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccbedrockagentcoreoauth2credentialprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference interface {
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
	InternalValue() *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpoint
	SetInternalValue(val *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpoint)
	ManagedVpcResource() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointManagedVpcResourceOutputReference
	SelfManagedLatticeResource() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResourceOutputReference
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

// The jsii proxy struct for DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference
type jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) InternalValue() *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpoint {
	var returns *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) ManagedVpcResource() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointManagedVpcResourceOutputReference {
	var returns DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointManagedVpcResourceOutputReference
	_jsii_.Get(
		j,
		"managedVpcResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) SelfManagedLatticeResource() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResourceOutputReference {
	var returns DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResourceOutputReference
	_jsii_.Get(
		j,
		"selfManagedLatticeResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockagentcoreOAuth2CredentialProvider.DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference_Override(d DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockagentcoreOAuth2CredentialProvider.DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference)SetInternalValue(val *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpoint) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesPrivateEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

