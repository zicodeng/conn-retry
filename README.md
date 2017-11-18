# Connection Auto-Retry Example

Some Docker containers, such as RabbitMQ, have to do some work at startup before they are ready to accept connections from clients. If you start your RabbitMQ container and your container that wants to connect to RabbitMQ at the same time (e.g., from a bash script), you might encounter an error when trying to connect to RabbitMQ because that container isn't yet done with its startup work, and is therefore not ready for new connections.

Knowing this, your container can automatically retry its connection to RabbitMQ for a set amount of times. Each time it sleeps for a progressively longer interval. After the max number of connection retries, it should report the final error and halt your program.

This repo contains an example of this technique. The `main.go` file starts a goroutine that tries to connect to a RabbitMQ server at the address supplied in the `MQADDR` environment variable. If it encounters an error, it will retry the connection, up to `maxConnRetries` times.

To run this example, execute the `build.sh` script to build the go executable and Docker container. Then run the `start.sh` script to create a Docker private network, run the RabbitMQ container, and run the go client container. Use `docker logs gateway` to inspect the logs of the go client container. Use the `stop.sh` script to stop/remove the containers, and remove the private network.

```bash
# build the go client container
./build.sh
# create a private network and start the containers
./start.sh
# follow the logs of the gateway (ctrl+c to stop)
docker logs --follow gateway
# tear-down everything
./stop.sh
```
