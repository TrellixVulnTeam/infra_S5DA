(function() {
  'use strict';

  // Refresh delay for on-call rotations in milliseconds.
  // This does not need to refresh very frequently.
  const refreshDelayMs = 60 * 60 * 1000;

  Polymer({
    is: 'som-drawer',

    properties: {
      defaultTree: String,
      _isTrooperPage: {
        type: Boolean,
        computed: '_computeIsTrooperPage(tree.name)',
        value: false,
      },
      path: {
        type: String,
        notify: true,
      },
      _rotations: {
        type: Object,
        value: {
          'android': 'android',
          'chromium': 'chrome',
          'chromiumos': 'chromeosgardener',
          'chromium.perf': 'perfbot',
        },
      },
      _sheriffRotations: Object,
      _sheriffs: {
        type: Array,
        computed: '_computeSheriffs(tree.name, _sheriffRotations, _rotations)',
        value: null,
      },
      _staticPageList: {
        type: Array,
        computed: '_computeStaticPageList(staticPages)',
        value: function() {
          return [];
        },
      },
      tree: Object,
      _treeList: {
        type: Array,
        value: function() {
          return [];
        },
      },
      trees: {
        type: Object,
        notify: true,
        computed: '_computeTrees(_treeList)',
      },
      _trooperRotations: String,
      _troopers: {
        type: Array,
        computed: '_computeTroopers(_trooperRotations)',
        value: null,
      },
      // Settings.
      collapseByDefault: {
        type: Boolean,
        notify: true,
      },
      linkStyle: {
        type: String,
        notify: true,
      },
      showInfraFailures: {
        type: Boolean,
        notify: true,
      },
    },

    created: function() {
      this.async(this._refreshAsync, refreshDelayMs);
    },

    _refresh: function() {
      this.$.fetchTrooper.generateRequest();
      this.$.fetchSheriffs.generateRequest();
    },

    _refreshAsync: function() {
      this._refresh();
      this.async(this._refreshAsync, refreshDelayMs);
    },

    _computeIsTrooperPage: function(treeName) {
      return treeName === 'trooper';
    },

    // Gets current tree sheriffs from file at:
    // http://chromium-build.appspot.com/p/chromium/all_rotations.js
    _computeSheriffs(treeName, sheriffData, rotations) {
      if (!treeName || !sheriffData || !(treeName in rotations)) {
        return null;
      }
      let rotation = rotations[treeName];
      let i = sheriffData['rotations'].indexOf(rotation);

      let date = new Date();

      // Date() is UTC by default. Convert to PST.
      date.setHours(date.getHours() - 8);

      let datestamp = this._formatDate(date);

      let dateData = sheriffData['calendar'].filter(function(obj) {
        return obj.date == datestamp;
      })[0];

      if (!dateData || !dateData['participants']) {
        return null;
      }

      let sheriffs = dateData['participants'][i];

      let sheriffMap = sheriffs.reduce((map, s) => {
        map[s] = date;
        return map;
      }, {});

      // Keep looping through dates until the current sheriffs no long have
      // shifts.
      let flag = true;
      let dayMs = 24 * 60 * 60 * 1000;
      while (flag) {
        flag = false;

        // If the day is a Friday, then we add more days to account for the
        // weekend.
        if (date.getDay() == 5) {
          date = new Date(date.getTime() + dayMs * 3);
        } else {
          date = new Date(date.getTime() + dayMs);
        }
        datestamp = this._formatDate(date);
        dateData = sheriffData['calendar'].filter(function(obj) {
          return obj.date == datestamp;
        })[0];
        if (dateData) {
          let nextSheriffs = dateData['participants'][i];

          let intersection = sheriffs.filter(function(s) {
            return nextSheriffs.indexOf(s) != -1;
          });

          intersection.forEach((s) => {
            flag = true;
            sheriffMap[s] = date;
          });
        }
      }

      let result = Object.keys(sheriffMap).map((key) => {
        return {'username': key, 'endDate': sheriffMap[key]};
      });
      return result;
    },

    _computeStaticPageList(staticPages) {
      let pageList = [];
      for (let key in staticPages) {
        let page = staticPages[key];
        page.name = key;
        pageList.push(page);
      }
      return pageList;
    },

    _computeTrees(treeList) {
      let trees = {};
      if (!treeList) {
        return trees;
      }
      treeList.forEach(function(tree) {
        trees[tree.name] = tree;
      });
      let defaultTree = this.defaultTree;
      if (this.path === '/') {
        if (defaultTree && defaultTree in trees) {
          this.path = '/' + defaultTree;
        } else {
          this.path = '/help-som';
        }
      }
      return trees;
    },

    _computeTroopers(trooperRotations) {
      if (!trooperRotations) {
        return [];
      }

      let troopers = trooperRotations.split(',');
      troopers[0] = troopers[0] + ' (primary)';
      if (troopers.length == 1) {
        return troopers;
      }
      troopers.slice(1).forEach(function(trooper, i) {
        troopers[i + 1] = trooper + ' (secondary)';
      });
      return troopers;
    },

    _formatDate(date) {
      return date.toISOString().substring(0, 10);
    },

    _formatDateShort(date) {
      return moment(date).format('MMM D');
    },

    _onSelected: function(evt) {
      let pathIdentifier = evt.srcElement.value;
      this.path = '/' + pathIdentifier;

      if (pathIdentifier && pathIdentifier in this.trees) {
        this.defaultTree = pathIdentifier;
      }
    },

    toggleMenu: function(e) {
      let target = e.target;
      let collapseId = target.getAttribute('data-toggle-target');
      let collapse = this.$[collapseId];
      collapse.opened = !collapse.opened;

      let icons = target.getElementsByClassName('toggle-icon');
      for (let i = 0; i < icons.length; i++) {
        icons[i].icon = collapse.opened ? 'remove' : 'add';
      }
    },
  });
})();
