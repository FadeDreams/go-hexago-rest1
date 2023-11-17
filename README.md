## Execution Instruction using Docker

## Running with Docker

1. Build the Docker image:

    ```bash
    docker build -t hexagoapp2 .
    ```

2. Run the Docker container, exposing port 8080:

    ```bash
    docker run --name hexagoapp2c -p 8080:8080 hexagoapp2
    ```

3. To test the endpoints of the app while the container is running, you'll need to find the container's IP address. Open a new terminal tab and execute:

    ```bash
    docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hexagoapp2c
    ```

    The command will print something like `172.17.0.2`. You should use this IP address to test the API app endpoints from your machine, not `127.0.0.1`. For example:

    - `http://172.17.0.2:8080/items/1009`

4. To create an item using a POST request, you can use the following `curl` command:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{
        "ItemID": "1011",
        "Name": "Kolahe Ghermez",
        "Description": "Some description",
        "Quantity": 100,
        "Color": "red",
        "Status": "available"
    }' http://172.17.0.2:8080/items
    ```

   This command sends a POST request to create an item with the provided JSON data.


