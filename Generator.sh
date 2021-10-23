#!/bin/bash

protoc ./Greet/GreetPb/Greet.proto --go_out=plugins=grpc:.