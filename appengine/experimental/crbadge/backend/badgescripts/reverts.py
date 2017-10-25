import json

from local_libs.git_checkout import local_git_repository as git_repo

START_REVISION = None

def main():
  counters = [
      RevertStreakCounter(),
      RevertCounter(),
      ComboCounter(),
  ]
  repo = git_repo.LocalGitRepository(
      repo_url='https://chromium.googlesource.com/chromium/src.git')
  all_revisions = repo.GetChangeLogs(START_REVISION, 'HEAD')
  for revision in all_revisions:
    for counter in counters:
      counter(revision)
  print json.dumps(sum((counter.ToDicts() for counter in counters), []),
                   indent=4)

class BaseCounter(object):
  def ToDicts(self):
    all_badges_data = []
    for k, v in self.result.iteritems():
      badge_data = {'badge_name': self.badge_names[k], 'data': []}
      for email, value in v.iteritems():
        badge_data['data'].append({'email': email, 'value': value})
      all_badges_data.append(badge_data)
    return all_badges_data

class RevertCounter(BaseCounter):
  def __init__(self):
    self.badge_names = {
        2: 'code-revert_x2',
        3: 'code-revert_x3',
        4: 'code-revert_x4_plus',
    }
    self.result = {}

  def __call__(self, revision):
      subject = revision.message.splitlines()[0]
      level = subject.count('Revert')
      if level > 4:
        level = 4
      if level > 1:
        self.result.setdefault(level, {})
        self.result[level].setdefault(revision.author.email, 0)
        self.result[level][revision.author.email] += 1

class RevertStreakCounter(BaseCounter):
  def __init__(self):
    self.result = {}
    self.badge_names = {
        2: 'code-revert_consecutive_x2',
        3: 'code-revert_consecutive_x3',
        4: 'code-revert_consecutive_x4_plus',
    }
    self.previous_revision = None
    self.previous_subject = None
    self.streak = 1

  def __call__(self, revision):
    subject = revision.message.splitlines()[0]
    if (self.previous_revision
        and self.previous_revision.author.email == revision.author.email):
      self.previous_subject = self.previous_revision.message.splitlines()[0]
      if subject.startswith('Revert') and self.previous_subject.startswith(
          'Revert'):
        self.streak += 1
        if self.streak > 1:
          if self.streak > 4:
            self.streak = 4
          self.result.setdefault(self.streak, {})
          self.result[self.streak].setdefault(revision.author.email, 0)
          self.result[self.streak][revision.author.email] += 1
      else:
        self.streak = 1
    else:
      self.streak = 1
    self.previous_revision = revision

class ComboCounter(BaseCounter):
  def __init__(self):
    self.result = {}
    # Based on Killer Instinct combo names:
    # https://www.gamefaqs.com/xboxone/718565-killer-instinct/faqs/70214
    self.badge_names = {
        3: 'code-triple_combo', 4: 'code-quad_combo', 5: 'code-solid_combo',
        6: 'code-hyper_combo', 7: 'code-brutal_combo', 8: 'code-master_combo',
        9: 'code-blaster_combo'
    }

  def __call__(self, revision):
      subject = revision.message.splitlines()[0]
      combo = (subject.count('Revert') + subject.count('Reland')
               + subject.count('revert') + subject.count('reland'))
      if combo > 2:
        if combo > 9:
          combo = 9
        self.result.setdefault(combo, {})
        self.result[combo].setdefault(revision.author.email, 0)
        self.result[combo][revision.author.email] += 1

if __name__ == '__main__':
  main()
