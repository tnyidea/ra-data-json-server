# ra-data-json-server

- getList
    - GET http://my.api.url/posts?_sort=title&_order=ASC&_start=0&_end=24&title=bar
- getOne
    - GET http://my.api.url/posts/123
- getMany
    - GET http://my.api.url/posts?id=123&id=456&id=789
- getManyReference
    - GET http://my.api.url/posts?author_id=345
- create
    - POST http://my.api.url/posts
- update
    - PUT http://my.api.url/posts/123
- updateMany
    - PUT http://my.api.url/posts/123, PUT http://my.api.url/posts/456, PUT http://my.api.url/posts/789
- delete
    - DELETE http://my.api.url/posts/123


- Note: The JSON Server REST Data Provider expects the API to include a X-Total-Count header in the response to getList and getManyReference calls. The value must be the total number of resources in the collection. This allows react-admin to know how many pages of resources there are in total, and build the pagination controls.
```
X-Total-Count: 319
```

- If your API is on another domain as the JS code, you'll need to whitelist this header with an Access-Control-Expose-Headers CORS  header.
```
Access-Control-Expose-Headers: X-Total-Count
```