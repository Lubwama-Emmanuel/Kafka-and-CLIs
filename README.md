# Chat CLI Application

[![codecov](https://codecov.io/gh/Lubwama-Emmanuel/Kafka-and-CLIs/branch/intial-setup/graph/badge.svg?token=VD1KMP2GG9)](https://codecov.io/gh/Lubwama-Emmanuel/Kafka-and-CLIs)

## Description
A command-line-driven program that allows message exchange between two clients(producer to consumer) thereby micking chatting and uses kafka as a message broker. <br>
The application uses cobra-cli, a command line program to generate cobra command files. With this library, client are able to interact with the application by passing in commands in the cli.

## Code style
Dependency Injection - a technique that uses programming practices to leverage decoupling amongst code. Pieces of code that have no logical correlation are separated from one another and just injected to establish a dependency. <a href="https://www.developer.com/languages/golang-dependency-injection">See more..</a> <br>
This project makes use of the above technique by creating an interface provider that creates an instance of a message broker and this case kafka. Other message brokers such as RabbitMQ, Redis can be used by simply creating their provider interfaces and simply injecting their instances into the application business logic.

## Usage
With one instance of the program, someone can run the command:<br>
<code>
    go run main.go send --channel channel_name --server “server:port” --group group_name --message message_to_send
</code>

With a second instance of the same program <br>
<code>
    go run main.go receive --channel channel_name --from start_from --server “server:port” --group group_name 
</code>

Where;<br>
commands:
<li>send command allows client to subscribe as a producer</li>
<li>receive comand allows client to subscribe as a consumer</li>

flags:
<li>Channel => topic subscribed to</li>
<li>Server => port for connection</li>
<li>Group => group id to join</li>
<li>Message => messages to send to the subscribed consumers</li>

## Additional Information
<li>Run each instance in a new terminal</li>
<li>You can run a one to many or a many to many communication</li>
