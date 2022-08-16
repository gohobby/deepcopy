window.BENCHMARK_DATA = {
  "lastUpdate": 1660615756914,
  "repoUrl": "https://github.com/gohobby/deepcopy",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "27848278+hgtgh@users.noreply.github.com",
            "name": "Hugo-T",
            "username": "hgtgh"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7797a8c49add54c649be7b9afd3b036834566111",
          "message": "Fix benchmark.yml",
          "timestamp": "2022-08-16T04:08:52+02:00",
          "tree_id": "0d427be9154adbd0dde80e830e2b223dfe7a5ae9",
          "url": "https://github.com/gohobby/deepcopy/commit/7797a8c49add54c649be7b9afd3b036834566111"
        },
        "date": 1660615756483,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCopyableMap",
            "value": 1412,
            "unit": "ns/op\t    1248 B/op\t      16 allocs/op",
            "extra": "781101 times\n2 procs"
          },
          {
            "name": "BenchmarkDeepCopy_Map",
            "value": 1406,
            "unit": "ns/op\t    1248 B/op\t      16 allocs/op",
            "extra": "847802 times\n2 procs"
          }
        ]
      }
    ]
  }
}