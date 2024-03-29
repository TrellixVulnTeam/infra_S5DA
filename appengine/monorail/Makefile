# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file or at
# https://developers.google.com/open-source/licenses/bsd

# Makefile to simplify some common AppEngine actions.
# Use 'make help' for a list of commands.

DEVID = monorail-dev
STAGEID= monorail-staging
PRODID= monorail-prod

GAE_PY?= python gae.py
DEV_APPSERVER_FLAGS?= --watcher_ignore_re="(.*/lib|.*/node_modules|.*/third_party|.*/venv)"

WEBPACK_PATH := ./node_modules/webpack-cli/bin/cli.js

TARDIR ?= "/workspace"

FRONTEND_MODULES?= default
BACKEND_MODULES?= besearch latency-insensitive api

BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)

PY_DIRS = api,businesslogic,features,framework,project,proto,search,services,sitewide,testing,tracker

_VERSION ?= $(shell ../../../infra/luci/appengine/components/tools/calculate_version.py)

RUN_CLOUD_TASKS := ../../cipd/cloud-tasks-emulator \
	  --queue projects/monorail-staging/locations/us-central1/queues/componentexport \
		--queue projects/monorail-staging/locations/us-central1/queues/default \
		--queue projects/monorail-staging/locations/us-central1/queues/notifications \
		--queue projects/monorail-staging/locations/us-central1/queues/outboundemail \
		--queue projects/monorail-staging/locations/us-central1/queues/recomputederivedfields \
		--queue projects/monorail-staging/locations/us-central1/queues/spamexport \
		--queue projects/monorail-staging/locations/us-central1/queues/wipeoutsendusers \
		--queue projects/monorail-staging/locations/us-central1/queues/wipeoutdeleteusers \
		--queue projects/monorail-staging/locations/us-central1/queues/deleteusers \
		--queue projects/monorail-staging/locations/us-central1/queues/pubsub-issueupdates \
		-host localhost \
		-port 9090

default: help

check:
ifndef NPM_VERSION
	$(error npm not found. Install from nodejs.org or see README)
endif

help:
	@echo "Available commands:"
	@sed -n '/^[a-zA-Z0-9_.]*:/s/:.*//p' <Makefile

# Run "eval `../../go/env.py`" before running the following prpc_proto commands
prpc_proto_v0:
	touch ../../ENV/lib/python2.7/site-packages/google/__init__.py
	PYTHONPATH=../../ENV/lib/python2.7/site-packages \
	PATH=../../luci/appengine/components/tools:$(PATH) \
	../../cipd/protoc \
	--python_out=. --prpc-python_out=. api/api_proto/*.proto
	cd ../../go/src/infra/monorailv2 && \
	cproto -proto-path ../../../../appengine/monorail/ ../../../../appengine/monorail/api/api_proto/
prpc_proto_v3:
	touch ../../ENV/lib/python2.7/site-packages/google/__init__.py
	PYTHONPATH=../../ENV/lib/python2.7/site-packages \
	PATH=../../luci/appengine/components/tools:$(PATH) \
	../../cipd/bin/protoc \
	--python_out=. --prpc-python_out=. api/v3/api_proto/*.proto
	cd ../../go/src/infra/monorailv2 && \
	cproto -proto-path ../../../../appengine/monorail/ ../../../../appengine/monorail/api/v3/api_proto/

business_proto:
	touch ../../ENV/lib/python2.7/site-packages/google/__init__.py
	PYTHONPATH=../../ENV/lib/python2.7/site-packages \
	PATH=../../luci/appengine/components/tools:$(PATH) \
	../../cipd/protoc \
	--python_out=. --prpc-python_out=. proto/*.proto

test:
	../../test.py test appengine/monorail

test_no_coverage:
	../../test.py test appengine/monorail --no-coverage

coverage:
	@echo "Running tests + HTML coverage report in ~/monorail-coverage:"
	../../test.py test appengine/monorail --html-report ~/monorail-coverage --coveragerc appengine/monorail/.coveragerc

# Shows coverage on the tests themselves, helps illuminate when we have test
# methods that aren't used.
test_coverage:
	@echo "Running tests + HTML coverage report (for tests) in ~/monorail-test-coverage:"
	../../test.py test appengine/monorail --html-report ~/monorail-test-coverage --coveragerc appengine/monorail/.testcoveragerc

# Commands for running locally using dev_appserver.
# devserver requires an application ID (-A) to be specified.
# We are using `-A monorail-staging` because ml spam code is set up
# to impersonate monorail-staging in the local environment.
serve:
	@echo "---[Starting SDK AppEngine Server]---"
	$(GAE_PY) devserver -A monorail-staging -- $(DEV_APPSERVER_FLAGS)\
	& $(WEBPACK_PATH) --watch\
	& $(RUN_CLOUD_TASKS)

serve_email:
	@echo "---[Starting SDK AppEngine Server]---"
	$(GAE_PY) devserver -A monorail-staging -- $(DEV_APPSERVER_FLAGS) --enable_sendmail=True\
	& $(WEBPACK_PATH) --watch\
	& $(RUN_CLOUD_TASKS)

# The _remote commands expose the app on 0.0.0.0, so that it is externally
# accessible by hostname:port, rather than just localhost:port.
serve_remote:
	@echo "---[Starting SDK AppEngine Server]---"
	$(GAE_PY) devserver -A monorail-staging -o -- $(DEV_APPSERVER_FLAGS)\
	& $(WEBPACK_PATH) --watch\
	& $(RUN_CLOUD_TASKS)

serve_remote_email:
	@echo "---[Starting SDK AppEngine Server]---"
	$(GAE_PY) devserver -A monorail-staging -o -- $(DEV_APPSERVER_FLAGS) --enable_sendmail=True\
	& $(WEBPACK_PATH) --watch\
	& $(RUN_CLOUD_TASKS)

run: serve

deps: node_deps
	rm -f static/dist/*

build_js:
	$(WEBPACK_PATH) --mode=production

clean_deps:
	rm -rf node_modules

node_deps:
	npm ci --no-save

dev_deps:
	python -m pip install --no-deps -r requirements.dev.txt

karma:
	npx karma start --debug --coverage

karma_debug:
	npx karma start --debug

pylint:
	pylint -f parseable *py {$(PY_DIRS)}{/,/test/}*py

py3lint:
	pylint --py3k *py {$(PY_DIRS)}{/,/test/}*py

deploy_dev: build_frontend
	$(eval BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD))
	@echo "---[Dev $(DEVID)]---"
	$(GAE_PY) upload --tag $(BRANCH_NAME) -A $(DEVID) $(FRONTEND_MODULES) $(BACKEND_MODULES)

# Build target used by LUCI CD and manual process to build frontend.
build_frontend: clean_deps deps build_js

package_release:
	rsync -aLK . $(TARDIR)/package

lsbuilds:
	gcloud builds list --filter="tags='monorail'"

# AppEngine apps can be tested locally and in non-default versions upload to
# the main app-id, but it is still sometimes useful to have a completely
# separate app-id.  E.g., for testing inbound email, load testing, or using
# throwaway databases.
deploy_staging: build_frontend
	@echo "---[Staging $(STAGEID)]---"
	$(GAE_PY) upload -A $(STAGEID) $(FRONTEND_MODULES) $(BACKEND_MODULES)

# This is our production server that users actually use.
deploy_prod: build_frontend
	@echo "---[Deploying prod instance $(PRODID)]---"
	$(GAE_PY) upload -A $(PRODID) $(FRONTEND_MODULES) $(BACKEND_MODULES)

# Note that we do not provide a command-line way to make the newly-uploaded
# version the default version. This is for two reasons: a) You should be using
# your browser to confirm that the new version works anyway, so just use the
# console interface to make it the default; and b) If you really want to use
# the command line you can use gae.py directly.
