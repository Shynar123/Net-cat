# TCPChat

TCPChat is a command-line application that implements a Server-Client Architecture similar to NetCat (nc) command. It allows multiple clients to connect to a server and participate in a group chat. The application supports TCP connections, requires clients to provide a name, and provides control over the number of connections allowed.

## Features

1. **TCP Connection**: The server establishes a TCP connection and acts as a central hub for multiple clients to connect to.
2. **Client Name**: Each client is required to provide a unique name when connecting to the server.
3. **Control Connections**: The server can control the number of client connections allowed. _This TCP chat allows up to 10 connections_
4. **Chat Messaging**: Clients can send messages to the chat, which are broadcasted to all other connected clients.
5. **Non-Empty Messages**: Empty messages from clients are not broadcasted to the chat.
6. **Message Timestamp**: Each message sent to the chat is identified with the timestamp and the name of the client who sent it. The format is: `[YYYY-MM-DD HH:MM:SS][client.name]: [client.message]`.
7. **Message History**: When a new client joins the chat, they receive the entire history of previously sent messages.
8. **Client Notifications**: When a client connects or disconnects, the server notifies all other clients about the event.
9. **Message Broadcasting**: All clients receive messages sent by other clients in real-time.
10. **Default Port**: If no port is specified, the default port 8989 is used. If an incorrect usage is provided, the program displays the usage message: `[USAGE]: ./TCPChat $port`.

## Getting Started

To use TCPChat, follow these steps:

1. Clone the repository or download the project files.
2. Compile the code if necessary.
3. Run the server by executing the following command:
   ```
   go run . [port]
   ```
   Replace `[port]` with the desired port number. If no port is specified, the default port 8989 will be used.
4. Connect clients to the server by executing the following command on each client:
   ```
   nc localhost [port]
   ```
   Replace `[port]` with the corresponding port number.
5. Clients will be prompted to enter their name upon connection.
6. Start chatting with other connected clients by typing messages and pressing Enter.

## Usage Examples

Here are some examples of how to use TCPChat:

- Run the server on the default port (8989):

  ```
  go run .
  ```

- Run the server on a specific port:

  ```
  go run . [port]
  ```

- Connect a client to the server:
  ```
  nc localhost [port]
  ```

## Dependencies

TCPChat does not have any external dependencies. It is built using standard networking libraries available in the programming language used.

## Contributors

- [@szhigero](https://01.alem.school/git/szhigero)

- [@ykhamito](https://01.alem.school/git/ykhamito)