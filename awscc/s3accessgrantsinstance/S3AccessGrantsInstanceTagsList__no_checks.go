// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3accessgrantsinstance

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3AccessGrantsInstanceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessGrantsInstanceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3AccessGrantsInstanceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantsInstanceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantsInstanceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantsInstanceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantsInstanceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3AccessGrantsInstanceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

