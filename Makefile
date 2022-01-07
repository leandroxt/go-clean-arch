## run: run the cmd/api application
.PHONY: run
run:
	@ go run ./cmd/api -giphy-api-key=${GIPHY_API_KEY}
