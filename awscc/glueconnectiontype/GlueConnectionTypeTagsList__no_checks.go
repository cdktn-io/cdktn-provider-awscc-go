// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package glueconnectiontype

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueConnectionTypeTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnectionTypeTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueConnectionTypeTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueConnectionTypeTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueConnectionTypeTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueConnectionTypeTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueConnectionTypeTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueConnectionTypeTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

