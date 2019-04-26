# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Updates the Git Cache zip files."""

from recipe_engine import recipe_api

DEPS = [
  'depot_tools/depot_tools',
  'depot_tools/gclient',
  'recipe_engine/buildbucket',
  'recipe_engine/context',
  'recipe_engine/json',
  'recipe_engine/path',
  'recipe_engine/properties',
  'recipe_engine/python',
  'recipe_engine/runtime',
  'recipe_engine/url',
]


PROPERTIES = {
  'bucket': recipe_api.Property(
      default=None, help='override GS bucket to upload cached git repos to'),
  'repo_urls': recipe_api.Property(
      default=None,
      help='List of repo urls to limit work to just these repos. Each must:\n'
           ' * not have /a/ as path prefix\n'
           ' * no trailing slash\n'
           ' * no .git suffix\n'
           'For example, "https://chromium.googlesource.com/infra/infra".')
}


BUILDER_MAPPING = {
  'git-cache-chromium': 'https://chromium.googlesource.com/',
}

TEST_REPOS = """All-Projects
All-Users
apps
chromium/src
foo/bar"""


def list_host_repos(api, host_url):
  with api.depot_tools.on_path():
    output = api.url.get_text('%s?format=TEXT' % host_url,
                              default_test_data=TEST_REPOS).output
    return ['%s%s' % (host_url, repo)
            for repo in output.splitlines()
            if repo.lower() not in ['all-projects', 'all-users']]


def RunSteps(api, bucket, repo_urls):
  if not repo_urls:
    repo_urls = list_host_repos(
        api, BUILDER_MAPPING[api.buildbucket.builder_name])

  api.gclient.set_config('infra')
  api.gclient.c.solutions[0].revision = 'origin/master'
  api.gclient.checkout()
  api.gclient.runhooks()

  # Turn off the low speed limit, since checkout will be long.
  env = {
    'GIT_HTTP_LOW_SPEED_LIMIT': '0',
    'GIT_HTTP_LOW_SPEED_TIME': '0',
  }
  if api.runtime.is_experimental:
    assert bucket, 'bucket property is required in experimental mode'
    env['OVERRIDE_BOOTSTRAP_BUCKET'] = bucket

  with api.context(env=env):
    # Run the updater script.
    with api.depot_tools.on_path():
      for url in repo_urls:
        api.python(
            'Updating %s' % url,
            api.path['start_dir'].join('infra', 'run.py'),
            [
              'infra.services.git_cache_updater',
              '--repo', url,
            '--work-dir', api.path['start_dir'].join('cache_dir')
          ],
          # It's fine for this to fail, just move on to the next one.
          ok_ret='any')


def GenTests(api):
  yield (
      api.test('git-cache-chromium') +
      api.buildbucket.try_build(builder='git-cache-chromium')
  )
  yield (
      api.test('git-cache-chromium-led-triggered') +
      api.runtime(is_luci=True, is_experimental=True) +
      api.properties(bucket='experimental-gs-bucket',
                     repo_urls=['https://chromium.googlesource.com/v8/v8']) +
      api.buildbucket.try_build(builder='git-cache-chromium')
  )
