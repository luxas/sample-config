SHELL := /bin/bash
PROJECT = github.com/luxas/sample-config
APIS_DIR = ${APIS_DIR}

shell:
	docker run -it \
		-v $(shell pwd):/go/src/github.com/luxas/sample-config \
		-w /go/src/github.com/luxas/sample-config \
		-e GO111MODULE=on \
		golang:1.11

build: 
	docker run -it \
		-v $(shell pwd):/go/src/github.com/luxas/sample-config \
		-w /go/src/github.com/luxas/sample-config \
		-e GO111MODULE=on \
		golang:1.11 \
		make binary

binary: autogen vendor
	go build -o bin/sample-config github.com/luxas/sample-config/cmd/sample-config

vendor: binary
	if [[ ! -f go.mod ]]; then go mod init; fi
	go mod tidy
	go mod vendor
	go mod verify

autogen: /go/bin/deepcopy-gen /go/bin/defaulter-gen /go/bin/conversion-gen
	# Let the boilerplate be empty
	touch /tmp/boilerplate
	/go/bin/deepcopy-gen \
		--input-dirs ${APIS_DIR}/config,${APIS_DIR}/config/v1,${APIS_DIR}/config/v1beta1 \
		--bounding-dirs ${APIS_DIR} \
		-O zz_generated.deepcopy \
		-h /tmp/boilerplate

	/go/bin/defaulter-gen \
		--input-dirs ${APIS_DIR}/pkg/apis/config/v1,${APIS_DIR}/config/v1beta1 \
		-O zz_generated.defaults \
		-h /tmp/boilerplate

	/go/bin/conversion-gen \
		--input-dirs ${APIS_DIR}/config,${APIS_DIR}/config/v1,${APIS_DIR}/config/v1beta1 \
		-O zz_generated.conversion \
		-h /tmp/boilerplate

/go/bin/%:
	go install /go/src/github.com/luxas/sample-config/vendor/k8s.io/code-generator/cmd/$*

clean:
	rm -r bin vendor
	find . -type f | grep zz_generated | grep -v vendor | xargs -r rm
