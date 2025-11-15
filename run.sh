#!/bin/bash

# Dev Echo Server - One-Command Runner
# This script handles everything needed to run the application

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print colored output
print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

# Print banner
echo -e "${BLUE}"
echo "╔════════════════════════════════════════╗"
echo "║      Dev Echo Server Runner            ║"
echo "╔════════════════════════════════════════╗"
echo -e "${NC}"

# Check if .env file exists, if not create from example
if [ ! -f .env ]; then
    print_warning ".env file not found"
    if [ -f .env.example ]; then
        print_info "Creating .env from .env.example..."
        cp .env.example .env
        print_success ".env file created"
    else
        print_error ".env.example not found!"
        exit 1
    fi
else
    print_success ".env file found"
fi

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
    print_success "Environment variables loaded"
fi

# Function to check if Docker is installed
check_docker() {
    if command -v docker &> /dev/null; then
        # Check for docker compose (new) or docker-compose (old)
        if docker compose version &> /dev/null || command -v docker-compose &> /dev/null; then
            return 0
        fi
    fi
    return 1
}

# Function to run docker compose (handles both new and old syntax)
run_docker_compose() {
    if docker compose version &> /dev/null 2>&1; then
        docker compose "$@"
    elif command -v docker-compose &> /dev/null; then
        docker-compose "$@"
    else
        print_error "Neither 'docker compose' nor 'docker-compose' found!"
        return 1
    fi
}

# Function to check if Go is installed
check_go() {
    if command -v go &> /dev/null; then
        return 0
    else
        return 1
    fi
}

# Main menu
echo ""
echo "Select run mode:"
echo "  1) Docker (Recommended - No Go installation needed)"
echo "  2) Native Go (Requires Go installed)"
echo "  3) Build Docker image only"
echo "  4) Stop and remove containers"
echo "  5) Exit"
echo ""
read -p "Enter choice [1-5]: " choice

case $choice in
    1)
        print_info "Running with Docker..."
        if ! check_docker; then
            print_error "Docker not found or not properly configured!"
            print_info "Please install Docker: https://docs.docker.com/get-docker/"
            echo ""
            print_warning "Falling back to native Go option..."
            sleep 2
            choice=2
        fi
        
        if [ "$choice" = "1" ]; then
            print_info "Building and starting containers..."
            run_docker_compose up --build -d
            
            print_success "Container started successfully!"
            print_info "Server running at: http://localhost:${PORT:-3000}"
            echo ""
            print_info "View logs with: docker compose logs -f (or docker-compose logs -f)"
            print_info "Stop server with: docker compose down (or docker-compose down)"
        fi
        ;;&  # Fall through to case 2 if docker failed
    
    2)
        print_info "Running with native Go..."
        if ! check_go; then
            print_error "Go not found!"
            print_info "Please install Go: https://golang.org/doc/install"
            exit 1
        fi
        
        print_info "Downloading dependencies..."
        go mod download
        
        print_info "Building application..."
        go build -o dev-echo-server .
        
        print_success "Build complete!"
        print_info "Starting server..."
        
        # Check if server is already running
        if pgrep -f "dev-echo-server" > /dev/null; then
            print_warning "Server already running. Stopping it first..."
            pkill -f dev-echo-server
            sleep 1
        fi
        
        ./dev-echo-server &
        SERVER_PID=$!
        
        sleep 2
        
        if ps -p $SERVER_PID > /dev/null; then
            print_success "Server started successfully! (PID: $SERVER_PID)"
            print_info "Server running at: http://localhost:${PORT:-3000}"
            echo ""
            print_info "Stop server with: pkill -f dev-echo-server"
            print_info "Or: kill $SERVER_PID"
        else
            print_error "Failed to start server"
            exit 1
        fi
        ;;
    
    3)
        print_info "Building Docker image..."
        if ! check_docker; then
            print_error "Docker not found!"
            exit 1
        fi
        
        run_docker_compose build
        print_success "Docker image built successfully!"
        ;;
    
    4)
        print_info "Stopping containers..."
        if check_docker; then
            run_docker_compose down
            print_success "Containers stopped and removed"
        else
            print_error "Docker not found!"
        fi
        ;;
    
    5)
        print_info "Exiting..."
        exit 0
        ;;
    
    *)
        print_error "Invalid choice!"
        exit 1
        ;;
esac

echo ""
print_success "Done!"
