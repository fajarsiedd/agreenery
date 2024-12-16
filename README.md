# **Agreenery**

## **Introduction**

Agreenery is an agricultural application to help farmers in planting and caring for plants so that they can grow healthily.

## **Prerequisite**
* **Go:** 1.23 or higher
* **Database:** MySQL 8.0 or higher
* **API Keys:**
    * `GOOGLE_MAPS_API_KEY` for use of Google Maps API. [Get your API key](https://mapsplatform.google.com/)    
    * `GEMINI_API_KEY` for use of Gen AI Gemini 1.5 Flash. [Get your API key](https://ai.google.dev/gemini-api/docs?gad_source=1&gclid=CjwKCAiAxea5BhBeEiwAh4t5K-uLKpnHMmmUmfdgAgRQG-WXsX2AP1N9CETrOASezuTErrrTuhGiHBoCTaUQAvD_BwE&hl=id)
* **Tools:**
    * `Docker` (optional, for local development)

## **Installation**

To install Agreenery, follow these steps:

1. Clone the repository: 
    ```bash
    git clone https://github.com/fajarsiedd/agreenery.git
2. Navigate to the project directory:
    ```bash
    cd agreenery
3. Install dependencies:
    ```bash
    go get .
3. Add ENV file:
    ```bash
    copy .env.example .env
4. Modify ENV file:
    ```bash
    # DB CONFIG
    DB_HOST="localhost"
    DB_PORT="3306"
    DB_USERNAME="root"
    DB_PASSWORD=""
    DB_NAME="agreenery"
    
    # JWT
    JWT_SECRET_KEY="your_jwt_secret"
    
    # API KEYS
    GOOGLE_MAPS_API_KEY="your_api_key"
    GEMINI_API_KEY="your_api_key"
6. Run tests:
    ```bash
    go test ./...
7. Start the project:
    ```bash
    go run main.go


## **Docker**

If you'd like to use docker to run Agreenery, here are some guidelines:

1. Make sure you have followed Agreenery installation steps 1-5
2. Open your Docker Desktop or turn on your Docker Engine. If you haven't installed Docker yet. [Follow these steps](https://docs.docker.com/desktop/setup/install/windows-install/)
3. Update ENV file:
    ```bash
    #DB CONFIG
    DB_HOST="host.docker.internal" //Change this line
    ...
4. Build docker image:
    ```bash
    docker build -t agreenery:1.0.0 .
5. Run container:
    ```bash
    docker run -itd --name agreenery-app -p 1323:1323 agreenery:1.0.0
6. Check the logs:
    ```bash
    docker logs carbon-tracker
7. Now you can test the API using [Postman](https://www.postman.com/) ✨

## **Contact Me**
**Email:** fajarsidiq999@gmail.com
    
Happy Coding ❤