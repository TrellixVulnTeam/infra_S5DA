# Copyright (c) 2008-2009 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""AppEngine scripts to manage the chromium tree status.

   Logged in people with a @chromium.org email can change the status that
   appears on the waterfall page. The previous status are kept in the
   database and the last 100 topics.
"""

from google.appengine.ext import webapp
from google.appengine.ext.webapp import util

import base_page
import static_blobs
import breakpad
import event_push
import lkgr
import status
import utils
import xmpp


# Application configuration.
URLS = [
  ('/', status.MainPage),
  ('/allstatus', status.AllStatusPage),
  ('/current', status.CurrentPage),
  ('/status', status.StatusPage),
  ('/status_viewer', status.StatusViewerPage),
  ('/([^/]+\.(?:gif|png|jpg|ico))', static_blobs.ServeHandler),
  ('/static_blobs/list', static_blobs.ListPage),
  ('/static_blobs/(.*)', static_blobs.ServeHandler),
  ('/restricted/static_blobs/upload/(.*)', static_blobs.FormPage),
  ('/restricted/static_blobs/upload_internal/(.*)', static_blobs.UploadHandler),
  ('/revisions', lkgr.Revisions),
  ('/lkgr', lkgr.LastKnownGoodRevision),
  ('/breakpad', breakpad.BreakPad),
  ('/restricted/breakpad/im', breakpad.SendIM),
  ('/restricted/breakpad/cleanup', breakpad.Cleanup),
  ('/recent-events', event_push.RecentEvents),
  ('/status-receiver', event_push.StatusReceiver),
  ('/restricted/status-processor', event_push.StatusProcessor),
  ('/_ah/xmpp/message/chat/', xmpp.XMPPHandler),
]
APPLICATION = webapp.WSGIApplication(URLS, debug=True)


# Register django filters
webapp.template.register_template_library('filters')


def main():
  """Manages and displays chromium tree and revisions status."""
  util.run_wsgi_app(APPLICATION)


if __name__ == "__main__":
  # Do some one-time initializations.
  base_page.bootstrap()
  utils.bootstrap()
  main()
