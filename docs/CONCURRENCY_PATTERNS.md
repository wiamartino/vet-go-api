# Concurrency Patterns Quick Reference

## Thread-Safe Patterns Used in vet-go

### 1. Singleton with Double-Checked Locking
**Location**: `infrastructure/database/database.go`

```go
var (
    dbInstance *DB
    mu         sync.RWMutex
)

func ConnectDatabase() (*DB, error) {
    // First check (with read lock)
    mu.RLock()
    if dbInstance != nil {
        mu.RUnlock()
        return dbInstance, nil
    }
    mu.RUnlock()
    
    // Acquire write lock for initialization
    mu.Lock()
    defer mu.Unlock()
    
    // Double-check after acquiring write lock
    if dbInstance != nil {
        return dbInstance, nil
    }
    
    // Initialize
    database, err := gorm.Open(...)
    if err != nil {
        return nil, err
    }
    
    dbInstance = &DB{DB: database}
    return dbInstance, nil
}
```

**Why it works**: 
- Read locks for frequent checks (fast path)
- Write lock only when actually initializing
- Double-check prevents race during lock acquisition

### 2. Once Initialization
**Location**: `utils/jwt/jwt.go`

```go
var (
    once sync.Once
    mu   sync.RWMutex
)

func init() {
    initializeJWT()
}

func initializeJWT() {
    once.Do(func() {
        loadJWTConfig()
    })
}

func loadJWTConfig() {
    mu.Lock()
    defer mu.Unlock()
    
    jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
    isInitialized = true
}
```

**Why it works**:
- `sync.Once` guarantees function runs exactly once
- Perfect for initialization code
- No contention after first call

### 3. Atomic Counters
**Location**: `middlewares/metrics.go`

```go
var totalRequests int64
var totalErrors int64

func MetricsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Increment request counter atomically
        reqCount := atomic.AddInt64(&totalRequests, 1)
        
        // Conditionally increment error counter
        if c.Writer.Status() >= http.StatusBadRequest {
            errCount = atomic.AddInt64(&totalErrors, 1)
        } else {
            errCount = atomic.LoadInt64(&totalErrors)
        }
    }
}
```

**Why it works**:
- Lock-free operation (faster than mutexes)
- Atomic operations are thread-safe
- Perfect for counters and statistics

### 4. Read-Write Mutex
**Location**: `infrastructure/database/database.go`

```go
// Many readers, few writers pattern
func GetDB() *DB {
    mu.RLock()  // Multiple goroutines can hold read lock
    defer mu.RUnlock()
    return dbInstance
}

func updateConfig() {
    mu.Lock()   // Exclusive lock for writing
    defer mu.Unlock()
    // Update configuration
}
```

**Why it works**:
- Optimized for read-heavy workloads
- Multiple readers can proceed concurrently
- Writers get exclusive access

## Common Pitfalls to Avoid

### ❌ DON'T: Unprotected shared state
```go
// WRONG - Race condition!
var counter int

func increment() {
    counter++  // Not thread-safe!
}
```

### ✅ DO: Use atomic operations
```go
// CORRECT
var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}
```

### ❌ DON'T: Defer unlock before checking
```go
// WRONG - Holding lock too long!
func getData() Data {
    mu.Lock()
    defer mu.Unlock()  // Lock held for entire function
    
    if cachedData != nil {
        return cachedData
    }
    
    // Expensive operation while holding lock
    cachedData = fetchFromDatabase()
    return cachedData
}
```

### ✅ DO: Minimize lock duration
```go
// CORRECT
func getData() Data {
    mu.RLock()
    if cachedData != nil {
        data := cachedData
        mu.RUnlock()
        return data
    }
    mu.RUnlock()
    
    // Expensive operation without lock
    newData := fetchFromDatabase()
    
    mu.Lock()
    cachedData = newData
    mu.Unlock()
    
    return newData
}
```

## Testing for Race Conditions

### During Development
```bash
# Test single package
go test -race ./package/...

# Test with verbose output
go test -race -v ./...

# Test specific function
go test -race -run TestFunctionName
```

### In CI/CD
```bash
# Add to CI pipeline
go test -race -short ./...

# With coverage
go test -race -coverprofile=coverage.out ./...
```

### Stress Testing
```bash
# Run tests multiple times
go test -race -count=100 ./...

# Parallel execution
go test -race -parallel=10 ./...
```

## Quick Checklist

Before committing code with shared state:

- [ ] Is shared data protected by mutex or atomic operations?
- [ ] Are locks released in all code paths (use defer)?
- [ ] Is lock duration minimized?
- [ ] Have you run `go test -race`?
- [ ] Are there tests covering concurrent access?

## Race Detector Flags

```bash
# Enable race detector
go test -race

# More sensitive detection (slower)
GORACE="halt_on_error=1" go test -race

# Log all races to file
GORACE="log_path=/tmp/race" go test -race

# Custom options
GORACE="halt_on_error=1 log_path=./race.log" go test -race
```

## When to Use What

| Pattern | Use Case | Performance |
|---------|----------|-------------|
| `sync.Mutex` | Protecting shared state | Moderate |
| `sync.RWMutex` | Read-heavy workloads | Good for reads |
| `sync.Once` | One-time initialization | Excellent |
| `atomic.*` | Simple counters/flags | Excellent |
| Channels | Communication between goroutines | Good |

## Resources

- [Go Race Detector](https://go.dev/doc/articles/race_detector)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Memory Model](https://go.dev/ref/mem)
