# act cache

The *act cache* allows the Github Actions to cache their artifacts. This package is copied from the the [Gitea arc-runner implementation](https://gitea.com/gitea/act_runner/src/branch/main/internal/app/artifactcache) to provide a cache before Gitea contributing back.

## installation

Either download a prebuilt release from the Release page or build using `go build .`

## usage

### Act config
We need to let act to know the Cache Server URL and use which token. There are two ways to do so (choose one).

1. Ensure you add the following configuration to your `~/.actrc` file:
```
--env ACTIONS_CACHE_URL=http://localhost:8111/
--env ACTIONS_RUNTIME_URL=http://localhost:8111/

# run act cache
act_cache > cache.log &

# just run act
act {event}
```

2. Run act with env file (for example, put the file in project root) 
```
# ./act/cache.env
ACTIONS_CACHE_URL=http://localhost:8080/
ACTIONS_RUNTIME_URL=http://localhost:8080/

# run act cache
act_cache > cache.log &

# run act with .env file
act {event} --env-file act/cache.env
```
