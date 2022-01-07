ARTIFACT_DIR := ./artifacts

###############################################################################
###                               Build                                     ###
###############################################################################
install: go.sum
	@echo "--> Installing realestated"
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/realestated

build: go.sum
	@echo "--> Building realestate chain"
	@go build -o $(ARTIFACT_DIR)/realestated -mod=readonly ./cmd/realestated