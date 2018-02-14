# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import os

from . import util

# Default output path of generated Markdown document.
DEFAULT_PATH = os.path.join(util.PKG_DIR, 'wheels.md')


class Generator(object):
  _INFRA_GITILES_PATH = 'https://chromium.googlesource.com/infra/infra'
  _INFRA_DOCKERBUILD_PATH = 'infra/tools/dockerbuild'

  _HEADER = """\
# dockerbuild-generated wheel packages

This file is automatically generated by `dockerbuild`'s `wheel-dump` subcommand.
It was generated from revision
[%(git_revision)s](%(dockerbuild_url)s).

[TOC]

# Wheel List

This list represents the current set of configured `dockerbuild` wheels in
[wheel.py](%(wheel_py_url)s).

"""

  _WHEEL_NAME_TEMPLATE = """\
## **%(name)s**

"""

  _WHEEL_TEMPLATE = """\
### %(version)s

```protobuf
wheel: <
  name: "%(package_name)s"
  version: "%(package_tag)s"
>
```

%(supported)s

"""

  _FOOTER = """\

# Contact

If a wheel is needed, but is not in this list, please
contact Chrome Operations:

* `luci-eng@google.com`
* [File a Bug](https://bugs.chromium.org/p/chromium/issues/entry?\
template=Build%20Infrastructure)
"""

  def __init__(self, git_revision):
    self._git_revision = git_revision
    self._packages = {}

  @classmethod
  def create(cls, system):
    # Try and get the current Git revision.
    rc, stdout = git_revision = system.run(
        ['git', '-C', util.PKG_DIR, 'rev-parse', 'origin/master'])
    git_revision = stdout.strip() if rc == 0 else None
    return cls(git_revision)

  def dockerbuild_url(self, *subpath):
    parts = [
        self._INFRA_GITILES_PATH,
        '+',
        self._git_revision if self._git_revision else 'master',
        self._INFRA_DOCKERBUILD_PATH,
    ]
    parts.extend(subpath)
    return '/'.join(parts)

  def add_package(self, whl, plat):
    key = (whl.spec.name, whl.spec.version)
    _, v = self._packages.setdefault(key, (whl, set()))
    if plat:
      v.add(plat)

  def write(self, fd):
    fd.write(self._HEADER % dict(
        git_revision=self._git_revision or 'Unknown',
        dockerbuild_url=self.dockerbuild_url(),
        wheel_py_url=self.dockerbuild_url('wheel.py'),
    ))

    categorized_packages = {}
    for (name, version), (whl, plats) in self._packages.iteritems():
      categorized_packages.setdefault(name, {})[version] = (whl, plats)

    for name, versions in sorted(categorized_packages.items()):
      fd.write(self._WHEEL_NAME_TEMPLATE % dict(
          name=name,
      ))

      for version, (whl, plats) in sorted(versions.items()):
        package = whl.cipd_package(self._git_revision, templated=True)

        # Build an italic list of supported platforms.
        plat_names = [plat.name for plat in plats] or ('universal',)
        supported = '\n'.join([''] + [
          '* *%s*' % (plat_name,)
          for plat_name in plat_names])

        fd.write(self._WHEEL_TEMPLATE % dict(
            version=version,
            supported=supported,
            package_name=package.name,
            package_tag=package.tags[0],
        ))

    fd.write(self._FOOTER)
