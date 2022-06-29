# URLShortener
Created url shortener REST APIs using Golang with gin package.

## Algorithm
* The mapping between the original Url and the generated short Url is stored in a map
* Hashing the long url with SHA256 and using base62 for encoding it

## Requests
* POST - http://localhost:9000/create_url request body:
```
{
  "long_url": "<long_url>"
}
```
  response:
```
{
  "short_url": "<short_url>"
}
```

* GET - <short_url> redirects to the original url
