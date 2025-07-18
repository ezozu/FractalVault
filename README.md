# FractalVault - Secure and Scalable Fractal Generation Service

FractalVault is a Go-based service designed to generate and store fractal images on demand. It provides a robust and scalable solution for creating visually stunning fractal art, suitable for integration into various applications, from educational tools to artistic platforms. The service leverages concurrent processing and efficient storage mechanisms to deliver high-quality fractal images with minimal latency. It's built with performance and security in mind, offering configurable parameters to fine-tune the generation process and protect against resource exhaustion.

This project addresses the need for a reliable and customizable fractal generation engine that can be easily integrated into existing systems. Unlike monolithic fractal generation applications, FractalVault provides a service-oriented architecture, enabling developers to offload fractal creation to a dedicated server. This separation of concerns improves application performance and maintainability. FractalVault supports various fractal types, including the Mandelbrot set, Julia sets, and Burning Ship fractals, with customizable parameters for color palettes, iteration limits, and zoom levels. The generated images are stored securely and can be accessed via a REST API, providing a flexible and efficient way to retrieve fractal images.

FractalVault distinguishes itself by its focus on scalability and security. The concurrent processing model allows the service to handle a large number of requests simultaneously, making it suitable for high-traffic environments. Input validation and resource limits are implemented to prevent malicious attacks and ensure system stability. Furthermore, the service is designed to be easily deployable on various platforms, including cloud environments and containerized infrastructures. FractalVault empowers developers to incorporate fractal art into their applications without the complexities of implementing their own fractal generation algorithms.

## Key Features

*   **Concurrent Fractal Generation:** Utilizes Go's goroutines to generate fractal images concurrently, maximizing CPU utilization and minimizing latency. Employs a worker pool pattern to manage the number of concurrent goroutines and prevent resource exhaustion. Specifically, a `sync.WaitGroup` is used to ensure all goroutines complete before returning the generated image.
*   **Configurable Fractal Parameters:** Allows users to customize fractal parameters such as the fractal type (`mandelbrot`, `julia`, `burning_ship`), color palette, maximum iterations, zoom level, and image dimensions. These parameters are validated to prevent invalid or malicious input. The configuration is achieved via command line arguments and environment variables.
*   **Secure Image Storage:** Stores generated fractal images securely, potentially using object storage services like AWS S3 or Google Cloud Storage, or locally on the file system. When using cloud storage, encryption at rest is enabled to protect sensitive data. The storage location and access credentials are configurable via environment variables.
*   **REST API:** Provides a RESTful API for generating and retrieving fractal images. The API supports various request parameters and returns images in common formats such as PNG or JPEG. The API endpoints are documented using Swagger/OpenAPI for easy integration.
*   **Input Validation:** Implements robust input validation to prevent invalid or malicious fractal parameters from causing errors or security vulnerabilities. Regular expressions and range checks are used to ensure that input values are within acceptable limits.
*   **Resource Management:** Implements resource limits, such as maximum memory usage and CPU time, to prevent fractal generation from consuming excessive resources and impacting system performance. Go's `runtime` package is used to monitor memory usage, and the `context` package is used to implement timeouts.
*   **Logging and Monitoring:** Includes comprehensive logging and monitoring capabilities to track the performance and health of the service. Logs are generated using the `log` package and can be configured to output to different destinations, such as files or systemd journal. Metrics are exposed via Prometheus for monitoring with tools like Grafana.

## Technology Stack

*   **Go:** The primary programming language used for developing the service, leveraging its concurrency features and performance.
*   **net/http:** Go's standard library package for building HTTP servers and handling API requests.
*   **image/png, image/jpeg:** Go's standard library packages for encoding and decoding image data in PNG and JPEG formats.
*   **sync:** Go's standard library package for providing synchronization primitives, such as mutexes and wait groups, for concurrent programming.
*   **viper:** A configuration management library for reading configuration from environment variables, files, and command-line arguments.
*   **zap:** A fast, structured, leveled logging package.

## Installation

1.  **Prerequisites:**
    *   Go 1.18 or higher installed and configured.
    *   Git installed.

2.  **Clone the repository:**
    git clone https://github.com/ezozu/FractalVault.git
    cd FractalVault

3.  **Build the application:**
    go build -o fractalvault .

4.  **(Optional) Install Dependencies:**
    While `go build` handles dependencies implicitly, to manage dependencies explicitly with modules, use:
    go mod tidy

## Configuration

FractalVault can be configured using environment variables or a configuration file. The following environment variables are supported:

*   `PORT`: The port on which the service will listen (default: 8080).
*   `STORAGE_TYPE`: The type of storage to use (e.g., `file`, `s3`).
*   `STORAGE_PATH`: The path to store images locally (if `STORAGE_TYPE` is `file`).
*   `AWS_S3_BUCKET`: The name of the S3 bucket to use (if `STORAGE_TYPE` is `s3`).
*   `AWS_REGION`: The AWS region.
*   `MAX_CONCURRENT_REQUESTS`: The maximum number of concurrent fractal generation requests (default: 10).
*   `DEFAULT_ITERATIONS`: The default maximum number of iterations for fractal generation (default: 100).
*   `LOG_LEVEL`: The logging level (e.g., `debug`, `info`, `warn`, `error`) (default: `info`).

Alternatively, you can create a configuration file named `config.yaml` in the same directory as the executable. Refer to the repository's `config.example.yaml` for an example.

## Usage

To start the service, run the following command:
./fractalvault

The service will then be accessible at `http://localhost:8080` (or the port specified by the `PORT` environment variable).

**API Endpoints:**

*   `POST /fractals`: Generates a new fractal image.

    *   Request Body (JSON):
    {
    "type": "mandelbrot",
    "iterations": 200,
    "zoom": 1.5,
    "width": 512,
    "height": 512,
    "color_palette": "rainbow"
    }

    *   Response: Returns the ID of the generated image.

*   `GET /fractals/{id}`: Retrieves a fractal image by its ID.

    *   Response: Returns the fractal image data as a PNG file.

Example using curl:



## Contributing

We welcome contributions to FractalVault! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise commit messages.
4.  Submit a pull request with a detailed description of your changes.
5.  Ensure your code adheres to the Go coding standards.
6.  Include unit tests for your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/ezozu/FractalVault/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the Go community for providing excellent tools and libraries.