# Compile and Install the application

1. ```go build```
2. On linux: ```./<name of the executable generated>```
3. `go list -f '{{.Target}}'`
4. `export PATH=$PATH:/path/to/your/install/directory`
5. `go install`
6. Now you can execute your application using the name of your application.