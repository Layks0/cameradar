sudo apt-get update && apt-get upgrade
sudo apt-get install golang
sudo apt-get install libcur14-openss1-dev pkg
go get github.com/Ullaakut/cameradar
cd $GOPATH/src/github.com/Ullaakut/cameradar
cd cmd/cameradar
go install
cameradar -t <109.237.2.66>

