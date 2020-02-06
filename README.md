# Generator

A Go package to generate data for a Virtual Machine metadata service. 

## Installation

Retrieve package:

```bash
go get -u gitlab.com/jmarhee/generatorpkg
```

Include:

```
import (
	generatorpkg "gitlab.com/jmarhee/generatorpkg"
)
```

## Data Types Generated

This package can generate MAC Addresses:

```go
mac := generatorpkg.NewMacAddr()
```

and UUIDs:

```go
uuid := generatorpkg.NewUuid()
```

