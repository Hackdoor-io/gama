PORT=8080
MAIN=main.go

run:
	ENV=production go run $(MAIN)

dev:
	ENV=development PORT=$(PORT) go run $(MAIN)
