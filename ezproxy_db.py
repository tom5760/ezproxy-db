import sys
import json
import os.path

from google.appengine.api import mail
from google.appengine.api import users
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
            'proxies': Proxy.all().filter('approved =', True).order('name'),
            'login_url': users.create_login_url('/'),
            'logout_url': users.create_logout_url('/'),
            'is_logged_in': user is not None,
            'is_admin': users.is_current_user_admin(),
        }

        if users.is_current_user_admin():
            context['moderated_proxies'] = (
                    Proxy.all().filter('approved =', False).order('name'))

        self.render_response('index.html', context)

class AddHandler(BaseHandler):
    def post(self):
        user = users.get_current_user()

        if user is None:
            self.redirect('/')

        proxy = Proxy(
            name=self.request.get('name'),
            url=self.request.get('url'),
            approved=users.is_current_user_admin(),
        )
        proxy.put()

        if not proxy.approved:
            message = 'New EZProxy URL: <a href="{1}">{0} - {1}</a>'.format(
                    proxy.name, proxy.url)
            mail.send_mail_to_admins('no-reply@ezproxy-db.appspotmail.com',
                    'EZProxy DB Moderation Request', message)

            self.render_response('addproxy.html', {})
        else:
            self.redirect('/')

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

        self.redirect('/')

class JSONHandler(BaseHandler):
    def get(self):
        self.response.headers['Content-Type'] = 'application/json'

        proxies = [{'name': p.name, 'url': p.url}
                for p in Proxy.all().filter('approved =', True).order('name')]

        json.dump(proxies, self.response, separators=(',', ':'))

app = webapp2.WSGIApplication([
        webapp2.Route('/', 'ezproxy_db.MainHandler', 'main'),
        webapp2.Route('/addproxy', 'ezproxy_db.AddHandler', 'add'),
        webapp2.Route('/editproxy', 'ezproxy_db.EditHandler', 'edit'),
        webapp2.Route('/proxies.json', 'ezproxy_db.JSONHandler', 'json'),
    ], debug=True)
