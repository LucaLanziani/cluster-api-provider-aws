/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// SetupWebhookWithManager will setup the webhooks for the EKSConfigTemplate.
func (r *EKSConfigTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-bootstrap-cluster-x-k8s-io-v1beta1-eksconfigtemplate,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=bootstrap.cluster.x-k8s.io,resources=eksconfigtemplate,versions=v1beta1,name=validation.eksconfigtemplates.bootstrap.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/mutate-bootstrap-cluster-x-k8s-io-v1beta1-eksconfigtemplate,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=bootstrap.cluster.x-k8s.io,resources=eksconfigtemplate,versions=v1beta1,name=default.eksconfigtemplates.bootstrap.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var _ webhook.Defaulter = &EKSConfigTemplate{}
var _ webhook.Validator = &EKSConfigTemplate{}

// ValidateCreate will do any extra validation when creating a EKSConfigTemplate.
func (r *EKSConfigTemplate) ValidateCreate() error {
	return nil
}

// ValidateUpdate will do any extra validation when updating a EKSConfigTemplate.
func (r *EKSConfigTemplate) ValidateUpdate(old runtime.Object) error {
	return nil
}

// ValidateDelete allows you to add any extra validation when deleting.
func (r *EKSConfigTemplate) ValidateDelete() error {
	return nil
}

// Default will set default values for the EKSConfigTemplate.
func (r *EKSConfigTemplate) Default() {
}
