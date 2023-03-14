# Chat CLI Application

[![codecov](https://codecov.io/gh/Lubwama-Emmanuel/Kafka-and-CLIs/branch/intial-setup/graph/badge.svg?token=VD1KMP2GG9)](https://codecov.io/gh/Lubwama-Emmanuel/Kafka-and-CLIs)

## Description
A command-line-driven program that allows message exchange between consumers and producers and uses kafka as a message broker. 

## Usage
With one instance of the program, someone can run the command:<br>
<code>
    go main.go send --channel channel_name --server “server:port” --group group_name --message message_to_send
</code>

With a second instance of the same program <br>
<code>
    go main.go receive --channel channel_name --from start_from --server “server:port” --group group_name 
</code>

## Additional Information
<ul>
<li>Run each instance in a new terminal</li>
<li>You can run a one to many or a many to many communication</li>
</ul>
