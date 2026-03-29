````markdown
## Getting Started

### 1. Start the server

```bash
make docker-run
```

### 2. Run migrations

Run the following migrations **in order**:

```bash
make migration
# enter: 001_create_hospital

make migration
# enter: 002_create_patient

make migration
# enter: 003_create_staff
```

> Migration filenames can be found in `runner.go`

### 3. Seed hospital data

```bash
make seed-hospital
```
````
