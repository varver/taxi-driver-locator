# Taxi Driver Locator

## Stack used
Used MongoDB since we need faster reads and writes and mongo db can handle 50,000 reads and writes per minute with ease.

Used golang as it provide capability to run things concurrent and parallel manner by utilizing least resources.

Used Martini library of golang to provide rest api framework.


## Installation/Setup
### Note : folder structure is important here so please follow all the instructions provided here.

run these commands
~~~
mkdir -p ~/goprojects/src/gitlab.com/varver/
cd ~/goprojects/src/gitlab.com/varver/
sudo apt-get -y install git
git clone https://gitlab.com/varver/wmd.git
cd wmd/setup
~~~

Lets setup project with these commands , it will install golang and mongodb.
~~~
make setup
source ~/.goenv
make gotest
~~~

if you see "Hello, Go is working fine" then golang is successfully installed.

run these commands :
~~~
cd ~/goprojects/src/gitlab.com/varver/wmd/
go get -u -v github.com/tools/godep
go install github.com/tools/godep
godep restore
~~~

##### change configurations in config.ini file and run this command
~~~
go run server.go
~~~

##### The api will be available on localhost:3000

## How to run unit test ?
~~~
cd controllers
go test -v
cd ../utils
go test -v  
~~~


## Configuration modes :
in config.ini file EnvMode can have "dev" or "live" or "test" mode .

**1) "dev"** : log all the messages and errors in console.

**2) "live"** : log all messages and errors in syslog file.

**3) "test"** : activated only by test suite and uses "TestDatabase" config variable as db name for testing purpose.


## File/Folder structure
**Config folder** : contains package to load configurations.

**config.ini** : is our configuration file , open it to know about configurations available.

**controllers** : package which contain all the controllers linked with url routes in server.go file.

**server.go** : file contains main function and routes of api.

**db folder** : contains package that connects with the database.

**logger** : package to provide better logs. Provide logging as std out in "dev" mode and in syslog in "live" mode.

**models** : package contains struct for driver's document that will be stored in database.

**utils** : this package contains utilities and helper functions.

**Godep** : contains all the dependencies of the project .

**vendor** : contains files linked with Godep.
