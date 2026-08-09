// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package glueblueprint

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueBlueprintTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlueBlueprintTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueBlueprintTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueBlueprintTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueBlueprintTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueBlueprintTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueBlueprintTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueBlueprintTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

