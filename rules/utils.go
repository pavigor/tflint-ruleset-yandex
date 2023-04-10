package rules

var ValidComputePlatforms = map[string]bool {
	"standard-v1": true,
	"standard-v2": true,
	"standard-v3": true,
	"gpu-standard-v1": true,
	"gpu-standard-v2": true,
	"standard-v3-t4": true,
	"gpu-standard-v3": true,
}

var ValidAvailabilityZones = map[string]bool {
	"ru-central1-a": true,
	"ru-central1-b": true,
	"ru-central1-c": true,
}

var ValidChallengeType = map[string]bool {
	"DNS_CNAME": true,
	"DNS_TXT": true,
	"HTTP": true,
}

var ValidDiskTypes = map[string]bool {
	"network-ssd": true,
	"network-hdd": true,
	"network-ssd-nonreplicated": true,
}