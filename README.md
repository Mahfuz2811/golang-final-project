# Golang Final Project - Dockerized

A containerized Go REST API with user authentication and product management, featuring JWT authentication, Redis caching, and RabbitMQ messaging.

## 🏗️ Architecture

- **Backend**: Go with Gin framework
- **Database**: MySQL (with GORM)
- **Cache**: Redis
- **Message Queue**: RabbitMQ
- **Authentication**: JWT tokens
- **Container**: Docker & Docker Compose

## 📋 Prerequisites

- Docker and Docker Compose installed
- Existing containers running on `common-net` network:
  - MySQL (port 3306)
  - Redis (port 6379)
  - RabbitMQ (port 5672, management 15672)

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <your-repo>
cd golang-final-project
```

### 2. Configure Environment

Edit `docker/.envs/app.env` if needed:

```env
# Application Configuration
APP_PORT=8080
GIN_MODE=release

# Database Configuration
DB_HOST=common-mysql-1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=secret
DB_NAME=registration

# Redis Configuration
REDIS_HOST=common-redis-1
REDIS_PORT=6379
REDIS_PASSWORD=redispassword
REDIS_DB=0

# RabbitMQ Configuration
RABBITMQ_HOST=common-rabbitmq-1
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_QUEUE=email_queue

# JWT Configuration
JWT_SECRET=my_secret_key
```

### 3. Build and Run

Using the automated setup script:

```bash
./setup.sh
```

Or manually:

```bash
# Build images
docker-compose -f docker/docker-compose.yaml build

# Start services
docker-compose -f docker/docker-compose.yaml up -d

# Check status
docker-compose -f docker/docker-compose.yaml ps
```

### 4. Verify Deployment

- **Health Check**: http://localhost:8080/health
- **API Base**: http://localhost:8080

## 📚 API Endpoints

### Authentication

- `POST /register` - User registration
- `POST /login` - User login
- `GET /user` - Get user info (protected)

### Products (Protected)

- `POST /products/create` - Create product
- `GET /products/all` - Get all products

### Example Usage

```bash
# Register user
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# Login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Use returned token for protected endpoints
curl -X GET http://localhost:8080/user \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🛠️ Development

### Available Make Commands

```bash
make help        # Show all available commands
make build       # Build Docker images
make up          # Start services
make down        # Stop services
make logs        # View logs
make logs-app    # View app logs only
make restart     # Restart services
make clean       # Clean up containers and images
make rebuild     # Rebuild and restart
make status      # Show service status
make migrate     # Run database migration
make shell       # Open shell in app container
```

### Development Mode

For development with hot reload:

```bash
# Use override file for development
docker-compose -f docker/docker-compose.yaml -f docker/docker-compose.override.yaml up
```

### Running Tests

```bash
# Run tests locally
go test ./...

# Run specific test packages
go test ./services/...
```

## 🔧 Configuration

### Environment Variables

All configuration is handled through environment variables in `docker/.envs/app.env`:

| Variable         | Description              | Default           |
| ---------------- | ------------------------ | ----------------- |
| `APP_PORT`       | Application port         | 8080              |
| `GIN_MODE`       | Gin mode (debug/release) | release           |
| `DB_HOST`        | MySQL hostname           | common-mysql-1    |
| `DB_USER`        | MySQL username           | root              |
| `DB_PASSWORD`    | MySQL password           | secret            |
| `DB_NAME`        | Database name            | registration      |
| `REDIS_HOST`     | Redis hostname           | common-redis-1    |
| `REDIS_PASSWORD` | Redis password           | redispassword     |
| `RABBITMQ_HOST`  | RabbitMQ hostname        | common-rabbitmq-1 |
| `JWT_SECRET`     | JWT signing secret       | my_secret_key     |

### Network Configuration

The application connects to existing services via the `common-net` Docker network:

```bash
# Verify network exists
docker network ls | grep common-net

# If needed, create the network
docker network create common-net
```

## 📊 Monitoring & Logs

### View Logs

```bash
# All services
make logs

# Application only
make logs-app

# Follow logs
docker-compose -f docker/docker-compose.yaml logs -f golang-final-project
```

### Health Monitoring

- Application health: http://localhost:8080/health
- Redis Commander: http://localhost:8081 (if available)
- RabbitMQ Management: http://localhost:15672

## 🐛 Troubleshooting

### Common Issues

1. **Connection refused errors**

   - Ensure all prerequisite containers are running
   - Check network connectivity: `docker network inspect common-net`

2. **Migration fails**

   - Check database credentials in environment file
   - Ensure MySQL container is healthy

3. **Build failures**
   - Clear Docker cache: `docker builder prune`
   - Rebuild: `make rebuild`

### Debug Mode

Enable debug mode for more verbose logging:

```bash
# Edit docker/.envs/app.env
GIN_MODE=debug
```

## 🏗️ Project Structure

```
.
├── cmd/                    # CLI commands
├── db/                     # Database configuration
├── docker/                 # Docker configuration
│   ├── app/               # Application Dockerfile
│   ├── .envs/             # Environment files
│   └── docker-compose.yaml
├── handlers/              # HTTP handlers
├── middlewares/           # HTTP middlewares
├── models/                # Data models
├── repositories/          # Data access layer
├── services/              # Business logic
├── utils/                 # Utilities
├── Makefile              # Development commands
└── setup.sh              # Automated setup script
```

## 📝 Contributing

1. Make changes to the codebase
2. Test locally: `go test ./...`
3. Build Docker image: `make build`
4. Test in container: `make up`
5. Submit pull request

## 📄 License

[Your License Here]
