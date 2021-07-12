# Apache Kafka examples

A simple Kafka testing environment for different libraries and languages.

## Runnning locally

- In order to setup the infra (Zookeeper and Kafka broker) runs:

    ```bash
    docker-compose up -d
    ```

- Next commands should be executed on the kafka container, so first log in into the container by typing:

    ``` bash
    docker-compose exec kafka bash
    ```

- Then inside the container, runs:

    ```bash
    # To create a topic
    /bin/kafka-topics --create --topic topic-name --bootstrap-server localhost:9092
    ```

    ```bash
    # Runs a consumer
    /bin/kafka-console-consumer --topic topic-name --from-beginning --bootstrap-server localhost:9092
    ```

    ```bash
    # Runs a producer
    /bin/kafka-console-producer --topic topic-name --bootstrap-server localhost:9092
    ```
