# Case Study Backend

This case study comprises of 2 challenges.\
For the second challenge, i build RESTful API for managing and uploading movies using Go, MySQL, and MongoDB (GridFS).\
Here are the description of the API and a little bit of guide to run the API:

## Features
- Upload video files to MongoDB
- Upload movie metadata to MySQL
- Get, update, and search movies

## Here is the application flow of how this API is gonna be used by Front-End
- Admin is given a form to upload the movie file video. That form consists of several fields of movie metadata such as title, description, duration, artists, genres. There is button in the form to upload the video. When hit, the video is uploaded to MongoDB through an endpoint. This endpoint returns the id of the file from the MongoDB and the video filename, etc. The video filename is gonna be used to auto-fill the title field in the form. After admin customizing the right value for each field, then the form is submitted and all of those movie metadatas are stored to MySQL through another endpoint.
- Admin can update movie. It triggers a form similar to the 'upload movie' scenario.
- Admin can display movies. This page will use the endpoint that uses pagination.
- Admin can search movie by title, description, artists, genres. 

## How to start

### 1. Start MongoDB (Docker)
```bash
docker run -d -p 27017:27017 --name mongodb mongo:6
```
### 2. Start MySQL Server

### 3. Import the DB to your MySQL Server

### 4. Open the API documentation with the link below to help you test the API
https://documenter.getpostman.com/view/32572193/2sB2qcCLe3
