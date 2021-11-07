# Copyright 2021 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
from . import cipd_manager
from . import git_manager
from . import gcs_manager
from . import helper


class Source:
  """ Source handles all the pinning, downloading, and uploading artifacts to
      Git, GCS and CIPD repositories. See (git|gcs|cipd)_manager.py for
      implementation details on how each is handled.
  """

  def __init__(self, cache, step, path, m_file, raw_io, cipd, gsutil, gitiles):
    """ __init__ generates the src managers (git, gcs and cipd) and stores them
        in class variables.
        Args:
          cache: path to dir that can be used to download artifacts
          step: ref to recipe_engine/step module object
          path: ref to recipe_engine/path module object
          m_file: ref to recipe_engine/file module object
          raw_io: ref to recipe_engine/raw_io module object
          cipd: ref to recipe_engine/cipd module object
          gsutil: ref to depot_tools/gsutil module object
          gitiles: ref to depot_tools/gitiles module object
    """
    # dir to store CIPD downloaded packages
    cipd_dir = cache.join('CIPDPkgs')
    # dir to store GIT downloaded packages
    git_dir = cache.join('GITPkgs')
    # dir to store GCS downloaded packages
    gcs_dir = cache.join('GCSPkgs')
    helper.ensure_dirs(m_file, [
        cipd_dir, git_dir, gcs_dir,
        gcs_dir.join('chrome-gce-images', 'WIB-WIM')
    ])
    self._cipd = cipd_manager.CIPDManager(step, cipd, path, cipd_dir)
    self._gcs = gcs_manager.GCSManager(step, gsutil, path, m_file, raw_io,
                                       gcs_dir)
    self._git = git_manager.GITManager(step, gitiles, m_file, path, git_dir)

  def record_download(self, src):
    """ record_download records src for pinning and downloading purposes
        Args:
          src: sources.Src proto object that contains ref to an artifact
    """
    self._gcs.record_package(src)
    self._git.record_package(src)
    self._cipd.record_package(src)

  def pin(self):
    """ pin pins all the recorded packages to static refs """
    self._gcs.pin_packages()
    self._cipd.pin_packages()
    self._git.pin_packages()

  def download(self):
    """ download downloads all the pinned packages to cache on disk """
    self._gcs.download_packages()
    self._git.download_packages()
    self._cipd.download_packages()

  def record_upload(self, gcs_src, orig=''):
    """ record_upload records an upload that is pending to be performed
        Args:
          gcs_src: sources.GCSSrc proto object that refs a file to be uploaded
          orig: url representing the static ref for this upload
    """
    self._gcs.record_upload(gcs_src, orig)

  def get_local_src(self, src):
    """ get_local_src returns path on the disk that points to the given src ref
        Args:
          src: sources.Src proto object that is ref to a downloaded artifact
    """
    if src and src.WhichOneof('src') == 'gcs_src':
      return self._gcs.get_local_src(src)
    if src and src.WhichOneof('src') == 'git_src':
      return self._git.get_local_src(src)
    if src and src.WhichOneof('src') == 'cipd_src':
      return self._cipd.get_local_src(src)
    if src and src.WhichOneof('src') == 'local_src':
      return src.local_src  # pragma: no cover

  def get_url(self, src):
    """ get_url returns string containing an url referencing the given src
        Args:
          src: sources.Src proto object that contains ref to an artifact
    """
    return self._gcs.get_gs_url(src)

  def upload(self):
    """ upload uploads all the available files to be uploaded if available """
    self._gcs.upload_packages()