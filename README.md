# weather-api-server

A server that accepts multiple cities as input and fetches real-time weather from various weather APIs. It responds with the real-time weather results.

# Description

The server is integrated with a free weather API that accepts one city at a time. However, the weather-api-server allows users to request weather information for multiple cities simultaneously by iterating over the user's input cities on the server and querying the free weather API accordingly.

# Features

- Accepts multiple cities as input for weather information retrieval.
- Utilizes Go's performance capabilities for efficient processing of multiple requests.
- Integrates with a free weather API to fetch real-time weather data.
- Responds with the real-time weather results to the client.

## Installation

To run the weather-api-server, follow these steps:

1.Ensure you have Go installed on your system. If not, you can download and install it from the official Go website: https://golang.org/

2.Clone the repository:

```bash
git clone https://github.com/ThanmayNath/weather-api-server.git
```

3.Navigate to the project directory:

```bash
cd weather-api-server
```

4.Install the dependencies:

```bash
go mod download
```

5.Set up your free Open Weather API from rapidapi.com
6.Run the server:

```bash
go run main.go
```

The server should now be up and running on`http://localhost:8800/weather`

## Usage

To retrieve weather information for multiple cities, make a POST request to the server's endpoint with the desired cities. The endpoint URL format is as follows:

```bash
http://localhost:8800/weather
```

```bash
{
  "cities": [
    "toronto",
    "mumbai",
    "london"
  ]
}
```

A successful response from the server will result in the following output:

```bash
[
    {
        "city": "toronto",
        "temperature": 0
    },
    {
        "city": "mumbai",
        "temperature": 0
    },
    {
        "city": "london",
        "temperature": 0
    }
]
```

## Contributing

Contributions to the weather-api-server project are welcome! If you find a bug or have an idea for a new feature, please open an issue or submit a pull request on the GitHub repository.
