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
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	mtgpupodv1 "lsalab.nthu/mtgpu/pkg/apis/mtgpupod/v1"
	versioned "lsalab.nthu/mtgpu/pkg/client/clientset/versioned"
	internalinterfaces "lsalab.nthu/mtgpu/pkg/client/informers/externalversions/internalinterfaces"
	v1 "lsalab.nthu/mtgpu/pkg/client/listers/mtgpupod/v1"
)

// MtgpuPodInformer provides access to a shared informer and lister for
// MtgpuPods.
type MtgpuPodInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MtgpuPodLister
}

type mtgpuPodInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMtgpuPodInformer constructs a new informer for MtgpuPod type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMtgpuPodInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMtgpuPodInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMtgpuPodInformer constructs a new informer for MtgpuPod type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMtgpuPodInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LsalabV1().MtgpuPods(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LsalabV1().MtgpuPods(namespace).Watch(options)
			},
		},
		&mtgpupodv1.MtgpuPod{},
		resyncPeriod,
		indexers,
	)
}

func (f *mtgpuPodInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMtgpuPodInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mtgpuPodInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mtgpupodv1.MtgpuPod{}, f.defaultInformer)
}

func (f *mtgpuPodInformer) Lister() v1.MtgpuPodLister {
	return v1.NewMtgpuPodLister(f.Informer().GetIndexer())
}