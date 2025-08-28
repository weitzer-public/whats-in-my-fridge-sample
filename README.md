# What's in my Fridge? 🥑

This web application suggests a simple recipe based on a list of ingredients you have at home.

## How to run locally

1.  **Prerequisites:** Make sure you have Go installed on your system.

2.  **Download the files:**
    You can clone this repository or download the files directly.

3.  **Configure the Application:**
    - Open the `config.json` file.
    - Replace the placeholder values for `gcp_project_id` and `gemini_api_key_secret_name` with your actual Google Cloud project ID and the full resource name of your Secret Manager secret.

4.  **Authenticate with Google Cloud:**
    Use the following command to authenticate your local environment with Google Cloud. This will allow the application to access Secret Manager.
    ```bash
    gcloud auth application-default login
    ```

5.  **Initialize the module:**
    Open your terminal, navigate to the project directory, and run the following command to download the dependencies:
    ```bash
    go mod tidy
    ```

6.  **Run the application:**
    Execute the following command to start the web server:
    ```bash
    go run .
    ```

7.  **Access the application:**
    Open your web browser and go to [http://localhost:8080](http://localhost:8080).

## How to use

1.  **Enter your ingredients:**
    In the text area, type the ingredients you have on hand. It's best to put one ingredient per line.

2.  **Suggest a recipe:**
    Click the "Suggest a Recipe" button.

3.  **View the recipe:**
    The application will display a suggested recipe based on your ingredients.

## Running with Docker

This project includes a `Dockerfile` to run the application in a container.

1.  **Build the Docker image:**
    Open your terminal in the project root and run the following command:
    ```bash
    docker build -t what-fridge .
    ```

2.  **Configure Authentication:**
    When running inside a Docker container, the application needs a way to authenticate with Google Cloud to use the Gemini API. You can do this in one of two ways:

    **Method 1: Using Application Default Credentials (if you have `gcloud` installed)**

    If you have the `gcloud` CLI installed on your machine, you can use your local credentials.

    a. **Login with `gcloud`:**
    ```bash
    gcloud auth application-default login
    ```

    b. **Run the container with your credentials mounted:**
    ```bash
    docker run -p 8080:8080 \
      -v ~/.config/gcloud/application_default_credentials.json:/app/credentials.json \
      -e GOOGLE_APPLICATION_CREDENTIALS=/app/credentials.json \
      what-fridge
    ```
    *(Note for Windows users: Replace `~/.config/gcloud/...` with `%APPDATA%\gcloud\application_default_credentials.json`)*

    **Method 2: Using a Service Account Key (if you don't have `gcloud`)**

    If you don't have `gcloud` installed, you can use a service account key file.

    a. **Create a Service Account:**
       - In the Google Cloud Console, go to the "Service Accounts" page for your project.
       - Click "Create Service Account".
       - Give it a name (e.g., "what-fridge-sa").
       - Grant it the "Vertex AI User" role.
       - Click "Done".

    b. **Create a Service Account Key:**
       - Click on the newly created service account.
       - Go to the "Keys" tab.
       - Click "Add Key" -> "Create new key".
       - Select "JSON" as the key type and click "Create".
       - A JSON file will be downloaded to your computer. **Treat this file like a password and keep it secure.**

    c. **Run the container with the key file mounted:**
       Move the downloaded JSON key file to your project directory and run the following command (replace `[PATH_TO_YOUR_KEY_FILE].json` with the actual path to your key file):
       ```bash
       docker run -p 8080:8080 \
         -v $(pwd)/[PATH_TO_YOUR_KEY_FILE].json:/app/credentials.json \
         -e GOOGLE_APPLICATION_CREDENTIALS=/app/credentials.json \
         what-fridge
       ```

3.  **Access the application:**
    Open your web browser and go to [http://localhost:8080](http://localhost:8080).

