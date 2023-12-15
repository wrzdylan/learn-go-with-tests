# Init Go Project

## Commande

- `go mod init <modulepath>`
Permet de spécifier la version de go et les dépendances pour un projet.     

- Linter Go
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2

golangci-lint --version
```

- `go fmt` pour formatter le code       

- `go run {filename.go}`       

- `go test`   

- `go install golang.org/x/tools/cmd/godoc@latest` puis `godoc -http :8000`     
Permet de voir tout les packages installer sur notre système.         

## Test

- Le nom du fichier doit avoir le format : `xxx_test.go`       
- Le nom de la fonction test doit commencer par : `Test`        
- La fonction test prend seulement 1 argument : `t *testing.T`      