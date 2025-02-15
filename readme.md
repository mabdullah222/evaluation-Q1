# Absolute Path Resolver

## Overview
This Go project converts a relative path into an absolute path using helper functions in `utils/utils.go`.

## Running the Code
### Prerequisites
Ensure Go is installed:
```sh
 go version
```

### Steps to Run
1. Run the program:
   ```sh
   go run main.go
   ```

## Tests
To verify the functionality, test with the following cases:
- **Redundant slashes:** `currDir = "/home///user//documents"`, `relPath = "../projects/code"` → Expected: `/home/user/projects/code`
- **Current directory references:** `currDir = "/home/user/./documents"`, `relPath = "./code"` → Expected: `/home/user/documents/code`
- **Multiple parent directories:** `currDir = "/home/user/documents"`, `relPath = "../../projects"` → Expected: `/home/projects`

## Author
- **Muhammad Abdullah**
- Email: [abdullah.bscs21seecs@seecs.edu.pk](mailto:abdullah.bscs21seecs@seecs.edu.pk)


