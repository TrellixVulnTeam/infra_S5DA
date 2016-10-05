"""Provides API wrapper for the codesite issue tracker"""

from endpoints import endpoints
from issue_tracker.issue import Issue
from issue_tracker.comment import Comment


class IssueTrackerAPI(object):  # pragma: no cover
  CAN_ALL = 'all'

  """A wrapper around the issue tracker api."""
  def __init__(self, project_name):
    self.project_name = project_name
    self.client = endpoints.build_client(
        'monorail', 'v1', 'https://monorail-prod.appspot.com/_ah/api/discovery'
        '/v1/apis/{api}/{apiVersion}/rest')

  def create(self, issue, send_email=True):
    body = {}
    assert issue.summary
    body['summary'] = issue.summary
    if issue.description:
      body['description'] = issue.description
    if issue.status:
      body['status'] = issue.status
    if issue.owner:
      body['owner'] = {'name': issue.owner}
    if issue.labels:
      body['labels'] = issue.labels
    if issue.components:
      body['components'] = issue.components
    if issue.cc:
      body['cc'] = [{'name': user} for user in issue.cc]
    request = self.client.issues().insert(
        projectId=self.project_name, sendEmail=send_email, body=body)
    tmp = endpoints._retry__request(request)
    issue.id = int(tmp['id'])
    issue.dirty = False
    return issue

  def update(self, issue, comment=None, send_email=True):
    if not issue.dirty and not comment:
      return issue

    updates = {}
    if 'summary' in issue.changed:
      updates['summary'] = issue.summary
    if 'status' in issue.changed:
      updates['status'] = issue.status
    if 'owner' in issue.changed:
      updates['owner'] = issue.owner
    if 'blocked_on' in issue.changed:
      updates['blockedOn'] = issue.blocked_on
    if issue.labels.isChanged():
      updates['labels'] = list(issue.labels.added)
      updates['labels'].extend('-%s' % label for label in issue.labels.removed)
    if issue.components.isChanged():
      updates['components'] = list(issue.components.added)
      updates['components'].extend(
          '-%s' % comp for comp in issue.components.removed)
    if issue.cc.isChanged():
      updates['cc'] = list(issue.cc.added)
      updates['cc'].extend('-%s' % cc for cc in issue.cc.removed)

    body = {'id': issue.id,
            'updates': updates}

    if comment:
      body['content'] = comment

    request = self.client.issues().comments().insert(
        projectId=self.project_name, issueId=issue.id, sendEmail=send_email,
        body=body)
    endpoints.retry_request(request)

    if issue.owner == '----':
      issue.owner = ''

    issue.dirty = False
    return issue

  def addComment(self, issue_id, comment, send_email=True):
    issue = self.getIssue(issue_id)
    self.update(issue, comment, send_email)

  def getCommentCount(self, issue_id):
    request = self.client.issues().comments().list(
        projectId=self.project_name, issueId=issue_id, startIndex=1,
        maxResults=0)
    feed = endpoints.retry_request(request)
    return feed.get('totalResults', '0')

  def getComments(self, issue_id):
    rtn = []

    request = self.client.issues().comments().list(
        projectId=self.project_name, issueId=issue_id)
    feed = endpoints.retry_request(request)
    rtn.extend([Comment(entry) for entry in feed['items']])
    total_results = feed['totalResults']
    if not total_results:
      return rtn

    while len(rtn) < total_results:
      request = self.client.issues().comments().list(
          projectId=self.project_name, issueId=issue_id, startIndex=len(rtn))
      feed = endpoints.retry_request(request)
      rtn.extend([Comment(entry) for entry in feed['items']])

    return rtn

  def getFirstComment(self, issue_id):
    request = self.client.issues().comments().list(
        projectId=self.project_name, issueId=issue_id, startIndex=0,
        maxResults=1)
    feed = endpoints.retry_request(request)
    if 'items' in feed and len(feed['items']) > 0:
      return Comment(feed['items'][0])
    return None

  def getLastComment(self, issue_id):
    total_results = self.getCommentCount(issue_id)
    request = self.client.issues().comments().list(
        projectId=self.project_name, issueId=issue_id,
        startIndex=total_results-1, maxResults=1)
    feed = endpoints.retry_request(request)
    if 'items' in feed and len(feed['items']) > 0:
      return Comment(feed['items'][0])
    return None

  def getIssue(self, issue_id):
    """Retrieve a set of issues in a project."""
    request = self.client.issues().get(
        projectId=self.project_name, issueId=issue_id)
    entry = endpoints.retry_request(request)
    return Issue(entry)
