package rules

var ValidComputePlatforms = map[string]bool{
	"standard-v1":     true,
	"standard-v2":     true,
	"standard-v3":     true,
	"gpu-standard-v1": true,
	"gpu-standard-v2": true,
	"standard-v3-t4":  true,
	"gpu-standard-v3": true,
}

var ValidAvailabilityZones = map[string]bool{
	"ru-central1-a": true,
	"ru-central1-b": true,
	"ru-central1-c": true,
}

var ValidChallengeType = map[string]bool{
	"DNS_CNAME": true,
	"DNS_TXT":   true,
	"HTTP":      true,
}

var ValidDiskTypes = map[string]bool{
	"network-ssd":               true,
	"network-hdd":               true,
	"network-ssd-nonreplicated": true,
}

var ValidGpuInterconnectType = map[string]bool{
	"infiniband": true,
}

var ValidImageOsTypes = map[string]bool{
	"LINUX":   true,
	"WINDOWS": true,
}

var ValidNetworkAccelecrationTypes = map[string]bool{
	"standard":             true,
	"software_accelerated": true,
}

var ValidDiskAccessMode = map[string]bool{
	"READ_WRITE": true,
	"READ_ONLY":  true,
}

// https://cloud.yandex.com/en-ru/docs/container-registry/security/#service-roles
var ValidContainerRegistryServiceRoles = map[string]bool{
	"container-registry.admin":          true,
	"container-registry.images.puller":  true,
	"container-registry.images.pusher":  true,
	"resource-manager.clouds.member":    true,
	"resource-manager.clouds.owner":     true,
	"container-registry.viewer":         true,
	"container-registry.editor":         true,
	"container-registry.images.scanner": true,
}

//https://cloud.yandex.com/en/docs/dns/concepts/resource-record
var ValidDnsRecordTypes = map[string]bool{
	"A":     true,
	"AAAA":  true,
	"CAA":   true,
	"CNAME": true,
	"ANAME": true,
	"MX":    true,
	"NS":    true,
	"PTR":   true,
	"SOA":   true,
	"SRV":   true,
	"TXT":   true,
}

//https://cloud.yandex.com/en-ru/docs/functions/security/#roles
var ValidFunctionServiceRoles = map[string]bool{
	"functions.viewer":               true,
	"functions.auditor":              true,
	"functions.functionInvoker":      true,
	"functions.editor":               true,
	"functions.mdbProxiesUser":       true,
	"functions.admin":                true,
	"resource-manager.auditor":       true,
	"resource-manager.viewer":        true,
	"resource-manager.editor":        true,
	"resource-manager.admin":         true,
	"resource-manager.clouds.member": true,
}
