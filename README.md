# Exofish Backend

The backend to [exo.fish](https://exo.fish)

## About

A backend needs to be performant, well-organized, well-written, and human-readable. 

Luckily, Go makes all of those requirements easy to meet. 

This code makes up any and all server-side functions related to the function of [exo.fish](https://exo.fish).

Because web requests need to be made in a way that makes the frontend _not_ seem like it's making tons of 
API calls to get various data, Go was an obvious choice. As per [this article](https://medium.com/@avelinorun/go-vs-python-more-request-per-second-1ee0ca7e8681)
which I just found via a Google search, Go's performance in terms of web requests blows Python (which would've been my first choice)
out of the water. 

## Building

`go build`
