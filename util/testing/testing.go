package testing

import (
	core_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

// FakeK8sImplementer - fake implementer used for testing
type FakeK8sImplementer struct {
	NamespacesList   *v1.NamespaceList
	DeploymentSingle *v1beta1.Deployment
	DeploymentList   *v1beta1.DeploymentList

	// stores value of an updated deployment
	Updated *v1beta1.Deployment

	AvailableSecret *v1.Secret

	AvailablePods *v1.PodList

	// error to return
	Error error
}

// Namespaces - available namespaces
func (i *FakeK8sImplementer) Namespaces() (*v1.NamespaceList, error) {
	return i.NamespacesList, nil
}

// Deployment - available deployment, doesn't filter anything
func (i *FakeK8sImplementer) Deployment(namespace, name string) (*v1beta1.Deployment, error) {
	return i.DeploymentSingle, nil
}

// Deployments - available deployments
func (i *FakeK8sImplementer) Deployments(namespace string) (*v1beta1.DeploymentList, error) {
	return i.DeploymentList, nil
}

// Update - update deployment
func (i *FakeK8sImplementer) Update(deployment *v1beta1.Deployment) error {
	i.Updated = deployment
	return nil
}

// Secret - get secret
func (i *FakeK8sImplementer) Secret(namespace, name string) (*v1.Secret, error) {
	if i.Error != nil {
		return nil, i.Error
	}
	return i.AvailableSecret, nil
}

// Pods - available pods
func (i *FakeK8sImplementer) Pods(namespace, labelSelector string) (*v1.PodList, error) {
	return i.AvailablePods, nil
}

// ConfigMaps - returns nothing (not implemented)
func (i *FakeK8sImplementer) ConfigMaps(namespace string) core_v1.ConfigMapInterface {
	panic("not implemented")
	return nil
}
