/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	naeem4265comv1 "github.com/naeem4265/custom-controller/pkg/apis/naeem4265.com/v1"
	versioned "github.com/naeem4265/custom-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/naeem4265/custom-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/naeem4265/custom-controller/pkg/client/listers/naeem4265.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BookInformer provides access to a shared informer and lister for
// Books.
type BookInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BookLister
}

type bookInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBookInformer constructs a new informer for Book type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBookInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBookInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBookInformer constructs a new informer for Book type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBookInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Naeem4265V1().Books(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Naeem4265V1().Books(namespace).Watch(context.TODO(), options)
			},
		},
		&naeem4265comv1.Book{},
		resyncPeriod,
		indexers,
	)
}

func (f *bookInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBookInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bookInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&naeem4265comv1.Book{}, f.defaultInformer)
}

func (f *bookInformer) Lister() v1.BookLister {
	return v1.NewBookLister(f.Informer().GetIndexer())
}
