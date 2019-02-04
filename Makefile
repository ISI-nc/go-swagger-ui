.PHONY: all;

all: swagger-ui/dist/*
	packr -z
