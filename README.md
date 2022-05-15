whoami
======

Simple HTTP docker service that prints it's container ID
    $ docker build -t whoami .

    $ docker rm -f whoami

    $ docker run -d -p 9000:8000 --name whoami -t whoami
    736ab83847bb12dddd8b09969433f3a02d64d5b0be48f7a5c59a594e3a6a3541
    
    $ curl $(hostname --all-ip-addresses | awk '{print $1}'):9000
    I'm 1d5da65f3705 time now: 2022-05-15 15:40:02.5668707 +0000 UTC m=+11.412986601


```
go mod init github.com/vimsydoodle/whoami
go mod tidy

skaffold init  --generate-manifests

Configuration skaffold.yaml was written
You can now run [skaffold build] to build the artifacts
or [skaffold run] to build and deploy
or [skaffold dev] to enter development mode, with auto-redeploy

skaffold.yaml? Yes
Generated manifest deployment.yaml was written
```