# go_for_fun_projects
Create Go as executable:

go mod init go_for_fun_projects

go build

Execute with: "./go_for_fun_projects"

Testing: go test ./tests || go test ./tests -v (verbose) || go test ./tests -v -coverprofile cover.out || go tool cover -html=cover.out
