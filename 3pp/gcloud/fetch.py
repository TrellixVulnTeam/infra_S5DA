#!/usr/bin/env python
# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import json
import os
import sys
import urllib


BASE_URL = 'https://dl.google.com/dl/cloudsdk/channels/rapid'


def do_latest():
  print json.load(urllib.urlopen(BASE_URL + '/components-2.json'))['version']


def do_checkout(version, platform, checkout_path):
  targ_os, targ_arch = platform.split('-')
  ext = 'zip' if targ_os == 'windows' else 'tar.gz'
  download_url = (
    BASE_URL + '/downloads/google-cloud-sdk-%(version)s-%(os)s-%(arch)s.%(ext)s'
    % {
      'version': version,
      'os': {'mac': 'darwin'}.get(targ_os, targ_os),
      'arch': {
        '386':   'x86',
        'amd64': 'x86_64',
      }[targ_arch],
      'ext': ext
    })
  print >>sys.stderr, "fetching", download_url
  urllib.urlretrieve(download_url, os.path.join(checkout_path,
                                                'archive.' + ext))


def main():
  ap = argparse.ArgumentParser()
  sub = ap.add_subparsers()

  latest = sub.add_parser("latest")
  latest.set_defaults(func=lambda _opts: do_latest())

  checkout = sub.add_parser("checkout")
  checkout.add_argument("checkout_path")
  checkout.set_defaults(
    func=lambda opts: do_checkout(
      os.environ['_3PP_VERSION'], os.environ['_3PP_PLATFORM'],
      opts.checkout_path))

  opts = ap.parse_args()
  return opts.func(opts)


if __name__ == '__main__':
  sys.exit(main())
