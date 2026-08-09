// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package imagebuilderimagerecipe

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImagebuilderImageRecipeComponentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageRecipeComponentsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageRecipeComponentsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageRecipeComponentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageRecipeComponentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageRecipeComponentsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageRecipeComponentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImagebuilderImageRecipeComponentsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

