// This file was automatically generated by informer-gen

package v1

import (
	authorization_v1 "github.com/openshift/api/authorization/v1"
	versioned "github.com/openshift/client-go/authorization/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/authorization/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/authorization/listers/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterPolicyInformer provides access to a shared informer and lister for
// ClusterPolicies.
type ClusterPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterPolicyLister
}

type clusterPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterPolicyInformer constructs a new informer for ClusterPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterPolicyInformer constructs a new informer for ClusterPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthorizationV1().ClusterPolicies().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthorizationV1().ClusterPolicies().Watch(options)
			},
		},
		&authorization_v1.ClusterPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authorization_v1.ClusterPolicy{}, f.defaultInformer)
}

func (f *clusterPolicyInformer) Lister() v1.ClusterPolicyLister {
	return v1.NewClusterPolicyLister(f.Informer().GetIndexer())
}
