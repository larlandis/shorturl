# hash

```go
import "github.com/larlandis/shorturl/internal/pkg/hash"
```

## Index

- [type Hash](<#type-hash>)
  - [func New(storage Storage) *Hash](<#func-new>)
  - [func (s Hash) CreateNewHash(ctx context.Context, input string, length uint) (string, error)](<#func-hash-createnewhash>)
  - [func (s Hash) GetFromHash(ctx context.Context, hash string) (string, error)](<#func-hash-getfromhash>)
- [type Storage](<#type-storage>)


## type [Hash](<https://github.com/larlandis/shorturl/blob/main/api/internal/pkg/hash/shorturl.go#L11-L13>)

```go
type Hash struct {
    // contains filtered or unexported fields
}
```

### func [New](<https://github.com/larlandis/shorturl/blob/main/api/internal/pkg/hash/shorturl.go#L43>)

```go
func New(storage Storage) *Hash
```

### func \(Hash\) [CreateNewHash](<https://github.com/larlandis/shorturl/blob/main/api/internal/pkg/hash/shorturl.go#L21>)

```go
func (s Hash) CreateNewHash(ctx context.Context, input string, length uint) (string, error)
```

CreateNewHash creates and saves a new short hash from a given string

### func \(Hash\) [GetFromHash](<https://github.com/larlandis/shorturl/blob/main/api/internal/pkg/hash/shorturl.go#L30>)

```go
func (s Hash) GetFromHash(ctx context.Context, hash string) (string, error)
```

GetFromHash searches and returns a saved string from a hash

## type [Storage](<https://github.com/larlandis/shorturl/blob/main/api/internal/pkg/hash/shorturl.go#L14-L17>)

```go
type Storage interface {
    SavePair(ctx context.Context, input string, short string) error
    Search(ctx context.Context, short string) (url string, err error)
}
```