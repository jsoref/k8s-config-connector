// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ratelimiter

import (
	"time"

	"golang.org/x/time/rate"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/ratelimiter"
)

func NewRateLimiter() ratelimiter.RateLimiter {
	// This is based on workqueue.DefaultControllerRateLimiter, but with different parameters better suited to KRM reconciliation.
	// Context is in b/188203307

	// We have both overall and per-item rate limiting.
	// The overall is a token bucket and the per-item is exponential
	// The per-item rate limiter initially retries much more slowly (2 seconds vs 2 milliseconds),
	// and a much faster ultimate limit (120 seconds instead of 1000 seconds).

	// If we implement b/190097904 we should revisit these values, in particular the max delay could
	// likely be much higher again.

	return workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(2*time.Second, 120*time.Second),
		// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)
}

func RequeueRateLimiter() ratelimiter.RateLimiter {
	return workqueue.NewMaxOfRateLimiter(
		// 5 qps, 50 bucket size.  This is the overall factor, and must be slower than the NewRateLimiter limit, to leave "room" for new items.
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(5), 50)},
	)
}
