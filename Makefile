push:
	@go run utils/parser/problems/main.go
	@go run utils/plotter/main.go
	@go run utils/readme_builder/main.go
	@git add .
	@git commit -m  "$(MESSAGE)"
	@git push

