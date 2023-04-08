package rules

var validComputePlatforms = map[string]bool {
	"standard-v1": true,
	"standard-v2": true,
	"standard-v3": true,
	"gpu-standard-v1": true,
	"gpu-standard-v2": true,
	"standard-v3-t4": true,
	"gpu-standard-v3": true,
}

var validAvailabilityZones = map[string]bool {
	"ru-central1-a": true,
	"ru-central1-b": true,
	"ru-central1-c": true,
}
