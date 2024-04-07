go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon

CompileDaemon -build="go build -o dev" -command="./dev"

# if necessary, uninstall with:
# rm -f (which CompileDaemon)
# go get github.com/githubnemo/CompileDaemon@none