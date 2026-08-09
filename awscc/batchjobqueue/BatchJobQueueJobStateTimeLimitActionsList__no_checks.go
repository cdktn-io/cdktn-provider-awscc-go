// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchjobqueue

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchJobQueueJobStateTimeLimitActionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueJobStateTimeLimitActionsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueJobStateTimeLimitActionsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchJobQueueJobStateTimeLimitActionsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

