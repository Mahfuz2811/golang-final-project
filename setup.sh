#!/bin/bash

# Setup script for Golang Final Project

set -e

echo "🚀 Setting up Golang Final Project..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

# Check if common-net network exists
if ! docker network ls | grep -q common-net; then
    echo "❌ Network 'common-net' not found. Please ensure your MySQL, Redis, and RabbitMQ containers are running on this network."
    echo "You can create the network with: docker network create common-net"
    exit 1
fi

# Build and start the application
echo "🔨 Building Docker images..."
make build

echo "🚀 Starting services..."
make up

echo "📋 Checking service status..."
sleep 5
make status

echo ""
echo "✅ Setup complete!"
echo ""
echo "Your application should be running at: http://localhost:8080"
echo "Health check: http://localhost:8080/health"
echo ""
echo "Available commands:"
echo "  make logs     - View application logs"
echo "  make down     - Stop services"
echo "  make restart  - Restart services"
echo "  make help     - Show all available commands"
