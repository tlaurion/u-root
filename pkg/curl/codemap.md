# pkg/curl/

## Responsibility

Generic file fetching library supporting multiple URL schemes (HTTP, TFTP,
local files). Provides a pluggable scheme registry, caching/non-caching readers,
lazy fetching, and automatic retries with backoff.

## Design Patterns

- **Strategy pattern**: `FileScheme` interface with per-scheme implementations
- **Registry pattern**: `Schemes` map of scheme → `FileScheme`
- **Decorator pattern**: `SchemeWithRetries` wraps any `FileScheme` with retry logic
- **`io.ReaderAt` vs `io.Reader`**: dual fetch API (cached vs streaming)
- **Lazy evaluation**: `LazyFetch` delays fetching until the first read
- **Composable retry predicates**: `DoRetry` functions composed with `RetryOr`

## Integration Points

- `cmds/boot/*`: used for fetching kernels and initrds over network
- `cmds/core/curl`: CLI HTTP/TFTP/file download tool
- `pkg/dhclient`: uses for network boot URL handling

## Key Types/Functions

- `FileScheme` interface — Fetch(context, *url.URL) and FetchWithoutCache
- `Schemes` map — Registry of URL scheme → FileScheme
- `DefaultSchemes` — Default HTTP, TFTP, and local file support
- `File` interface — Stringer + URL() for fetched file reference
- `FileWithCache` — io.ReaderAt + File
- `FileWithoutCache` — io.Reader + File
- `LazyFetch(u *url.URL) (FileWithCache, error)` — Lazily fetch a URL
- `SchemeWithRetries` — Retry wrapper with backoff
- `HTTPClient` — HTTP FileScheme implementation
- `TFTPClient` — TFTP FileScheme implementation
- `LocalFileClient` — Local file scheme implementation
- `NewHTTPClient(c *http.Client) *HTTPClient`
- `NewTFTPClient(opts ...tftp.ClientOpt) FileScheme`
- `RetryHTTP(u *url.URL, err error) bool` — HTTP-specific retry predicate
- `RetryTFTP(u *url.URL, err error) bool` — TFTP-specific retry predicate
- `RetryConnectErrors(u *url.URL, err error) bool` — Connect error retry
- `RetryTemporaryNetworkErrors(u *url.URL, err error) bool` — Temp network error retry
