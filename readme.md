Create Go as executable:

1. go mod init go_for_fun_projects
2. go build
3. Execute with: "./go_for_fun_projects"

4. Testing: go test ./tests     || go test ./tests -v (verbose)      || go test ./tests -v -coverprofile cover.out || go tool cover -html=cover.out