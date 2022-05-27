# jobtome
The application exposes three endpoints 
- /api/create-short-url
- /api/{urlShortCode}

The first endpoint is the one that invokes the short url algorith. The endpoint method is an http.Post and a body which is a json of the following format
{
    "original-url": "here-enter-the-original-url"
}
the endpoint produces a short url 
  
The second endpoint is the one that does the actual redirect to the original url. The endpoint method is an http.Get and in the urlShortCode tag the short url must be inserted in order to redirect you to the the original url
 
A version of redis must be installed in order for the application to properly function
DB_URL --> redis:6379
DB_VERSION --> redis
PORT --> 8080

# Using docker
Clone the repo
Docker compose
docker-compose up --build
