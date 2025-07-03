# Docker VM Monitor

**Docker VM Monitor** is an open-source system for monitoring and managing Docker containers across distributed virtual machines. It provides lightweight agents (`monitor-client`) to run on each VM, and a centralized server (`monitor-server`) to collect logs, visualize machine activity, and control containers remotely.

**Note:** This project is in an early stage of development. More features and improvements will be added over time.

## Features

- **Real-time container monitoring** on each VM
- **Remote logging** from clients to the central server via HTTPS
- **Structured log storage** per machine
- **API to list machines**, their IPs, and active containers
- **Start/Stop containers** on individual VMs through HTTP requests
- **API to retrieve logs** for a specific machine

## How it works

1. Each VM runs `monitor-client`, which periodically collects container status and sends logs to monitor-server.
2. `monitor-server` stores logs per VM and exposes APIs for:

- Listing machine information
- Reading logs
- Controlling containers via HTTP

3. A front-end or dashboard can consume these APIs to monitor system health.

## Installation

Clone the repository:

```bash
git clone https://github.com/Josedzzz/test-monitor.git
cd test-monitor
```

### For monitor-client (VM agent):

```bash
cd monitor-client
go build -o client ./cmd
```

Set your `.env` file with:

```bash
SERVER_URL=http://<server-ip>:8080
VM_ID=vm-01
```

Run:

```bash
./client
```

### For monitor-server (central node):

```bash
cd monitor-server
go build -o server ./cmd
./server
```

The server will start at `http://localhost:8080` by default.

## API endpoints

### monitor-server:

- `POST /logs`: Receives logs from client VMs
- `GET /machines`: Get list of all registered VMs
- `GET /logs/{vm_id}`: Get logs for a specific VM

### monitor-client:

- `GET /containers`: Get all the containers info
- `POST /containers/{id}/start`: Start a container by the id
- `POST /containers/{id}/stop`: Stop a container by the id

## Considerations

Since the project is still in early development, there may be bugs or areas where things could be done better. Suggestions, issues, and contributions are more than welcome!

## Contributing

Contributions are welcome! Feel free to fork this repo, submit pull requests, or open issues for bugs and improvements.

## <3

Developed with love for learning, DevOps experimentation, and container monitoring.
