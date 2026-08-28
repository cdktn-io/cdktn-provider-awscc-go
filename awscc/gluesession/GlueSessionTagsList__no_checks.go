// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gluesession

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueSessionTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlueSessionTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueSessionTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueSessionTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueSessionTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueSessionTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueSessionTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueSessionTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

