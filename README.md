# OpenTag Task

## Basic Overview
Task regarding OpenTag's Golang application. The program itself is a server, which
helps translate english words and sentences to gopher ones. The application uses
go modules.

## Endpoints
The server has 3 endpoints. One for word translation, one for translating whole sentences and one
responsible for returning history data meaning all the translated words and sentences in the current server
run.

### /word endpoint
This is a POST request, requires a JSON body with single key value pair in the given format:
{"english_word": "word provided"}. The response will be the word, translated in gopher.

### /sentence endpoint
This endpoint is again a POST one.A JSON body with single key value pair in the given format is required:
{"english_sentence": "sentence provided"}. The response will be the sentence, translated in gopher.

### /history
The history endpoint will provide the history of all request send to the server for translation, ordered in alphabetical order by the english word/sentence. No request body or query string is required for the request.

## Starting Server
The main idea of the deployment process is that the server is supposed to be run as docker container. For this purpose there are docker-compose file and Dockerfile are created. You can start the server via docker container with make compose command e.g. "make compose -port=3000". Now this command has another argument "containerport", which can be passed too. The difference between both flags is that "containerport" is used as container port, meaning the container will be exposed with the value provided in this flag in the internal docker network. The "port" is exposing the container port itself for the docker host. The default value for both flags in the Makefile is 8000.

### TL;DR
You can start the server with "make compose", "make compose -port=**value**" and "make compose -port=**value** -containerport=**value**" where **value** is numeric value.

## Additional Info
If you want you can start the server locally with command "make run -port=**value**", also you can run the tests with "make test". Go check the Makefile too see all available scripts.