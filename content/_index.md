+++
summary = """
Database of URLs from institutions using proxy servers, like
[EZproxy](http://en.wikipedia.org/wiki/EZproxy).
"""
+++

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
* [ScholarKey](https://github.com/ezraiiiiiiiiiiii/ScholarKey) ([Firefox](https://addons.mozilla.org/en-GB/firefox/addon/scholarkey/) by @ezraiiiiiiiiiiii

[extension-github]: https://github.com/tom5760/chrome-ezproxy
[extension-chrome]:  https://chromewebstore.google.com/detail/ezproxy-redirect/gfhnhcbpnnnlefhobdnmhenofhfnnfhi
[extension-firefox]: https://addons.mozilla.org/addon/firefox-ezproxy-redirect/

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
