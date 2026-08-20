// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2connector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/inspectorv2connector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference interface {
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
	ContainerImageScanning() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationContainerImageScanningOutputReference
	ContainerImageScanningInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ServerlessScanning() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference
	ServerlessScanningInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmScanning() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference
	VmScanningInput() interface{}
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
	PutContainerImageScanning(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationContainerImageScanning)
	PutServerlessScanning(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanning)
	PutVmScanning(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanning)
	ResetContainerImageScanning()
	ResetServerlessScanning()
	ResetVmScanning()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference
type jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ContainerImageScanning() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationContainerImageScanningOutputReference {
	var returns Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationContainerImageScanningOutputReference
	_jsii_.Get(
		j,
		"containerImageScanning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ContainerImageScanningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerImageScanningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ServerlessScanning() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference {
	var returns Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanningOutputReference
	_jsii_.Get(
		j,
		"serverlessScanning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ServerlessScanningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessScanningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) VmScanning() Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference {
	var returns Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanningOutputReference
	_jsii_.Get(
		j,
		"vmScanning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) VmScanningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmScanningInput",
		&returns,
	)
	return returns
}


func NewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference_Override(i Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.inspectorv2Connector.Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) PutContainerImageScanning(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationContainerImageScanning) {
	if err := i.validatePutContainerImageScanningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putContainerImageScanning",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) PutServerlessScanning(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationServerlessScanning) {
	if err := i.validatePutServerlessScanningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putServerlessScanning",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) PutVmScanning(value *Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationVmScanning) {
	if err := i.validatePutVmScanningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVmScanning",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ResetContainerImageScanning() {
	_jsii_.InvokeVoid(
		i,
		"resetContainerImageScanning",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ResetServerlessScanning() {
	_jsii_.InvokeVoid(
		i,
		"resetServerlessScanning",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ResetVmScanning() {
	_jsii_.InvokeVoid(
		i,
		"resetVmScanning",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_Inspectorv2ConnectorProviderConfigurationAzureScopeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

