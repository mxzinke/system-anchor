# System Anchors

Linking an app or something different and the end-user will get redirected to the os specific URL.

## Example Config

```yaml
directions:
  # redirecting per OS
  example:
    ios: "https://google.com"
    android: "https://medium.com"

  # just another example
  another:
    ios: "https://apps.apple.com/us"
    android: "https://play.google.com/store"
```

The key is `example` and you can access the redirection by use `/example` of the server.