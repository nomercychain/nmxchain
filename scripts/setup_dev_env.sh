#!/bin/bash

# NoMercyChain Development Environment Setup Script
# This script sets up the development environment for the NoMercyChain project

echo "Setting up NoMercyChain development environment..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go..."
    
    # Detect OS
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        wget https://go.dev/dl/go1.18.10.linux-amd64.tar.gz
        sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.10.linux-amd64.tar.gz
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
        echo 'export GOPATH=$HOME/go' >> ~/.profile
        echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.profile
        source ~/.profile
        rm go1.18.10.linux-amd64.tar.gz
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        brew install go@1.18
    else
        echo "Unsupported OS. Please install Go manually from https://golang.org/dl/"
        exit 1
    fi
fi

# Verify Go installation
go version
if [ $? -ne 0 ]; then
    echo "Failed to install Go. Please install manually and try again."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "Node.js is not installed. Installing Node.js..."
    
    # Detect OS
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
        sudo apt-get install -y nodejs
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        brew install node@16
    else
        echo "Unsupported OS. Please install Node.js manually from https://nodejs.org/"
        exit 1
    fi
fi

# Verify Node.js installation
node -v
if [ $? -ne 0 ]; then
    echo "Failed to install Node.js. Please install manually and try again."
    exit 1
fi

# Verify npm installation
npm -v
if [ $? -ne 0 ]; then
    echo "Failed to install npm. Please install manually and try again."
    exit 1
fi

# Set up Go dependencies
echo "Setting up Go dependencies..."
go mod tidy

# Set up frontend dependencies
echo "Setting up frontend dependencies..."
cd client
npm install
cd ..

# Create necessary directories
echo "Creating necessary directories..."
mkdir -p scripts/local_testnet

# Set up local testnet script
cat > scripts/local_testnet/setup_local_testnet.sh << 'EOF'
#!/bin/bash

# NoMercyChain Local Testnet Setup Script
# This script sets up a local testnet for development and testing

echo "Setting up NoMercyChain local testnet..."

# Initialize the chain
nmxchaind init local-testnet --chain-id nomercychain-local-1

# Create test accounts
nmxchaind keys add validator --keyring-backend test
nmxchaind keys add user1 --keyring-backend test
nmxchaind keys add user2 --keyring-backend test

# Add genesis accounts
nmxchaind add-genesis-account $(nmxchaind keys show validator -a --keyring-backend test) 10000000000unmx
nmxchaind add-genesis-account $(nmxchaind keys show user1 -a --keyring-backend test) 1000000000unmx
nmxchaind add-genesis-account $(nmxchaind keys show user2 -a --keyring-backend test) 1000000000unmx

# Create validator transaction
nmxchaind gentx validator 5000000000unmx \
  --chain-id nomercychain-local-1 \
  --moniker="local-validator" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --keyring-backend test

# Collect gentxs
nmxchaind collect-gentxs

# Validate genesis
nmxchaind validate-genesis

# Start the chain
echo "Starting the local testnet..."
nmxchaind start
EOF

chmod +x scripts/local_testnet/setup_local_testnet.sh

# Create build script
cat > scripts/build.sh << 'EOF'
#!/bin/bash

# NoMercyChain Build Script
# This script builds the NoMercyChain binary

echo "Building NoMercyChain..."

# Build the binary
go build -o build/nmxchaind ./cmd/nmxchaind

echo "Build completed. Binary is located at build/nmxchaind"
EOF

chmod +x scripts/build.sh

# Create test script
cat > scripts/run_tests.sh << 'EOF'
#!/bin/bash

# NoMercyChain Test Script
# This script runs all tests for the NoMercyChain project

echo "Running NoMercyChain tests..."

# Run Go tests
echo "Running Go tests..."
go test -v ./...

# Run frontend tests
echo "Running frontend tests..."
cd client
npm test
cd ..

echo "All tests completed."
EOF

chmod +x scripts/run_tests.sh

echo "Development environment setup complete!"
echo ""
echo "Next steps:"
echo "1. Review the DEVELOPMENT_GUIDE.md file for detailed instructions"
echo "2. Run 'scripts/build.sh' to build the NoMercyChain binary"
echo "3. Run 'scripts/local_testnet/setup_local_testnet.sh' to set up a local testnet"
echo "4. Run 'scripts/run_tests.sh' to run all tests"
echo ""
echo "Happy coding!"