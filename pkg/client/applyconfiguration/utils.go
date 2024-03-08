// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/jovik31/controller/pkg/apis/jovik31.dev/v1alpha1"
	jovik31devv1alpha1 "github.com/jovik31/controller/pkg/client/applyconfiguration/jovik31.dev/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=jovik31.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("Node"):
		return &jovik31devv1alpha1.NodeApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Tenant"):
		return &jovik31devv1alpha1.TenantApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TenantSpec"):
		return &jovik31devv1alpha1.TenantSpecApplyConfiguration{}

	}
	return nil
}
