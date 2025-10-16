# Acetlisto

A demo API for shopping list


## Development

Create a GitHub codespace and run the following command:

```
air
```

## Requests

### Create an item

```
curl -d '{"Name":"Oil", "Description":"Should be nice!"}' -H "Content-Type: application/json" -X POST http://localhost:3000/items/
```

### List all items

```
curl -v http://localhost:3000/items/
```

### Get a single item

```
curl -v http://localhost:3000/items/{ID}
```