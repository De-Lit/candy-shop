server-build:
	docker build --file=docker/server.docker --tag candy-server .

server-run:
	docker run --rm -d -p 3333:3333 --name candy-server candy-server

server-stop:
	docker stop candy-server