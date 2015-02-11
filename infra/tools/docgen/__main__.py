#!/usr/bin/env python
# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Convenience script to generate documentation with Sphinx."""

import argparse
import logging
import os
import shutil
import subprocess
import sys

import infra.path_hacks as path_hacks


def cmd_run():
  """Generate html files from rst files.
  """

  # Clean
  subprocess.check_call(os.path.join('bootstrap', 'remove_orphaned_pycs.py'),
                        cwd=path_hacks.full_infra_path)

  # Add missing rst files
  path = os.path.join(path_hacks.full_infra_path, 'ENV', 'bin', 'sphinx-apidoc')
  subprocess.check_call([path,
                         '-o', os.path.join('doc', 'source', 'reference'),
                         'infra/'], cwd=path_hacks.full_infra_path)

  # Build html documentation for rst files
  path = os.path.join(path_hacks.full_infra_path, 'ENV', 'bin', 'sphinx-build')
  subprocess.check_call([path, '-b', 'html',
                         os.path.join('doc', 'source'),
                         os.path.join('doc', 'html')],
                        cwd=path_hacks.full_infra_path)


def cmd_clean():
  """Remove all files generated by the 'generate' command."""
  paths = (os.path.join(path_hacks.full_infra_path,
                        'doc', 'source', 'reference'),
           os.path.join(path_hacks.full_infra_path, 'doc', 'html'))
  for path in paths:
    try:
      shutil.rmtree(path)
    except OSError:  # pragma: no cover
      pass
    else:
      logging.info('Removing %s ...' % path)


def main(argv):
  if len(argv) == 0:
    return(cmd_run())

  parser = argparse.ArgumentParser()
  subparsers = parser.add_subparsers(title='commands')
  subparsers.add_parser('run').set_defaults(func=cmd_run)
  subparsers.add_parser('clean').set_defaults(func=cmd_clean)

  args = parser.parse_args(argv)
  return args.func()


if __name__ == '__main__':
  sys.exit(main(sys.argv[1:]))
