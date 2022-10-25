# gobrewcleanup
Practice Golang by cleaning up brew packages. A bit overkill, but it's sharpening the axe.

# Idea

Get all packages installed by `brew`. Ask about all the manually installed packages if they're still needed or can be removed. Create a list to be removed.

Imagine something like

```
[+] Found 138 manually installed packages

(r)eview (l)ist (e)xport
> r
==> jq: stable 1.6 (bottled), HEAD
Lightweight and flexible command-line JSON processor
https://stedolan.github.io/jq/
/usr/local/Cellar/jq/1.6 (18 files, 1.1MB) *
  Poured from bottle on 2022-02-08 at 21:26:10
From: https://github.com/Homebrew/homebrew-core/blob/HEAD/Formula/jq.rb
License: MIT
==> Dependencies
Required: oniguruma ✔

(k)eep (r)emove
> r

Uninstalling /usr/local/Cellar/w3m/0.5.3_7... (28 files, 1.9MB)
```

Essentially a `for each package in $(brew leaves); do prompt && brew remove $package; done`
