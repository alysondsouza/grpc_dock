# grpc_dock

TAK NADJA!!!
============

TUTORIAL 

if you haven't installed google's protocol buffers, see the prerequisites part at the bottom.

Setup of repository
make go.mod file with:

$ go mod init [link to repo without "https://"]

your repo should be on the public github. i couldn't get it to work on the itu instance.

make a .proto file in a subfolder, for example time/time.proto and fill it with IDL.

notice line 3 and 4.

package time;
option go_package = "https://github.itu.dk/nako/DISYS-week5;time";
the go_package should equal a link to your repo, a semicolon and then the package name.

run command:

$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative time/time.proto

which should create the two pb.go files.

run command:

$ go mod tidy

to install dependencies and create the go.sum file.

create a client\client.go file with the client_template.txt as template.

create a server\server.go file with the server_template.txt as template.

switch out the "myPackage" with your actual package.

switch our the method names with actual method names.

add more methods to the client.go file, so that there's a method for each request in the .proto file.

when everything is compilable, open a terminal, change directory to the server folder, and run the command:

$ go run .

this will start your server.

open a new terminal, change directory to the client folder and run the command:

$ go run .

this will run the requests listed in the main method of the client file.

create a Dockerfile like the one in this repository.

change line 11, 12 and 16 to match your repository.

from now on

please remember to commit and push changes to the files in your repository before running the program.

the following docker commands will clone your repository (maybe to the virtual machine?), so changes to files will not be applied, if you don't git commit yeet before.

the only exception (i think) is changes to the client.go file, since it's run locally on your computer, but just connects to the server in docker.

run command:

docker build -t test .

to build the code.

run command:

docker run -p 9080:9080 -tid test

to run the code in a docker container.

change directory into your client folder.

run command:

go run .

to run the code.

Prerequisites
before starting, install google's protocol buffers:

go to this link: https://developers.google.com/protocol-buffers/docs/downloads
click on the "release page" link.
find the version you need and download it.
as per october 2021, if your on windows, it's the third from the bottom, protoc-3.18.1-win64.zip.
unzip the downloaded file somewhere "safe".

on my windows machine, i placed it in C:\Program Files\Protoc
add the path to the bin folder to your system variables.

on windows, click the windows key and search for "system", then there should come something up like "edit the system environment variables".

click the button "environment variables..." at the bottom.

in the bottom list select the variable called "path" and click "edit ..."

in the pop-up window, click "new..."

paste the total path to the bin folder into the text field.

my path is C:\Program Files\Protoc\bin.

click "ok".

open a terminal and run these commands:

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
