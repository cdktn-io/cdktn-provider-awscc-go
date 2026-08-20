// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/inspectorv2connector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2ConnectorProviderConfigurationAzureOutputReference interface {
	cdktn.ComplexObject
	AutoInstallVmScanner() interface{}
	SetAutoInstallVmScanner(val interface{})
	AutoInstallVmScannerInput() interface{}
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
	ScopeConfiguration() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference
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
	PutScopeConfiguration(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfiguration)
	ResetAutoInstallVmScanner()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspectorv2ConnectorProviderConfigurationAzureOutputReference
type jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) AutoInstallVmScanner() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoInstallVmScanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) AutoInstallVmScannerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoInstallVmScannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) AwsConfigConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsConfigConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) AwsConfigConnectorArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsConfigConnectorArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) AzureRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"azureRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) AzureRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"azureRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) ScopeConfiguration() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference {
	var returns Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference
	_jsii_.Get(
		j,
		"scopeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) ScopeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInspectorv2ConnectorProviderConfigurationAzureOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Inspectorv2ConnectorProviderConfigurationAzureOutputReference {
	_init_.Initialize()

	if err := validateNewInspectorv2ConnectorProviderConfigurationAzureOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInspectorv2ConnectorProviderConfigurationAzureOutputReference_Override(i Inspectorv2ConnectorProviderConfigurationAzureOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetAutoInstallVmScanner(val interface{}) {
	if err := j.validateSetAutoInstallVmScannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoInstallVmScanner",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetAwsConfigConnectorArn(val *string) {
	if err := j.validateSetAwsConfigConnectorArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsConfigConnectorArn",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetAzureRegions(val *[]*string) {
	if err := j.validateSetAzureRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureRegions",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) PutScopeConfiguration(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfiguration) {
	if err := i.validatePutScopeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putScopeConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) ResetAutoInstallVmScanner() {
	_jsii_.InvokeVoid(
		i,
		"resetAutoInstallVmScanner",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

