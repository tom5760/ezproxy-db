import json
import logging

from google.appengine.api import mail, memcache, users
from google.appengine.ext import db
from google.appengine.ext.db import BadKeyError

import webapp2
from webapp2_extras import jinja2

class Proxy(db.Model):
    name = db.StringProperty(required=True)
    url = db.LinkProperty(required=True)
    approved = db.BooleanProperty(default=False, required=True)

class BaseHandler(webapp2.RequestHandler):
    @webapp2.cached_property
    def jinja2(self):
        return jinja2.get_jinja2(app=self.app)

    def render_response(self, template, context):
        self.response.write(self.jinja2.render_template(template, **context))

class MainHandler(BaseHandler):
    def get(self):
        user = users.get_current_user()
        context = {
            'login_url': users.create_login_url('/'),
            'logout_url': users.create_logout_url('/'),
            'is_logged_in': user is not None,
            'is_admin': users.is_current_user_admin(),
        }

        result = memcache.get_multi(['proxies', 'moderated_proxies'])

        if 'proxies' in result:
            logging.info('proxies cache hit')
            context['proxies'] = result['proxies']
        else:
            logging.info('proxies cache miss')
            context['proxies'] = Proxy.all().filter('approved =', True).order('name')

        if 'moderated_proxies' in result:
            logging.info('moderated proxies cache hit')
            context['moderated_proxies'] = result['moderated_proxies']
        else:
            logging.info('moderated proxies cache miss')
            context['moderated_proxies'] = Proxy.all().filter('approved =', False).order('name')

        memcache.add_multi({
            'proxies': context['proxies'],
            'moderated_proxies': context['moderated_proxies'],
        })

        self.render_response('index.html', context)

class AddHandler(BaseHandler):
    def post(self):
        user = users.get_current_user()

        if user is None:
            self.redirect('/')

        name = self.request.get('name')
        url = self.request.get('url')

        dup_name = Proxy.all().filter('name =', name)
        dup_url = Proxy.all().filter('url =', url)

        error = False
        context = {
            'msg': 'Thanks for your addition!  Your URL will be checked by an administrator soon.',
        }

        if url.find('$@') == -1:
            error = True
            context['msg'] = 'Your URL does not include the token "$@".  Please see the other URLs in the database for examples.'

        if dup_name.count() > 0:
            error = True
            context['msg'] = 'A proxy with that name already exists in the database, please try again.'

        if dup_url.count() > 0:
            if error:
                context['msg'] = 'A proxy with that name and url already exists in the database, please try again.'
            else:
                error = True
                context['msg'] = 'A proxy with that url already exists in the database, please try again.'

        if not error:
            proxy = Proxy(
                name=self.request.get('name'),
                url=self.request.get('url'),
                approved=users.is_current_user_admin(),
            )
            proxy.put()

            if not proxy.approved:
                memcache.delete('moderated_proxies')
                message = 'New EZProxy URL: {0}: {1}\n\nhttp://ezproxy-db.appspot.com'.format(
                        proxy.name, proxy.url)
                mail.send_mail_to_admins('no-reply@ezproxy-db.appspotmail.com',
                        'EZProxy DB Moderation Request', message)
            else:
                memcache.delete('proxies')
                return self.redirect('/')

        self.render_response('addproxy.html', context)

class EditHandler(BaseHandler):
    def post(self):
        if not users.is_current_user_admin():
            self.redirect('/')

        try:
            proxy = Proxy.get(self.request.get('id'))
        except BadKeyError:
            self.redirect('/')
            return

        action = self.request.get('action').lower()
        if action == 'add':
            proxy.approved = True
            proxy.put()
        elif action == 'delete':
            proxy.delete()
        elif action == 'remove':
            proxy.approved = False
            proxy.put()

        memcache.delete_multi(['proxies', 'moderated_proxies'])
        self.redirect('/')

class JSONHandler(BaseHandler):
    def get(self):
        self.response.headers['Content-Type'] = 'application/json'

        proxies = memcache.get('proxies')
        if proxies is None:
            proxies = Proxy.all().filter('approved =', True).order('name')
            memcache.set('proxies', proxies)

        data = [{'name': p.name, 'url': p.url} for p in proxies]
        json.dump(data, self.response, separators=(',', ':'))

app = webapp2.WSGIApplication([
        webapp2.Route('/', 'ezproxy_db.MainHandler', 'main'),
        webapp2.Route('/addproxy', 'ezproxy_db.AddHandler', 'add'),
        webapp2.Route('/editproxy', 'ezproxy_db.EditHandler', 'edit'),
        webapp2.Route('/proxies.json', 'ezproxy_db.JSONHandler', 'json'),
    ], debug=True)
