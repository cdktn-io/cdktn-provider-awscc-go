// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package panoramapackage

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PanoramaPackageTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PanoramaPackageTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PanoramaPackageTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PanoramaPackageTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PanoramaPackageTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PanoramaPackageTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PanoramaPackageTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPanoramaPackageTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

