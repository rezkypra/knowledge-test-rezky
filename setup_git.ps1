$ErrorActionPreference = "Stop"
Write-Host "Setting up Git..."

# Remove origin if exists (ignore error)
try { git remote remove origin } catch {}

# Add remote
Write-Host "Adding remote..."
git remote add origin "https://github.com/rezkypra/knowledge-test-rezky.git"

# Pull
Write-Host "Pulling changes..."
git pull origin main --allow-unrelated-histories

# Push
Write-Host "Pushing changes..."
git push -u origin main
