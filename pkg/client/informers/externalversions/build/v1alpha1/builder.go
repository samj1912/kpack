/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	buildv1alpha1 "github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	versioned "github.com/pivotal/kpack/pkg/client/clientset/versioned"
	internalinterfaces "github.com/pivotal/kpack/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BuilderInformer provides access to a shared informer and lister for
// Builders.
type BuilderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BuilderLister
}

type builderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBuilderInformer constructs a new informer for Builder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBuilderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBuilderInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBuilderInformer constructs a new informer for Builder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBuilderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KpackV1alpha1().Builders(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KpackV1alpha1().Builders(namespace).Watch(context.TODO(), options)
			},
		},
		&buildv1alpha1.Builder{},
		resyncPeriod,
		indexers,
	)
}

func (f *builderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBuilderInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *builderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&buildv1alpha1.Builder{}, f.defaultInformer)
}

func (f *builderInformer) Lister() v1alpha1.BuilderLister {
	return v1alpha1.NewBuilderLister(f.Informer().GetIndexer())
}
