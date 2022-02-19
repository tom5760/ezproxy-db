Library Proxy URL Database
==========================

* Author: Tom Wambold <tom5760@gmail.com>
* Copyright © 2022 Tom Wambold
* URL: https://libproxy-db.org/

Collection of proxy URLs from various libraries and institutions using
software, like [EZproxy][ezproxy].  Can be used with the [Google Chrome
Extension][chrome] or [Mozilla Firefox Extension][firefox].

[ezproxy]: https://en.wikipedia.org/wiki/EZproxy
[chrome]:  https://chrome.google.com/webstore/detail/gfhnhcbpnnnlefhobdnmhenofhfnnfhi
[firefox]: https://addons.mozilla.org/en-US/firefox/addon/ezproxy-redirect

This database is also accessible in [JSON format][json].  Also, the code for
the [browser extensions][github-browser] is available on GitHub.

[json]:           https://github.com/tom5760/ezproxy-db/blob/main/static/proxies.json
[github-browser]: https://github.com/tom5760/chrome-ezproxy

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
