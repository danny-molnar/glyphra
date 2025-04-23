#!/bin/bash
SUMMARY_FILE="docs/_posts/$(date +%F)-summary.md"

echo "---" > "$SUMMARY_FILE"
echo "title: \"Terraform Plan Summary for $(date +%F)\"" >> "$SUMMARY_FILE"
echo "date: $(date +%F)" >> "$SUMMARY_FILE"
echo "layout: post" >> "$SUMMARY_FILE"
echo "tags: [glyphra, terraform, ai]" >> "$SUMMARY_FILE"
echo "---" >> "$SUMMARY_FILE"
echo "" >> "$SUMMARY_FILE"

echo "$1" >> "$SUMMARY_FILE"

echo "✅ Saved blog post to $SUMMARY_FILE"
