# Soal Kedua - Project Solution & Integration

This repository contains the solution for "Soal Kedua" and instructions to integrate it with "Soal Pertama".

## Structure

*   `web-dashboard/`: Next.js 14 application (Soal Kedua).
*   `mobile-app/`: Flutter application (Soal Kedua).
*   `Dockerfile`: Docker configuration for the Backend (Soal Pertama).
*   `docker-compose.yml`: Orchestration for Backend, Database, and Web Dashboard.

## Integration Instructions

To merge this with your existing "Soal Pertama" repository:

1.  **Copy Files:**
    Copy the following files/folders to the **root** of your `knowledge-test-rezky` repository:
    *   `web-dashboard/`
    *   `mobile-app/`
    *   `Dockerfile` (The one in the root, for the backend)
    *   `docker-compose.yml`

2.  **Verify Backend Structure:**
    Ensure your backend code (Go files) is in the root or matches the `Dockerfile` build path.

3.  **Run Everything:**
    Open a terminal in the root of your repository and run:
    ```bash
    docker-compose up --build
    ```

    This will start:
    *   **Backend:** `http://localhost:8080`
    *   **Database:** PostgreSQL on port `5432`
    *   **Web Dashboard:** `http://localhost:3000`

## Features

### Web Dashboard
*   **Dashboard:** Stats cards and Student Table.
*   **Mata Kuliah:** Course management view.
*   **Data Akademik:** Academic performance overview.
*   **API:** Connects to `http://localhost:8080` (Configured in `docker-compose.yml`).

### Mobile App
*   **Login:** Clean UI (Version 2).
*   **Home:** "Halo Dunia" screen.
*   **Dark Mode:** System adaptive.
*   **Run:** `cd mobile-app && flutter run`

## Git Push
After copying the files, commit and push to your repository:

```bash
git add .
git commit -m "Add Soal Kedua (Web, Mobile) and Docker Integration"
git push origin main
```
