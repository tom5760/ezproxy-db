+++
summary = """
Database of URLs from institutions using proxy servers, like
[EZproxy](http://en.wikipedia.org/wiki/EZproxy).
"""
+++

Collection of proxy URLs from various libraries and institutions using
software, like [EZproxy][ezproxy].  Can be used with the [Google Chrome
Extension][chrome] or [Mozilla Firefox Extension][firefox].

[ezproxy]: https://en.wikipedia.org/wiki/EZproxy
[chrome]:  https://chrome.google.com/webstore/detail/gfhnhcbpnnnlefhobdnmhenofhfnnfhi
[firefox]: https://addons.mozilla.org/en-US/firefox/addon/ezproxy-redirect

This database is also accessible in [JSON format][json].  The
[code for this site][github-site] and code for the
[browser extensions][github-browser] are available on GitHub.

[json]:           /proxies.json
[github-site]:    https://github.com/tom5760/ezproxy-db
[github-browser]: https://github.com/tom5760/chrome-ezproxy

## Contributing

To contribute your institution's proxy URL, or to update an existing entry,
[submit an issue][issue], or you can
[edit the database directly and create a pull request][pr].
Any questions or comments can be posted [on the discussions board][board].

[issue]: https://github.com/tom5760/ezproxy-db/issues/new/choose
[pr]:    https://github.com/tom5760/ezproxy-db/edit/main/static/proxies.json

### What happened to ezproxy-db.appspot.com?

[I][me] decided to update and simplify the site.  We didn't need a specialized
app to manage a simple set of URLs.  Instead, the database is just a single
file, and this site generated statically from it.  Keeping everything in a
normal Git repository gets us a bunch of features for free, like tracking
changes to the database, and lets me more easily add co-maintainers.

The [previous version][appengine] is still running for now.  Requests to the
HTML index page are redirected here.  Requests for the [proxies.json][json]
file are proxied through, in order to support upgrading browser extensions to
use the new URL.  This will be supported at least through the end of 2022.

Feel free to [discuss this on the discussion board][board].

[me]:        https://github.com/tom5760/
[appengine]: https://github.com/tom5760/ezproxy-db/tree/appengine
[board]:     https://github.com/tom5760/ezproxy-db/discussions
