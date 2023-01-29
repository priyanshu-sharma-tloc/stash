/*
Copyright 2019 The Stash Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	repositoriesv1alpha1 "github.com/appscode/stash/client/clientset/versioned/typed/repositories/v1alpha1"
	stashv1alpha1 "github.com/appscode/stash/client/clientset/versioned/typed/stash/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	RepositoriesV1alpha1() repositoriesv1alpha1.RepositoriesV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Repositories() repositoriesv1alpha1.RepositoriesV1alpha1Interface
	StashV1alpha1() stashv1alpha1.StashV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Stash() stashv1alpha1.StashV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	repositoriesV1alpha1 *repositoriesv1alpha1.RepositoriesV1alpha1Client
	stashV1alpha1        *stashv1alpha1.StashV1alpha1Client
}

// RepositoriesV1alpha1 retrieves the RepositoriesV1alpha1Client
func (c *Clientset) RepositoriesV1alpha1() repositoriesv1alpha1.RepositoriesV1alpha1Interface {
	return c.repositoriesV1alpha1
}

// Deprecated: Repositories retrieves the default version of RepositoriesClient.
// Please explicitly pick a version.
func (c *Clientset) Repositories() repositoriesv1alpha1.RepositoriesV1alpha1Interface {
	return c.repositoriesV1alpha1
}

// StashV1alpha1 retrieves the StashV1alpha1Client
func (c *Clientset) StashV1alpha1() stashv1alpha1.StashV1alpha1Interface {
	return c.stashV1alpha1
}

// Deprecated: Stash retrieves the default version of StashClient.
// Please explicitly pick a version.
func (c *Clientset) Stash() stashv1alpha1.StashV1alpha1Interface {
	return c.stashV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.repositoriesV1alpha1, err = repositoriesv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.stashV1alpha1, err = stashv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.repositoriesV1alpha1 = repositoriesv1alpha1.NewForConfigOrDie(c)
	cs.stashV1alpha1 = stashv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.repositoriesV1alpha1 = repositoriesv1alpha1.New(c)
	cs.stashV1alpha1 = stashv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
