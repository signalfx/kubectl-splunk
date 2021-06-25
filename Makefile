.PHONY: license
license:
	if [ -d third-party ]; then chmod -R u+w third-party; fi
	rm -rf third-party
	go-licenses save ./cmd/kubectl-splunk --save_path=third-party
	python3 notification.py

.PHONY: docs
docs:
	go run cmd/gen-docs/main.go

.PHONY: install
install:
	go install ./cmd/kubectl-splunk
