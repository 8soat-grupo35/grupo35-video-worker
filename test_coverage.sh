 go test $(go list ./internal/... | grep -v /mock | grep -v /wrapper) -coverprofile=profile.cov
go tool cover -func=profile.cov