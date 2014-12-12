# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import os
import sys

import endpoints
import webapp2

from components import utils

import buildbucket

APP_DIR = os.path.dirname(os.path.abspath(__file__))
sys.path.append(os.path.join(APP_DIR, 'third_party'))

# TODO(nodir): include ui and tasks
handlers = []


def create_html_app():
  """Returns WSGI app that serves HTML pages."""
  routes = [
  ]
  return webapp2.WSGIApplication(routes, debug=utils.is_local_dev_server())


def create_endpoints_app():
  """Returns WSGI app that serves cloud endpoints requests."""
  apis = [
      buildbucket.BuildBucketApi,
  ]
  return endpoints.api_server(apis)


def initialize_apps():
  """Bootstraps the global state and creates WSGI applications."""
  return create_html_app(), create_endpoints_app()
