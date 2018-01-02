.PHONY: alli

TARGET=crawler

#disable debug info: -ldflags "-s -w"
gflags=-gcflags "-N -l"

all:
	go build -o ${TARGET} ${gflags}

clean:
	rm -rf ${TARGET}
