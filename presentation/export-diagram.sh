#!/bin/bash
# Script to export Mermaid diagram to image
# Usage: ./export-diagram.sh [input.md] [output.png]

INPUT="${1:-term.md}"
OUTPUT="${2:-images/architectural-characteristics-radar.png}"

# Check if @mermaid-js/mermaid-cli is installed
if ! command -v mmdc &> /dev/null; then
    echo "⚠️  Mermaid CLI not found!"
    echo "📦 Installing @mermaid-js/mermaid-cli..."
    npm install -g @mermaid-js/mermaid-cli
fi

# Create output directory if it doesn't exist
mkdir -p "$(dirname "$OUTPUT")"

echo "🔄 Converting $INPUT to $OUTPUT..."
mmdc -i "$INPUT" -o "$OUTPUT" -t dark -b transparent

if [ $? -eq 0 ]; then
    echo "✅ Successfully exported to $OUTPUT"
    echo "📊 File size: $(ls -lh "$OUTPUT" | awk '{print $5}')"
else
    echo "❌ Failed to export diagram"
    exit 1
fi

