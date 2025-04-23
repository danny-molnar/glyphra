# === ENVIRONMENT ===
TOKEN_ENV := .envrc

# === MAIN COMMANDS ===

run:
	@echo "▶️  Running Glyphra plan explanation..."
	go run main.go plan --input plan.out

plan:
	@echo "🛠️  Generating fake Terraform plan..."
	./scripts/gen_plan.sh

blog:
	@echo "📝 Saving last summary to blog post..."
	go run main.go plan --input plan.out | ./scripts/save_summary.sh

lint:
	@echo "🔍 Running go fmt and vet..."
	go fmt ./...
	go vet ./...

test:
	@echo "🧪 Running tests..."
	go test ./...

build:
	@echo "🏗️  Building binary..."
	go build -o glyphra main.go

docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t glyphra .

# === UTILITIES ===

clean:
	@echo "🧹 Cleaning up..."
	rm -f plan.out
	rm -f glyphra
