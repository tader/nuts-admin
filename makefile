.PHONY: dev

run-generators: gen-api

install-tools:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0

gen-api:
	oapi-codegen -generate server,types -package api api/api.yaml > api/generated.go

	oapi-codegen -generate client,types -package client -exclude-schemas VerifiableCredential,VerifiablePresentation,DID,DIDDocument -generate types,skip-prune -o nuts/client/generated.go https://nuts-node.readthedocs.io/en/latest/_static/common/ssi_types.yaml
	oapi-codegen -generate client,types -package vdr -import-mapping='../common/ssi_types.yaml:github.com/nuts-foundation/nuts-admin/nuts/client' -o nuts/client/vdr/generated.go https://nuts-node.readthedocs.io/en/latest/_static/vdr/v2.yaml
	oapi-codegen -generate client,types -package vcr -import-mapping='../common/ssi_types.yaml:github.com/nuts-foundation/nuts-admin/nuts/client' -o nuts/client/vcr/generated.go https://nuts-node.readthedocs.io/en/latest/_static/vcr/vcr_v2.yaml
	oapi-codegen -generate client,types -package discovery -import-mapping='../common/ssi_types.yaml:github.com/nuts-foundation/nuts-admin/nuts/client' -o nuts/client/discovery/generated.go https://nuts-node.readthedocs.io/en/latest/_static/discovery/v1.yaml

dev:
	make -j3 watch run-nuts-node run-api

watch:
	npm install
	npm run watch

run-nuts-node:
	docker compose pull
	docker compose up --wait

run-api:
	go run . live --configfile=deploy/admin.config.yaml
