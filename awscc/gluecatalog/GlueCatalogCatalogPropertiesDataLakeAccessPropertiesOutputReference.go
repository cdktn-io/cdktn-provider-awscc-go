// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluecatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference interface {
	cdktn.ComplexObject
	AllowFullTableExternalDataAccess() *string
	SetAllowFullTableExternalDataAccess(val *string)
	AllowFullTableExternalDataAccessInput() *string
	CatalogType() *string
	SetCatalogType(val *string)
	CatalogTypeInput() *string
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
	DataLakeAccess() interface{}
	SetDataLakeAccess(val interface{})
	DataLakeAccessInput() interface{}
	DataTransferRole() *string
	SetDataTransferRole(val *string)
	DataTransferRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	ManagedWorkgroupName() *string
	ManagedWorkgroupStatus() *string
	RedshiftDatabaseName() *string
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
	ResetAllowFullTableExternalDataAccess()
	ResetCatalogType()
	ResetDataLakeAccess()
	ResetDataTransferRole()
	ResetKmsKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference
type jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) AllowFullTableExternalDataAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) AllowFullTableExternalDataAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) CatalogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) CatalogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) DataLakeAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLakeAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) DataLakeAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLakeAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) DataTransferRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) DataTransferRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTransferRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ManagedWorkgroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedWorkgroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ManagedWorkgroupStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedWorkgroupStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) RedshiftDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueCatalog.GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference_Override(g GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueCatalog.GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetAllowFullTableExternalDataAccess(val *string) {
	if err := j.validateSetAllowFullTableExternalDataAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFullTableExternalDataAccess",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetCatalogType(val *string) {
	if err := j.validateSetCatalogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogType",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetDataLakeAccess(val interface{}) {
	if err := j.validateSetDataLakeAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLakeAccess",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetDataTransferRole(val *string) {
	if err := j.validateSetDataTransferRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTransferRole",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ResetAllowFullTableExternalDataAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowFullTableExternalDataAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ResetCatalogType() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ResetDataLakeAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetDataLakeAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ResetDataTransferRole() {
	_jsii_.InvokeVoid(
		g,
		"resetDataTransferRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ResetKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogCatalogPropertiesDataLakeAccessPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

