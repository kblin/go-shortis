# Host-it-yourself URL shortener

A little host-it-yourself URL shortener.


## Usage

First, set up the database.

```
shortis init
```

To add a new shortened URL, use
```
shortis add shortId https://url.to.point/to/here
```

To delete a shortened URL, use
```
shortis remove shortId
```

To update an existing URL, use
```
shortis update shortId https://new.url.to.point/to/here/again
```

To list all available aliases and URLS, use
```
shortis list
```

To serve the URL shortener, use
```
shortis serve
```


## License
Licensed under the Apache 2.0 license, see [`LICENSE`](LICENSE) for details.
