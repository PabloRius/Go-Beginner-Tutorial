## Initialize a module:
```bash
go mod init example.com/<module_name>
``` 
In order to execute some code, it must be coded inside the package name as "main" and execute inside it's terminal:
```bash
go run .
```
## Update dependencies
```bash
go mod tidy
```
This command won't work if there's any existing locally referenced module. (Greetings exemplifies this case)
To next command works around referencing local modules:
```bash
go mod edit -replace example.com/<module_name>=../<module_name>
```
In this case the locally referenced module is in the parent directory
After that, executing:
```bash
go mod tidy
```
Will update the dependency to the local module instead of looking it up on the online base.
## Testing
To execute a testing script you must use the command:
```bash
go test
```
Or if you want to use the verbose flag to view the output of every function:
```bash
go test -v
```
## Build and Install
To build an application:
```bash
go build
```
Inside the terminal of the module you want to build as your main application
You can also add the built program to the executable programs inside your machine by "installing" it
This commands will add the executable to the env variable PATH
First wee need to know where are the installed executables estored by Go by default:
```bash
go list -f '{{.Target}}'
```
The commando will show you the route where the executable will be stored once installed. We need the base directory for the next command:
```bash
set PATH=%PATH%;<bin_directory>
```
Where bin directory is the directory shown by the last command
Now executing:
```bash
go install
```
The previously generated executable file will be stored in the /bin directory, hence callable from anywhere in your system!