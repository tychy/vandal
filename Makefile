
# run go run main.go for TARGET
run: 
	go run main.go -path=$(TARGET)


test:
	@for i in {1..774}; do \
		echo "sample$$i"; \
		go run main.go -path=sample$$i; \
	done