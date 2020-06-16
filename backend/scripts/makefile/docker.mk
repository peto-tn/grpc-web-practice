####################################
# docker
####################################
docker_build:
	docker build ./ -t blog -f docker/build/Dockerfile

docker_run:
	docker run blog
