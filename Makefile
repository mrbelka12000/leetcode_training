push:
	@go run plotter/main.go
	@git add .
	@git commit -m  "$(MESSAGE)"
	@git push