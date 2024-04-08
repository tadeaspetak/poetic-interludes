
.PHONY: build
build: css
	go build -o app

.PHONY: dev
dev:
	go get github.com/githubnemo/CompileDaemon && \
	go install github.com/githubnemo/CompileDaemon && \
	CompileDaemon -build="go build -o dev" -command="./dev"
# if necessary, uninstall with:
# rm -f (which CompileDaemon)
# go get github.com/githubnemo/CompileDaemon@none

.PHONY: css
css:
	npx tailwindcss -i ./static/main.css -o ./static/tailwind.css

.PHONY: css-watch
css-watch:
	npx tailwindcss --watch -i ./static/main.css -o ./static/tailwind.css