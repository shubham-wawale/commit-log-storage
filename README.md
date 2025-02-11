# **Distributed Log Service in Go**

## **Overview**

This project implements a **Distributed Log Service** in **Go**, designed for efficient storage, indexing, and retrieval of log records. Logs are essential in distributed systems for tracking transactions, commits, and events.

## **Problem Statement**

Logs need to be **fast, reliable, and scalable**. Simply appending records to a file is inefficient for retrieval. This project addresses these challenges by implementing:

- **Efficient Storage**: Logs are stored as structured binary data.
- **Indexing**: A separate index file maps offsets to log positions, enabling quick lookups.
- **Memory Mapping**: `gommap` is used for high-performance file access.
- **Networking & Security**: gRPC with TLS ensures secure communication.

## **Project Components**

1. **Store** - Handles writing and reading log records from disk.
2. **Index** - Maps record offsets to file positions for fast lookups.
3. **Networking** - gRPC API for log service interactions.
4. **Security** - Implements TLS for encrypted data transmission.

## **Installation & Usage**

```sh
git clone https://github.com/yourusername/distributed-log.git
cd distributed-log
go run main.go
```

## **Key Technologies**

- **Go** (Golang)
- **gRPC** for networking
- **Memory-mapped files** for performance
- **TLS Security**

## **License**

This project is open-source under the MIT License.
