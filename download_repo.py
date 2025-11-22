import urllib.request
import zipfile
import os
import sys

url = "https://github.com/rezkypra/knowledge-test-rezky/archive/refs/heads/main.zip"
zip_path = "backend.zip"
extract_path = "backend_repo"

print(f"Starting download from {url}...", flush=True)

try:
    # Use a different user agent
    headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'}
    req = urllib.request.Request(url, headers=headers)
    
    with urllib.request.urlopen(req, timeout=60) as response, open(zip_path, 'wb') as out_file:
        data = response.read()
        out_file.write(data)
        print(f"Download finished. Size: {len(data)} bytes", flush=True)

    if os.path.exists(zip_path):
        print(f"Extracting to {extract_path}...", flush=True)
        with zipfile.ZipFile(zip_path, 'r') as zip_ref:
            zip_ref.extractall(extract_path)
        print("Extraction complete.", flush=True)
    else:
        print("Error: File not found after download.", flush=True)

except Exception as e:
    print(f"Error occurred: {e}", flush=True)
    sys.exit(1)
