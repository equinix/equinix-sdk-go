module github.com/equinix/equinix-sdk-go/examples/services/securecabinetv1/check_availability

go 1.23

toolchain go1.24.4

require (
	github.com/equinix/equinix-sdk-go v0.54.2
	github.com/hashicorp/go-retryablehttp v0.7.8
)

require (
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	golang.org/x/oauth2 v0.26.0 // indirect
)

replace github.com/equinix/equinix-sdk-go => ../../../..
