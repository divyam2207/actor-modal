# Simple Actor Model Communication in Go using Proto.Actor

This project demonstrates basic communication between two actors using the [Proto.Actor](https://github.com/asynkron/protoactor-go) framework in Go. It involves two actors: `Actor_A` and `Actor_B`. `Actor_A` sends a message to `Actor_B`, which processes the message and can respond (currently, no response is sent in this example).

## Project Structure
- `main.go`: The main file where the actor system is created and communication takes place.

## Code Explanation

### Actors:
- **Actor_A**: Sends a message to `Actor_B`.
- **Actor_B**: Receives the message from `Actor_A` and prints the content.

### Messages:
- `RequestMsg`: A simple message that `Actor_A` sends to `Actor_B` with a string payload.
- `AcknowledgmentMsg`: This type of message is defined for further usage where `Actor_B` might acknowledge receipt of the message (though it's not used in this example).

### Communication Flow:
1. **Actor_A** sends a `RequestMsg` to **Actor_B**.
2. **Actor_B** receives the message and prints it to the console.

## How to Run

### Prerequisites:
- [Go](https://golang.org/doc/install)
- [Proto.Actor Go](https://github.com/asynkron/protoactor-go)

### Steps:
1. Clone this repository or download the code files.
2. Install Proto.Actor in your Go environment:
    ```bash
    go get github.com/asynkron/protoactor-go/actor
    ```
3. Run the `main.go` file:
    ```bash
    go run main.go
    ```

### Output:
You should see output similar to:
