// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package socket

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net v0.7.0x/net/publicsuffix"
)

type ptrRecord struct {
	hostname string
	error    error
	expires  time.Time
}

func (r ptrRecord) IsExpired(now time.Time) bool {
	return now.After(r.expires)
}

// ReverseLookupCache is a cache for storing and retrieving the results of
// reverse DNS queries. It caches the results of queries regardless of their
// outcome (success or failure). The result is cached for the amount of time
// specified by parameters and not based on the TTL from the PTR record.
type ReverseLookupCache struct {
	data                   map[string]ptrRecord
	successTTL, failureTTL time.Duration
}

// NewReverseLookupCache returns a new cache.
func NewReverseLookupCache(successTTL, failureTTL time.Duration) *ReverseLookupCache {
	c := &ReverseLookupCache{
		data:       map[string]ptrRecord{},
		successTTL: successTTL,
		failureTTL: failureTTL,
	}

	return c
}

// Cleanup removes expired entries from the cache.
func (c *ReverseLookupCache) Cleanup() {
	now := time.Now()
	for k, ptr := range c.data {
		if ptr.IsExpired(now) {
			delete(c.data, k)
		}
	}
}

// Lookup performs a reverse lookup on the given IP address. A cached result
// will be returned if it is contained in the cache, otherwise a lookup is
// performed.
func (c ReverseLookupCache) Lookup(ip net.IP) (string, error) {
	// Go doesn't expose a lookup method that accepts net.IP so
	// unfortunately we must convert the IP to a string.
	ipStr := ip.String()

	// XXX: This could be implemented using common.Cache with a separate
	// cleanup thread.
	c.Cleanup()

	// Check the cache.
	now := time.Now()
	if ptr, found := c.data[ipStr]; found && !ptr.IsExpired(now) {
		return ptr.hostname, ptr.error
	}

	// Do a new lookup.
	names, err := net.LookupAddr(ipStr)
	now = time.Now()

	var ptr ptrRecord
	switch {
	case err != nil:
		ptr.expires = now.Add(c.failureTTL)
		ptr.error = err
	case len(names) == 0:
		ptr.expires = now.Add(c.failureTTL)
		ptr.error = fmt.Errorf("empty dns response")
	default:
		ptr.expires = now.Add(c.successTTL)
		ptr.hostname = names[0]
	}

	c.data[ipStr] = ptr
	return ptr.hostname, ptr.error
}

// etldPlusOne returns the effective top-level domain plus one domain for the
// given hostname.
func etldPlusOne(hostname string) (string, error) {
	if hostname == "" {
		return "", nil
	}

	trimmed := false
	if hostname[len(hostname)-1] == '.' {
		hostname = hostname[:len(hostname)-1]
		trimmed = true
	}

	domain, err := publicsuffix.EffectiveTLDPlusOne(hostname)
	if err != nil {
		return "", err
	}

	if trimmed {
		return domain + ".", nil
	}
	return domain, nil
}
