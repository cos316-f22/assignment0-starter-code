CXX ?= g++

all: cpp_bin
	@(echo "select 42" | \
		sqlite3 && \
		./cpp_bin && \
		curl -s https://cos316.princeton.edu/assignment0-blob.txt) | \
		go run cos316.princeton.edu/assignment0/summarize | \
		tee answer

cpp_bin: cpp_bin.cpp
	$(CXX) -o $@ $^
