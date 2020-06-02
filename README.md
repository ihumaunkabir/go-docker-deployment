## Deploy go application in Docker Container

A go application has been deployed to docker container using an image configuration based on the requirements.


## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites
* Docker


### Installation
 
1. Clone the repo
```sh
git clone https://github.com/oasiscse/go-docker-deployment.git
```
2. Go to project directory and build image -
```sh
docker build . -t go-docker-dep
```
3. To run the image (a running image is container)
```sh
docker run -p 3000:3000 go-docker-dep
```

## Test the server is running
1. Go to http://localhost:3000
(YourHostIP:3000)

Response
```sh
{"Greetings" : "You are awesome!"}
```
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Open a Pull Request

## License

Distributed under the MIT License.

## Contact

Tweet me - [@oasiscse](https://twitter.com/oasiscse)  
Find me on [ihumaun.com](http://ihumaun.com)
