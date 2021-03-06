package openshift

import (
	"github.com/RHsyseng/operator-utils/internal/platform"
	"k8s.io/client-go/rest"
)

/*
GetPlatformInfo examines the Kubernetes-based environment and determines the running platform, version, & OS.
Accepts <nil> or instantiated 'cfg' rest config parameter.

Result: PlatformInfo{ Name: OpenShift, K8SVersion: 1.13+, OS: linux/amd64 }
*/
func GetPlatformInfo(cfg *rest.Config) (platform.PlatformInfo, error) {
	return platform.K8SBasedPlatformVersioner{}.GetPlatformInfo(nil, cfg)
}

/*
IsOpenShift is a helper method to simplify boolean OCP checks against GetPlatformInfo results
Accepts <nil> or instantiated 'cfg' rest config parameter.
*/
func IsOpenShift(cfg *rest.Config) (bool, error) {
	info, err := GetPlatformInfo(cfg)
	if err != nil {
		return false, err
	}
	return info.IsOpenShift(), nil
}

/*
LookupOpenShiftVersion fetches OpenShift version info from API endpoints
*** NOTE: OCP 4.1+ requires elevated user permissions, see PlatformVersioner for details
Accepts <nil> or instantiated 'cfg' rest config parameter.

Result: OpenShiftVersion{ Version: 4.1.2 }
*/
func LookupOpenShiftVersion(cfg *rest.Config) (platform.OpenShiftVersion, error) {
	return platform.K8SBasedPlatformVersioner{}.LookupOpenShiftVersion(nil, cfg)
}

/*
MapKnownVersion maps from K8S version of PlatformInfo to equivalent OpenShift version

Result: OpenShiftVersion{ Version: 4.1.2 }
*/
func MapKnownVersion(info platform.PlatformInfo) platform.OpenShiftVersion {
	k8sToOcpMap := map[string]string{
		"1.10+": "3.10",
		"1.11+": "3.11",
		"1.13+": "4.1",
	}
	return platform.OpenShiftVersion{Version: k8sToOcpMap[info.K8SVersion]}
}
