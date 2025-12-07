# Makefile for ITS Presentation (Marp)
# CO3017 - Software Architecture

# Variables
MARP := marp
SLIDES := slides.md
OUTPUT_DIR := .
THEME := theme.css
MERMAID_DIAGRAM := term.md
DIAGRAM_OUTPUT := images/architectural-characteristics-radar.png
MERMAID_CLI := mmdc

# Output files
PDF := slides.pdf
HTML := slides.html
PPTX := slides.pptx

# Marp options
MARP_OPTS := --allow-local-files
THEME_OPTS := --theme $(THEME)

# Colors for output
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
NC := \033[0m # No Color

.PHONY: all pdf html pptx preview clean help install check diagram install-mermaid check-mermaid

# Default target
all: pdf html
	@echo "$(GREEN)✓ All formats generated successfully!$(NC)"

# Generate PDF
pdf: check
	@echo "$(BLUE)→ Generating PDF...$(NC)"
	$(MARP) $(SLIDES) -o $(OUTPUT_DIR)/$(PDF) $(MARP_OPTS)
	@echo "$(GREEN)✓ PDF generated: $(PDF)$(NC)"
	@ls -lh $(PDF)

# Generate HTML
html: check
	@echo "$(BLUE)→ Generating HTML...$(NC)"
	$(MARP) $(SLIDES) -o $(OUTPUT_DIR)/$(HTML) $(MARP_OPTS)
	@echo "$(GREEN)✓ HTML generated: $(HTML)$(NC)"
	@ls -lh $(HTML)

# Generate PowerPoint
pptx: check
	@echo "$(BLUE)→ Generating PowerPoint...$(NC)"
	$(MARP) $(SLIDES) -o $(OUTPUT_DIR)/$(PPTX) $(MARP_OPTS)
	@echo "$(GREEN)✓ PowerPoint generated: $(PPTX)$(NC)"
	@ls -lh $(PPTX)

# Generate PDF with custom theme
pdf-theme: check
	@echo "$(BLUE)→ Generating PDF with custom theme...$(NC)"
	$(MARP) $(SLIDES) -o $(OUTPUT_DIR)/$(PDF) $(THEME_OPTS) $(MARP_OPTS)
	@echo "$(GREEN)✓ PDF with theme generated: $(PDF)$(NC)"
	@ls -lh $(PDF)

# Generate HTML with custom theme
html-theme: check
	@echo "$(BLUE)→ Generating HTML with custom theme...$(NC)"
	$(MARP) $(SLIDES) -o $(OUTPUT_DIR)/$(HTML) $(THEME_OPTS) $(MARP_OPTS)
	@echo "$(GREEN)✓ HTML with theme generated: $(HTML)$(NC)"
	@ls -lh $(HTML)

# Preview slides in browser (live reload)
preview: check
	@echo "$(BLUE)→ Starting preview server...$(NC)"
	@echo "$(YELLOW)Open http://localhost:8080 in your browser$(NC)"
	@echo "$(YELLOW)Press Ctrl+C to stop$(NC)"
	$(MARP) -p $(SLIDES) $(MARP_OPTS)

# Preview with custom theme
preview-theme: check
	@echo "$(BLUE)→ Starting preview server with custom theme...$(NC)"
	@echo "$(YELLOW)Open http://localhost:8080 in your browser$(NC)"
	@echo "$(YELLOW)Press Ctrl+C to stop$(NC)"
	$(MARP) -p $(SLIDES) $(THEME_OPTS) $(MARP_OPTS)

# Watch mode - auto rebuild on changes
watch: check
	@echo "$(BLUE)→ Watching for changes...$(NC)"
	@echo "$(YELLOW)Press Ctrl+C to stop$(NC)"
	$(MARP) -w $(SLIDES) -o $(OUTPUT_DIR)/$(PDF) $(MARP_OPTS)

# Generate all formats
build-all: pdf html pptx
	@echo "$(GREEN)✓ All formats generated!$(NC)"
	@echo "$(BLUE)Files:$(NC)"
	@ls -lh $(PDF) $(HTML) $(PPTX)

# Clean generated files
clean:
	@echo "$(YELLOW)→ Cleaning generated files...$(NC)"
	@rm -f $(PDF) $(HTML) $(PPTX)
	@echo "$(GREEN)✓ Clean complete$(NC)"

# Install Marp CLI
install:
	@echo "$(BLUE)→ Installing Marp CLI...$(NC)"
	@if command -v npm >/dev/null 2>&1; then \
		npm install -g @marp-team/marp-cli; \
		echo "$(GREEN)✓ Marp CLI installed$(NC)"; \
	else \
		echo "$(YELLOW)⚠ npm not found. Please install Node.js first.$(NC)"; \
		exit 1; \
	fi

# Check if Marp is installed
check:
	@if ! command -v $(MARP) >/dev/null 2>&1; then \
		echo "$(YELLOW)⚠ Marp CLI not found!$(NC)"; \
		echo "$(BLUE)Install with: make install$(NC)"; \
		echo "$(BLUE)Or manually: npm install -g @marp-team/marp-cli$(NC)"; \
		exit 1; \
	fi

# Count slides
count:
	@echo "$(BLUE)Slide count:$(NC)"
	@grep -c "^---$$" $(SLIDES) || echo "0"
	@echo "$(BLUE)Total lines:$(NC)"
	@wc -l $(SLIDES)

# Validate markdown syntax
validate:
	@echo "$(BLUE)→ Validating markdown...$(NC)"
	@if command -v markdownlint >/dev/null 2>&1; then \
		markdownlint $(SLIDES); \
		echo "$(GREEN)✓ Markdown valid$(NC)"; \
	else \
		echo "$(YELLOW)⚠ markdownlint not installed$(NC)"; \
		echo "$(BLUE)Install with: npm install -g markdownlint-cli$(NC)"; \
	fi

# Show file info
info:
	@echo "$(BLUE)=== Presentation Info ===$(NC)"
	@echo "$(BLUE)Source:$(NC) $(SLIDES)"
	@echo "$(BLUE)Theme:$(NC) $(THEME)"
	@echo "$(BLUE)Slides:$(NC) $$(grep -c '^---$$' $(SLIDES) 2>/dev/null || echo '0')"
	@echo "$(BLUE)Lines:$(NC) $$(wc -l < $(SLIDES))"
	@echo ""
	@echo "$(BLUE)=== Generated Files ===$(NC)"
	@ls -lh $(PDF) $(HTML) $(PPTX) 2>/dev/null || echo "No files generated yet"

# Open PDF in default viewer
open-pdf: pdf
	@echo "$(BLUE)→ Opening PDF...$(NC)"
	@if [ "$$(uname)" = "Darwin" ]; then \
		open $(PDF); \
	elif [ "$$(uname)" = "Linux" ]; then \
		xdg-open $(PDF); \
	else \
		echo "$(YELLOW)Please open $(PDF) manually$(NC)"; \
	fi

# Open HTML in browser
open-html: html
	@echo "$(BLUE)→ Opening HTML...$(NC)"
	@if [ "$$(uname)" = "Darwin" ]; then \
		open $(HTML); \
	elif [ "$$(uname)" = "Linux" ]; then \
		xdg-open $(HTML); \
	else \
		echo "$(YELLOW)Please open $(HTML) manually$(NC)"; \
	fi

# Export Mermaid diagram to image
diagram: check-mermaid
	@echo "$(BLUE)→ Exporting Mermaid diagram...$(NC)"
	@mkdir -p images
	@$(MERMAID_CLI) -i $(MERMAID_DIAGRAM) -o $(DIAGRAM_OUTPUT) -t dark -b transparent
	@if [ $$? -eq 0 ]; then \
		OUTPUT_FILE=$$(basename $(DIAGRAM_OUTPUT) .png); \
		FOUND_FILE=$$(find images -name "$$OUTPUT_FILE*.png" -type f | head -1); \
		if [ -n "$$FOUND_FILE" ]; then \
			if [ "$$FOUND_FILE" != "$(DIAGRAM_OUTPUT)" ]; then \
				mv "$$FOUND_FILE" $(DIAGRAM_OUTPUT); \
			fi; \
			echo "$(GREEN)✓ Diagram exported: $(DIAGRAM_OUTPUT)$(NC)"; \
			ls -lh $(DIAGRAM_OUTPUT); \
		else \
			echo "$(YELLOW)⚠ File created but not found at expected location$(NC)"; \
		fi; \
	else \
		echo "$(YELLOW)⚠ Failed to export diagram$(NC)"; \
		exit 1; \
	fi

# Install Mermaid CLI
install-mermaid:
	@echo "$(BLUE)→ Installing Mermaid CLI...$(NC)"
	@if command -v npm >/dev/null 2>&1; then \
		npm install -g @mermaid-js/mermaid-cli; \
		echo "$(GREEN)✓ Mermaid CLI installed$(NC)"; \
	else \
		echo "$(YELLOW)⚠ npm not found. Please install Node.js first.$(NC)"; \
		exit 1; \
	fi

# Check if Mermaid CLI is installed
check-mermaid:
	@if ! command -v $(MERMAID_CLI) >/dev/null 2>&1; then \
		echo "$(YELLOW)⚠ Mermaid CLI not found!$(NC)"; \
		echo "$(BLUE)Install with: make install-mermaid$(NC)"; \
		echo "$(BLUE)Or manually: npm install -g @mermaid-js/mermaid-cli$(NC)"; \
		exit 1; \
	fi

# Help
help:
	@echo "$(BLUE)=== ITS Presentation Makefile ===$(NC)"
	@echo ""
	@echo "$(GREEN)Main targets:$(NC)"
	@echo "  make pdf          - Generate PDF (default quality)"
	@echo "  make html         - Generate HTML"
	@echo "  make pptx         - Generate PowerPoint"
	@echo "  make all          - Generate PDF + HTML"
	@echo "  make build-all    - Generate all formats"
	@echo ""
	@echo "$(GREEN)Diagram export:$(NC)"
	@echo "  make diagram      - Export Mermaid diagram from term.md to PNG"
	@echo "  make install-mermaid - Install Mermaid CLI"
	@echo ""
	@echo "$(GREEN)Preview & Development:$(NC)"
	@echo "  make preview      - Start live preview server"
	@echo "  make watch        - Watch and auto-rebuild PDF"
	@echo "  make open-pdf     - Generate and open PDF"
	@echo "  make open-html    - Generate and open HTML"
	@echo ""
	@echo "$(GREEN)With custom theme:$(NC)"
	@echo "  make pdf-theme    - Generate PDF with theme.css"
	@echo "  make html-theme   - Generate HTML with theme.css"
	@echo "  make preview-theme - Preview with custom theme"
	@echo ""
	@echo "$(GREEN)Utilities:$(NC)"
	@echo "  make clean        - Remove generated files"
	@echo "  make install      - Install Marp CLI"
	@echo "  make check        - Check if Marp is installed"
	@echo "  make count        - Count slides"
	@echo "  make info         - Show presentation info"
	@echo "  make validate     - Validate markdown syntax"
	@echo "  make help         - Show this help"
	@echo ""
	@echo "$(YELLOW)Examples:$(NC)"
	@echo "  make pdf && make open-pdf    # Quick preview"
	@echo "  make preview                 # Live development"
	@echo "  make build-all               # Generate all formats"
	@echo "  make diagram                 # Export Mermaid diagram"
	@echo ""
	@echo "$(BLUE)Marp Documentation:$(NC) https://marp.app/"
