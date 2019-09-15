# Fondo-Backend-Proxy

This repository contains the code for a proxy endpoint to the [Unsplash.com](https://unsplash.com/developers) API. This because in the *Unsplash API Guidelines* > *Technical Guidelines* > *Bullet 4* said:

> Your application’s Access Key and Secret Key must remain confidential. This means that they cannot be included in the client or made public. In most cases, this will require proxying the API through your own endpoint to sign the request with your keys.

This Serverless Function is running on [zeit.co](https://zeit.co/) and It uses Go as the programming language to provide the necessary endpoints for the Fondo project.
* Get the today's photos
* Search photos by a query
* Get the download link for a specific photo

This is a public and open-source code used in the project which uses a Development Access Key and Secret Key because I want that users know what is happening when they're using Fondo and invite to the community to report issues, suggestions, and contributions.

For this reason for production deployment, a private repository is used to ensure that the Productions keys are secure and private. 

## Using this API

### Welcome page!

```
https://fondo-backend.clopezr1402.now.sh/api/unsplash/welcome
```

### Get the today photos with a page, per page and order by specified

```
https://fondo-backend.clopezr1402.now.sh/api/unsplash/photos
    ?per_page=2
    &page=1
    &order_by=popular
```

### Get photos by query search with query, page and per page specified

```
https://fondo-backend.clopezr1402.now.sh/api/unsplash/search
    ?query=cats
    &per_page=2
    &page=1
```

### Get the url to download an specific photo from the id

```
https://fondo-backend.clopezr1402.now.sh/api/unsplash/download
    ?id=Wm3EU84mWeg
```