module github.com/MabsIPCA/helm-tests/helm_fetcher

go 1.25.3

require (
	github.com/MabsIPCA/helm-tests/helmfix v0.0.0
	github.com/joho/godotenv v1.5.1
	github.com/rs/zerolog v1.34.0
	sigs.k8s.io/yaml v1.6.0
)

replace github.com/MabsIPCA/helm-tests/helmfix v0.0.0 => ../helmfix

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	go.yaml.in/yaml/v2 v2.4.2 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
