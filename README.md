Library Proxy URL Database
==========================

* Author: Tom Wambold <tom@wambold.dev>
* Copyright © 2026 Tom Wambold
* URL: https://libproxy-db.org/

Collection of proxy URLs from various libraries and institutions using
software, like [EZproxy][ezproxy].

[ezproxy]: https://en.wikipedia.org/wiki/EZproxy

This database is also accessible in [JSON format][json].  The
[code for this site][github-site] and code for the
[browser extensions][extension-github] are available on GitHub.

[json]:           /proxies.json
[github-site]:    https://github.com/tom5760/ezproxy-db

## Browser Extensions

There are several browser extensions using this database.  Let us know if you
want to add yours to this list!

* [EZProxy Redirect][extension-github] ([Chrome][extension-chrome], [Firefox][extension-firefox]) by @tom5760
* [ScholarKey](https://github.com/ezraiiiiiiiiiiii/ScholarKey) ([Firefox](https://addons.mozilla.org/en-GB/firefox/addon/scholarkey/)) by @ezraiiiiiiiiiiii

[extension-github]: https://github.com/tom5760/chrome-ezproxy
[extension-chrome]:  https://chromewebstore.google.com/detail/ezproxy-redirect/gfhnhcbpnnnlefhobdnmhenofhfnnfhi
[extension-firefox]: https://addons.mozilla.org/addon/firefox-ezproxy-redirect/

## Contributing

To contribute your institution's proxy URL, or to update an existing entry,
[edit the database and create a pull request][pr]. Any questions or comments
can be posted [on the discussions board][board].

[pr]: https://github.com/tom5760/ezproxy-db/edit/main/static/proxies.json
[board]: https://github.com/tom5760/ezproxy-db/discussions

## What Changed?

This used to be a Python app hosted on Google App Engine.  The code for that is
available on the [`appengine` branch][appengine].  I decided to simplify this
and just generate a static site with [Hugo][hugo].  The database is also just a
[single static JSON file][json]. This makes the site super-easy to host
anywhere, and can be served quickly.

[appengine]: https://github.com/tom5760/ezproxy-db/tree/appengine
[hugo]:      https://gohugo.io/

Also, will be easier to use GitHub's collaboration features to track and vet
changes to the database and add co-maintainers, without having to implement
that with code. Users can use [GitHub's built-in editor][pr] to edit the
database and submit a pull request, without using Git directly.
