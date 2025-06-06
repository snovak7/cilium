// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package builder

import (
	"github.com/cilium/cilium/cilium-cli/connectivity/check"
	"github.com/cilium/cilium/cilium-cli/connectivity/tests"
	"github.com/cilium/cilium/cilium-cli/utils/features"
	"github.com/cilium/cilium/pkg/versioncheck"
)

type bgpControlPlane struct{}

func (t bgpControlPlane) build(ct *check.ConnectivityTest, _ map[string]string) {
	// prefix the test name with `seq-` to run it sequentially
	newTest("seq-bgp-control-plane-v1", ct).
		WithCondition(func() bool {
			return versioncheck.MustCompile(">=1.16.0")(ct.CiliumVersion) && ct.Params().IncludeUnsafeTests
		}).
		WithFeatureRequirements(
			features.RequireEnabled(features.BGPControlPlane),
			features.RequireEnabled(features.NodeWithoutCilium),
		).
		WithScenarios(tests.BGPAdvertisements(1))

	// prefix the test name with `seq-` to run it sequentially
	newTest("seq-bgp-control-plane-v2", ct).
		WithCondition(func() bool {
			return versioncheck.MustCompile(">=1.16.0")(ct.CiliumVersion) && ct.Params().IncludeUnsafeTests
		}).
		WithFeatureRequirements(
			features.RequireEnabled(features.BGPControlPlane),
			features.RequireEnabled(features.NodeWithoutCilium),
		).
		WithScenarios(tests.BGPAdvertisements(2))
}
