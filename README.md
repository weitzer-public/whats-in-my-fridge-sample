# What's in my Fridge? 🥑

This web application suggests a simple recipe based on a list of ingredients you have at home.

## How to run locally

1.  **Prerequisites:** Make sure you have Go installed on your system.

2.  **Download the files:**
    You can clone this repository or download the files directly.

3.  **Initialize the module:**
    Open your terminal, navigate to the project directory, and run the following command to download the dependencies:
    ```bash
    go mod tidy
    ```

4.  **Run the application:**
    Execute the following command to start the web server:
    ```bash
    go run .
    ```

5.  **Access the application:**
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

2.  **Run the Docker container:**
    After the image is built, run the following command to start the container:
    ```bash
    docker run -p 8080:8080 what-fridge
    ```

    You can then access the application at [http://localhost:8080](http://localhost:8080).
