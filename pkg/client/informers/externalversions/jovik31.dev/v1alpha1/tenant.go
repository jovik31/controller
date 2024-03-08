// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	jovik31devv1alpha1 "github.com/jovik31/controller/pkg/apis/jovik31.dev/v1alpha1"
	versioned "github.com/jovik31/controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jovik31/controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/jovik31/controller/pkg/client/listers/jovik31.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TenantInformer provides access to a shared informer and lister for
// Tenants.
type TenantInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TenantLister
}

type tenantInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTenantInformer constructs a new informer for Tenant type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTenantInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTenantInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTenantInformer constructs a new informer for Tenant type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTenantInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Jovik31V1alpha1().Tenants(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Jovik31V1alpha1().Tenants(namespace).Watch(context.TODO(), options)
			},
		},
		&jovik31devv1alpha1.Tenant{},
		resyncPeriod,
		indexers,
	)
}

func (f *tenantInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTenantInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tenantInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jovik31devv1alpha1.Tenant{}, f.defaultInformer)
}

func (f *tenantInformer) Lister() v1alpha1.TenantLister {
	return v1alpha1.NewTenantLister(f.Informer().GetIndexer())
}
