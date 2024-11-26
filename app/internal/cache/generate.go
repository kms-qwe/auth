package cache

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i UserCache -o ./mocks/ -s "_minimock.go"
