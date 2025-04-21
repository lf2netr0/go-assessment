# Go Ethereum API - WETH Total Supply

This is a simple Go API project built with the Gin framework that provides an endpoint to fetch the current total supply of Wrapped Ether (WETH) from the Ethereum mainnet.

## Endpoint

- **GET `/weth-total-supply`**: Returns the total supply of WETH.

  **Example Response:**

  ```json
  {
    "totalSupply": "3019840729876543210987654" 
  }
  ```
  *(Note: The actual value will change based on the current state of the Ethereum blockchain.)*

## Prerequisites

- Go (version 1.21 or higher recommended)

## Running the Project

1.  **Clone the repository (if you haven't already):**
    ```bash
    git clone <your-repository-url>
    cd go-assessment 
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Run the server:**
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

4.  **Access the endpoint:**
    Open your browser or use `curl`:
    ```bash
    curl http://localhost:8080/weth-total-supply
    ``` 