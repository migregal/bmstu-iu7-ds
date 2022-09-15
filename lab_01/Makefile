.PHONY: build_app
build_app:
	go build -o app.out ./cmd/app/main.go

.PHONY: build_installer
build_installer:
	go build -o install.out cmd/installer/main.go

.PHONY: install
install: build_app build_installer
	./install.out

.PHONY: clean
clean:
	rm -rf *.out
