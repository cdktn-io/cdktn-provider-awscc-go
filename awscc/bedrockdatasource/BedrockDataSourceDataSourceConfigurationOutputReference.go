// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataSourceDataSourceConfigurationOutputReference interface {
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
	ConfluenceConfiguration() BedrockDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference
	ConfluenceConfigurationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManagedKnowledgeBaseConnectorConfiguration() BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference
	ManagedKnowledgeBaseConnectorConfigurationInput() interface{}
	S3Configuration() BedrockDataSourceDataSourceConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() interface{}
	SalesforceConfiguration() BedrockDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference
	SalesforceConfigurationInput() interface{}
	SharePointConfiguration() BedrockDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
	SharePointConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WebConfiguration() BedrockDataSourceDataSourceConfigurationWebConfigurationOutputReference
	WebConfigurationInput() interface{}
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
	PutConfluenceConfiguration(value *BedrockDataSourceDataSourceConfigurationConfluenceConfiguration)
	PutManagedKnowledgeBaseConnectorConfiguration(value *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfiguration)
	PutS3Configuration(value *BedrockDataSourceDataSourceConfigurationS3Configuration)
	PutSalesforceConfiguration(value *BedrockDataSourceDataSourceConfigurationSalesforceConfiguration)
	PutSharePointConfiguration(value *BedrockDataSourceDataSourceConfigurationSharePointConfiguration)
	PutWebConfiguration(value *BedrockDataSourceDataSourceConfigurationWebConfiguration)
	ResetConfluenceConfiguration()
	ResetManagedKnowledgeBaseConnectorConfiguration()
	ResetS3Configuration()
	ResetSalesforceConfiguration()
	ResetSharePointConfiguration()
	ResetWebConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataSourceDataSourceConfigurationOutputReference
type jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ConfluenceConfiguration() BedrockDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference
	_jsii_.Get(
		j,
		"confluenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ConfluenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ManagedKnowledgeBaseConnectorConfiguration() BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationOutputReference
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConnectorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ManagedKnowledgeBaseConnectorConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedKnowledgeBaseConnectorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) S3Configuration() BedrockDataSourceDataSourceConfigurationS3ConfigurationOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) SalesforceConfiguration() BedrockDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference
	_jsii_.Get(
		j,
		"salesforceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) SalesforceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) SharePointConfiguration() BedrockDataSourceDataSourceConfigurationSharePointConfigurationOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
	_jsii_.Get(
		j,
		"sharePointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) SharePointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharePointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) WebConfiguration() BedrockDataSourceDataSourceConfigurationWebConfigurationOutputReference {
	var returns BedrockDataSourceDataSourceConfigurationWebConfigurationOutputReference
	_jsii_.Get(
		j,
		"webConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) WebConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webConfigurationInput",
		&returns,
	)
	return returns
}


func NewBedrockDataSourceDataSourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataSourceDataSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataSourceDataSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceDataSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataSourceDataSourceConfigurationOutputReference_Override(b BedrockDataSourceDataSourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceDataSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) PutConfluenceConfiguration(value *BedrockDataSourceDataSourceConfigurationConfluenceConfiguration) {
	if err := b.validatePutConfluenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putConfluenceConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) PutManagedKnowledgeBaseConnectorConfiguration(value *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfiguration) {
	if err := b.validatePutManagedKnowledgeBaseConnectorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putManagedKnowledgeBaseConnectorConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) PutS3Configuration(value *BedrockDataSourceDataSourceConfigurationS3Configuration) {
	if err := b.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) PutSalesforceConfiguration(value *BedrockDataSourceDataSourceConfigurationSalesforceConfiguration) {
	if err := b.validatePutSalesforceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSalesforceConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) PutSharePointConfiguration(value *BedrockDataSourceDataSourceConfigurationSharePointConfiguration) {
	if err := b.validatePutSharePointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSharePointConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) PutWebConfiguration(value *BedrockDataSourceDataSourceConfigurationWebConfiguration) {
	if err := b.validatePutWebConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putWebConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ResetConfluenceConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetConfluenceConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ResetManagedKnowledgeBaseConnectorConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetManagedKnowledgeBaseConnectorConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		b,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ResetSalesforceConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetSalesforceConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ResetSharePointConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetSharePointConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ResetWebConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetWebConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceDataSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

