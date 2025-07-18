# FractalVault: Decentralized NFT Marketplace Scaffold

A secure and scalable platform for NFT trading, leveraging advanced cryptographic techniques for ownership verification.

## Detailed Description

FractalVault is a Go-based project designed as a decentralized NFT marketplace scaffold, prioritizing security and efficient ownership management. This repository provides a robust foundation for building a fully functional NFT trading platform that operates independently of centralized authorities. It distinguishes itself through its innovative use of elliptic curve cryptography for user authentication and Merkle tree-based ownership verification, implemented within a WASM smart contract. This combination ensures secure and transparent NFT transactions, while also optimizing gas costs associated with on-chain operations.

The project aims to address common challenges in existing NFT marketplaces, such as security vulnerabilities and high transaction fees. By leveraging elliptic curve cryptography, FractalVault minimizes the risk of unauthorized access and manipulation of user accounts. The Merkle tree implementation significantly reduces the amount of data that needs to be stored on-chain, thereby decreasing gas costs and improving the overall efficiency of the marketplace. The WASM smart contract facilitates portability and interoperability across different blockchain platforms, allowing the marketplace to adapt to evolving technological landscapes.

FractalVault provides developers with a comprehensive set of tools and libraries to facilitate the integration of its core functionalities into their own projects. The Go codebase is well-documented and modular, allowing developers to easily customize and extend the platform's features. This scaffold includes API endpoints for NFT creation, listing, bidding, and purchase, as well as utilities for managing user accounts and handling cryptographic operations. By utilizing FractalVault, developers can accelerate the development of secure and efficient decentralized NFT marketplaces, contributing to a more robust and accessible Web3 ecosystem.

## Key Features

*   **Elliptic Curve Cryptography (ECC) Authentication:** Utilizes secp256k1 for secure user authentication, providing robust protection against unauthorized access and account hijacking. Keys are managed locally, empowering users with complete control over their accounts. The underlying implementation relies on the `crypto/ecdsa` package in Go's standard library, ensuring a battle-tested and reliable cryptographic foundation.

*   **Merkle Tree-Based Ownership Verification:** Employs a Merkle tree structure to efficiently verify NFT ownership on-chain. This reduces gas costs significantly compared to traditional on-chain storage methods. The Merkle tree is constructed off-chain, and only the root hash is stored on the smart contract. Verification involves providing a Merkle proof, which is verified by the WASM smart contract.

*   **WASM Smart Contract:** The core logic for NFT management and transaction processing is implemented as a WebAssembly (WASM) smart contract. This allows the marketplace to be deployed on various blockchain platforms that support WASM, enhancing interoperability and portability. The WASM contract is written in Rust and compiled using `wasm-pack`.

*   **Go-Based Backend:** The backend infrastructure is built using Go, providing high performance and scalability. Go's concurrency model allows for efficient handling of multiple concurrent requests. The backend provides API endpoints for interacting with the WASM smart contract and managing user data.

*   **NFT Metadata Management:** Provides utilities for managing NFT metadata, including support for standard metadata formats like ERC-721. Metadata can be stored off-chain using decentralized storage solutions like IPFS, with the CID (Content Identifier) stored on-chain within the NFT's properties.

*   **Secure Transaction Processing:** Implements secure transaction processing mechanisms to prevent double-spending and other common vulnerabilities. Transactions are signed using the user's private key and verified by the WASM smart contract before being executed.

*   **API Endpoints for Marketplace Operations:** Includes a comprehensive set of API endpoints for key marketplace functionalities, such as NFT creation, listing, bidding, purchasing, and transferring. These endpoints are designed to be easily integrated into front-end applications.

## Technology Stack

*   **Go:** The primary programming language used for the backend infrastructure. Provides high performance, concurrency, and a rich set of libraries.
*   **Rust:** Used for developing the WASM smart contract due to its safety, performance, and excellent support for WebAssembly.
*   **WebAssembly (WASM):** Enables the smart contract to run efficiently on various blockchain platforms.
*   **Elliptic Curve Cryptography (secp256k1):** Used for secure user authentication and digital signatures.
*   **Merkle Trees:** Employed for efficient NFT ownership verification.
*   **IPFS (InterPlanetary File System):** Used for decentralized storage of NFT metadata.

## Installation

1.  **Install Go:** Ensure that Go (version 1.18 or later) is installed on your system. Instructions can be found at [https://go.dev/dl/](https://go.dev/dl/).

2.  **Install Rust and `wasm-pack`:** Rust is required to compile the WASM smart contract. Install Rust using `rustup`: [https://www.rust-lang.org/tools/install](https://www.rust-lang.org/tools/install). Then, install `wasm-pack`: `cargo install wasm-pack`.

3.  **Clone the Repository:**
    

4.  **Build the WASM Smart Contract:** Navigate to the `wasm` directory and build the contract:
    

5.  **Install Go Dependencies:**
    

6.  **Build the Go Backend:**
    

## Configuration

The FractalVault backend requires several environment variables to be configured. Create a `.env` file in the root directory of the project and define the following variables:

*   `PORT`: The port on which the backend server will listen (e.g., `8080`).
*   `DATABASE_URL`: The connection string for the database (e.g., using PostgreSQL). This is where user data and NFT metadata will be stored.
*   `WASM_CONTRACT_PATH`: The path to the compiled WASM contract file (e.g., `wasm/pkg/fractalvault_bg.wasm`).
*   `IPFS_GATEWAY`: The URL of the IPFS gateway to use for accessing NFT metadata (e.g., `https://ipfs.io/ipfs/`).
*   `PRIVATE_KEY`: A private key used for signing transactions related to contract deployment and maintenance. **This should be kept highly secure.**

Load these environment variables in your application, for example using a library like `github.com/joho/godotenv`.

## Usage

The FractalVault backend provides a REST API for interacting with the NFT marketplace. Example endpoints include:

*   `POST /api/users/register`: Registers a new user. Requires a public key (derived from their elliptic curve keypair) and other user information.
    Request Body:
    {
    "publicKey": "...",
    "username": "..."
    }

*   `POST /api/nfts/create`: Creates a new NFT. Requires metadata and the user's signature.
    Request Body:
    {
    "metadataCID": "...",
    "ownerPublicKey": "...",
    "signature": "..."
    }

*   `GET /api/nfts/{nftId}`: Retrieves information about a specific NFT.

*   `POST /api/nfts/{nftId}/transfer`: Transfers ownership of an NFT. Requires the sender's signature and the recipient's public key.
    Request Body:
    {
    "recipientPublicKey": "...",
    "signature": "..."
    }

Detailed API documentation, including request and response formats, will be provided separately.

## Contributing

We welcome contributions to FractalVault! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise code with thorough documentation.
4.  Submit a pull request with a detailed description of your changes.
5.  Ensure that your code adheres to the project's coding style and passes all tests.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/ezozu/FractalVault/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the open-source community for their contributions to the technologies used in this project.