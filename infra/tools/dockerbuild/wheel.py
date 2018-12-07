# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Manages the generation and uploading of Python wheel CIPD packages.

== A Note on Universiality ==

Our wheel generation uses CIPD package naming to represent the platform,
architecture, and ABI of generated wheels. Because of that, we offload the
actual Python wheel naming (which "pip" checks when installing wheels) to the
CIPD naming layer and represent every generated wheel filename as universal.

Naming wheels universally will cause "pip" to agreeably install any wheel onto
any platform, effectively establishing the CIPD naming scheme as the absolute
authority removing any protection against mismatched tags that "pip" supplies.

We are deferring to CIPD's platform resolution for backwards-compatible
architectures, notably:

  - CIPD's "armv6l" platform loads on "armv7l", "armv8l", "armv9l", ...
  - CIPD's "arm64" platform loads on "aarch64".

An alternative is to enumerate compatible platforms explicitly, such as
"linux_armv6l.linux_armv7l.linux_armv8l.linux_armv9l". This can work in the
short-term, but imagine adding the hypothetically-backwards-compatible
"linux_armv10l" platform:

  - CIPD will continue to resolve "${platform}" to "linux-armv6l".
  - The "linux-armv6l" CIPD package will contain a wheel with the above platform
    enumeration, which will NOT resolve on the new "linux_armv10l" platform.
  - We probably recognized this because users reported failures on their
    "linux_armv10l" systems, meaning that the "linux-armv6l" CIPD package is
    cached.

We can respond in one of three ways:
  - Forcefully differentiate "armv10l" from its predecessors. This is not
    technically consistent with other CIPD platform approaches, and would be
    done for convenience, rather than effective, reasons.
  - Append ".linux_armv10l" to the wheel's platform enumeration, generate new
    CIPD packages with different tags, and update every "vpython" spec file to
    reference the new tags. This is tedious.
  - Append ".linux_armv10l" to the wheel's platform enumeration, delete the
    current CIPD package, and manually purge any "linux-armv10l" systems' CIPD
    caches. This is tedious and unsupported.

None of these options is particularly great, nor are any of them guaranteed to
scale based on the new platform circumstances. Consequently, we remove the
architecture decision from "pip" altogether by using a universal name for every
wheel. Now, "pip" will install any packaged wheel on any platform, so it is
our decision to NOT install the wrong wheels on incompatible platforms. We use
CIPD and "vpython"'s naming resolution to ensure that this doesn't happen.

Internally to "dockerbuild", we still name wheels appropriately; we only
make them universal for CIPD packaging. This allows correct wheels to still be
generated by "dockerbuild" for non-CIPD purposes.
"""

from . import source
from . import platform

from .wheel_cryptography import CryptographyBuilder as Cryptography

from .wheel_wheel import PrebuiltBuilder as Prebuilt
from .wheel_wheel import SourceOrPrebuiltBuilder as SourceOrPrebuilt
from .wheel_wheel import UniversalBuilder as Universal
from .wheel_wheel import UniversalSourceBuilder as UniversalSource
from .wheel_wheel import MultiWheelBuilder as MultiWheel

from .wheel_opencv import OpenCVBuilder as OpenCV

from .wheel_infra import InfraBuilder as Infra

from .wheel_mysql import MySQLPythonBuilder as MySQLPython

# This is the NumPy-ecosystem list of platforms that their mac-x64 wheel
# supports.
_NUMPY_MAC_x64 = [
  'macosx_10_6_intel',
  'macosx_10_9_intel',
  'macosx_10_9_x86_64',
  'macosx_10_10_intel',
  'macosx_10_10_x86_64',
]


SPECS = {s.spec.tag: s for s in (
  SourceOrPrebuilt('bcrypt', '3.1.4',
      packaged=[
        'mac-x64',
        'manylinux-x86',
        'manylinux-x64',
        'windows-x86',
        'windows-x64',
      ],
  ),
  SourceOrPrebuilt('coverage', '4.3.4'),
  SourceOrPrebuilt('coverage', '4.5.1',
      skip_plat=['mac-x64'],
  ),
  SourceOrPrebuilt('cffi', '1.10.0',
    arch_map={
      'mac-x64': ['macosx_10_6_intel'],
    },
  ),
  SourceOrPrebuilt('numpy', '1.11.3',
      abi_map={
        'windows-x86': 'none',
        'windows-x64': 'none',
      },
      arch_map={'mac-x64': _NUMPY_MAC_x64},
  ),
  SourceOrPrebuilt('numpy', '1.12.1',
      abi_map={
        'windows-x86': 'none',
        'windows-x64': 'none',
      },
      arch_map={'mac-x64': _NUMPY_MAC_x64},
      skip_plat=('linux-arm64',),
  ),

  SourceOrPrebuilt('psutil', '1.2.1',
      packaged=[],
      only_plat=platform.ALL_LINUX + ['mac-x64'],
  ),

  SourceOrPrebuilt('psutil', '5.2.2',
      abi_map={
        'windows-x86': 'none',
        'windows-x64': 'none',
      },
      arch_map={'mac-x64': _NUMPY_MAC_x64},
      packaged=['windows-x86', 'windows-x64'],
  ),

  Prebuilt('pypiwin32', '219',
      ['windows-x86', 'windows-x64'],
  ),

  Prebuilt('scipy', '0.19.0',
      ['mac-x64', 'manylinux-x86', 'manylinux-x64'],
      arch_map={'mac-x64': _NUMPY_MAC_x64},
  ),

  SourceOrPrebuilt('pyasn', '1.6.0b1', packaged=(),
      only_plat=['manylinux-x64']),
  Prebuilt('pynacl', '1.2.1', ['manylinux-x64', 'mac-x64']),

  Prebuilt('pandas', '0.23.4',
      ['manylinux-x86', 'manylinux-x64',
       'mac-x64',
       'windows-x86', 'windows-x64'],
      arch_map={'mac-x64': _NUMPY_MAC_x64},
  ),

  OpenCV('opencv_python', '2.4.13.2', '1.11.3',
      only_plat=['manylinux-x86', 'manylinux-x64']),

  OpenCV('opencv_python', '3.2.0.7', '1.12.1',
      packaged=[
        'mac-x64',
        'manylinux-x86',
        'manylinux-x64',
        'windows-x86',
        'windows-x64',
      ],
      arch_map={'mac-x64': _NUMPY_MAC_x64},
      only_plat=[
        'mac-64', 'manylinux-x86', 'manylinux-x64', 'windows-x86',
        'windows-x64'],
  ),

  Cryptography('cryptography',
      source.pypi_sdist('cryptography', '2.0.3'),
      source.remote_archive(
          name='openssl',
          version='1.1.0f',
          url='https://www.openssl.org/source/openssl-1.1.0f.tar.gz',
      ),
      arch_map={
        'mac-x64': ['macosx_10_6_intel'],
      },
      packaged=[
        'manylinux-x86',
        'manylinux-x64',
        'mac-x64',
        'windows-x86',
        'windows-x64',
      ],
  ),

  SourceOrPrebuilt('crcmod', '1.7', packaged=(),
      skip_plat=[
        'windows-x86',
        'windows-x64',
      ],
  ),
  SourceOrPrebuilt('grpcio', '1.4.0'),
  SourceOrPrebuilt('scan-build', '2.0.8', packaged=(),
      skip_plat=[
        'mac-x64',
        'windows-x86',
        'windows-x64',
      ],
  ),
  SourceOrPrebuilt('scandir', '1.7',
      packaged=[
        'windows-x86',
        'windows-x64',
      ],
      skip_plat=['mac-x64'],
  ),

  SourceOrPrebuilt('simplejson', '3.13.2',
      packaged=[
        'windows-x86',
        'windows-x64',
      ],
  ),

  # Prefer to use 'cryptography' instead of PyCrypto, if possible. We have to
  # use PyCrypto for GAE dev server (it's the only crypto package available on
  # GAE). Since we support it only on Linux and OSX, build only for these
  # platforms.
  SourceOrPrebuilt('pycrypto', '2.6.1',
      packaged=(),
      only_plat=['manylinux-x64', 'mac-x64'],
  ),

  # List cultivated from "pyobjc-2.5.1"'s "setup.py" as a superset of available
  # packages.
  #
  # - This must be built on Mac 10.9 or lower due to a version string
  #   parsing error in "setup.py" that rates "10.10" as a lower version than
  #   "10.9".
  # - The package requires that "setuptools==1" package be installed. Since
  #   "run.py" doesn't support this version, the user must create their own
  #   VirtualEnv, manually install "setuptools==1", and then use the
  #   "--native-python" Dockerbuild flag to build using that Python.
  MultiWheel('pyobjc', '2.5.1', (
    [SourceOrPrebuilt(name, '2.5.1', packaged=[])
     for name in ['pyobjc-core'] + ['pyobjc-framework-%s' % (v,) for v in [
        'Accounts', 'AddressBook', 'AppleScriptKit', 'AppleScriptObjC',
        'Automator', 'CFNetwork', 'CalendarStore', 'Cocoa', 'Collaboration',
        'CoreData', 'CoreLocation', 'CoreText', 'DictionaryServices',
        'EventKit', 'ExceptionHandling', 'FSEvents', 'InputMethodKit',
        'InstallerPlugins', 'InstantMessage', 'LatentSemanticMapping',
        'LaunchServices', 'Message', 'PreferencePanes', 'PubSub', 'QTKit',
        'Quartz', 'ScreenSaver', 'ScriptingBridge', 'SearchKit',
        'ServerNotification', 'ServiceManagement', 'Social', 'SyncServices',
        'SystemConfiguration', 'WebKit',
      ]]
    ]),
    only_plat=['mac-x64'],
    # Because this requires a specialized environment, we will not include it in
    # the default wheel list.
    default=False,
  ),

  # List cultivated from "pyobjc-4.1"'s "setup.py" as a superset of available
  # packages.
  #
  # This package is designed to be built on 10.12, and had to omit the following
  # framework packages, which require 10.13 to build:
  # - CoreML
  # - CoreSpotlight
  # - ExternalAccessory
  # - Vision
  MultiWheel('pyobjc', '4.1', (
    [SourceOrPrebuilt(name, '4.1', packaged=[])
     for name in ['pyobjc-core'] + ['pyobjc-framework-%s' % (v,) for v in [
        'AVFoundation', 'AVKit', 'Accounts', 'AddressBook', 'AppleScriptKit',
        'AppleScriptObjC', 'ApplicationServices', 'Automator', 'CFNetwork',
        'CalendarStore', 'CloudKit', 'Cocoa', 'Collaboration', 'ColorSync',
        'Contacts', 'ContactsUI', 'CoreBluetooth', 'CoreData', 'CoreLocation',
        'CoreServices', 'CoreText', 'CoreWLAN', 'CryptoTokenKit',
        'DictionaryServices', 'DiskArbitration', 'EventKit',
        'ExceptionHandling', 'FSEvents', 'FinderSync', 'GameCenter',
        'GameController', 'GameKit', 'GameplayKit', 'IMServicePlugIn',
        'IOSurface', 'ImageCaptureCore', 'InputMethodKit', 'InstallerPlugins',
        'InstantMessage', 'Intents', 'InterfaceBuilderKit',
        'LatentSemanticMapping', 'LaunchServices', 'LocalAuthentication',
        'MapKit', 'MediaAccessibility', 'MediaLibrary', 'MediaPlayer',
        'Message', 'ModelIO', 'MultipeerConnectivity', 'NetFS',
        'NetworkExtension', 'NotificationCenter', 'OpenDirectory', 'Photos',
        'PhotosUI', 'PreferencePanes', 'PubSub', 'QTKit', 'Quartz',
        'SafariServices', 'SceneKit', 'ScreenSaver', 'ScriptingBridge',
        'SearchKit', 'Security', 'SecurityFoundation', 'SecurityInterface',
        'ServerNotification', 'ServiceManagement', 'Social', 'SpriteKit',
        'StoreKit', 'SyncServices', 'SystemConfiguration', 'WebKit',
        'XgridFoundation', 'iTunesLibrary', 'libdispatch',
      ]]
    ]),
    only_plat=['mac-x64'],
  ),


  SourceOrPrebuilt('lazy-object-proxy', '1.3.1',
      packaged=[
        'manylinux-x86',
        'manylinux-x64',
        'windows-x86',
        'windows-x64',
      ],
  ),
  Prebuilt('lxml', '4.2.5',
      ['mac-x64',
       'manylinux-x86', 'manylinux-x64',
       'windows-x86', 'windows-x64'],
  ),
  SourceOrPrebuilt('MarkupSafe', '1.0', packaged=(),
      only_plat=platform.ALL_LINUX,
  ),
  MySQLPython('1.2.5',
      only_plat=[
        'manylinux-x86', 'manylinux-x64',
        'linux-arm64', 'linux-armv6',
        'linux-mips64',
      ],
  ),
  SourceOrPrebuilt('PyYAML', '3.12',
      packaged=[
        'windows-x86',
        'windows-x64',
      ],
  ),
  SourceOrPrebuilt('SQLAlchemy', '1.2.5', packaged=(),
      only_plat=['manylinux-x86', 'manylinux-x64'],
  ),
  SourceOrPrebuilt('subprocess32', '3.5.0rc1',
      packaged=['mac-x64'],
      skip_plat=[
        'windows-x86',
        'windows-x64',
      ],
  ),
  SourceOrPrebuilt('wrapt', '1.10.11', packaged=(),
      only_plat=['manylinux-x86', 'manylinux-x64'],
  ),

  Universal('appdirs', '1.4.3'),
  UniversalSource('apache-beam', '2.0.0'),
  UniversalSource('Appium_Python_Client', '0.24',
                   pypi_name='Appium-Python-Client'),
  Universal('asn1crypto', '0.22.0'),
  Universal('astroid', '1.6.3'),
  Universal('astunparse', '1.5.0'),
  Universal('attrs', '17.4.0'),
  UniversalSource('backport_ipaddress', '0.1', pyversions=['py2']),
  Universal('backports.functools_lru_cache', '1.5'),
  Universal('boto', '2.48.0'),
  Universal('cachetools', '2.0.1'),
  Universal('contextlib2', '0.5.5'),
  Universal('cheroot', '6.2.4'),
  Universal('CherryPy', '14.2.0'),
  UniversalSource('configparser', '3.5.0'),
  UniversalSource('Django', '1.5.1'),
  Universal('Django', '1.9'),
  Universal('enum34', '1.1.6', pyversions=['py2', 'py3']),
  Universal('fabric', '1.14.0'),
  Universal('funcsigs', '1.0.2'),
  UniversalSource('future', '0.16.0'),
  Universal('futures', '3.1.1'),
  Universal('gitdb2', '2.0.3'),
  Universal('GitPython', '2.1.9'),

  Universal('google-api-core', '0.1.1'),
  Universal('google-auth', '1.2.1'),
  Universal('google-cloud-bigquery', '0.28.0'),
  Universal('google-cloud-bigtable', '0.28.1'),
  Universal('google-cloud-core', '0.28.0'),
  Universal('google-cloud-datastore', '1.6.0'),
  Universal('google-cloud-dns', '0.28.0'),
  Universal('google-cloud-firestore', '0.28.0'),
  Universal('google-cloud-language', '1.0.0'),
  Universal('google-cloud-logging', '1.4.0'),
  Universal('google-cloud-monitoring', '0.28.0'),
  Universal('google-cloud-pubsub', '0.29.0'),
  Universal('google-cloud-resource-manager', '0.28.0'),
  Universal('google-cloud-runtimeconfig', '0.28.0'),
  Universal('google-cloud-spanner', '0.29.0'),
  Universal('google-cloud-speech', '0.30.0'),
  Universal('google-cloud-storage', '1.6.0'),
  Universal('google-cloud-translate', '1.3.0'),
  Universal('google-cloud-videointelligence', '0.28.0'),
  Universal('google-resumable-media', '0.3.1'),
  Universal('google_api_python_client', '1.6.2'),
  UniversalSource('google-cloud-trace', '0.16.0'),
  UniversalSource('google_compute_engine', '2.6.2',
                  pypi_name='google-compute-engine'),
  UniversalSource('googleapis-common-protos', '1.5.3'),

  UniversalSource('httplib2', '0.10.3'),
  Universal('idna', '2.5'),
  UniversalSource('inotify_simple', '1.1.7'),
  Universal('ipaddress', '1.0.18', pyversions=['py2']),
  Universal('isort', '4.3.4'),
  Universal('Jinja2', '2.10'),
  Universal('json5', '0.6.0'),
  Universal('mccabe', '0.6.1'),
  Universal('mock', '2.0.0'),
  Universal('more-itertools', '4.1.0'),
  UniversalSource('mox', '0.5.3'),
  Universal('nose', '1.3.7'),
  UniversalSource('oauth2client', '3.0.0'),
  Universal('oauth2client', '4.0.0'),
  Universal('oauth2client', '4.1.2'),
  Universal('packaging', '16.8'),
  Universal('paramiko', '2.4.1'),
  UniversalSource('PeakUtils', '1.0.3'),
  Universal('pluggy', '0.6.0'),
  Universal('pbr', '3.0.0'),
  Universal('ply', '3.11'),
  Universal('portend', '2.2'),
  Universal('protobuf', '3.2.0'),
  Universal('py', '1.5.3'),
  Universal('pyasn1', '0.2.3'),
  Universal('pyasn1_modules', '0.0.8'),
  UniversalSource('pycparser', '2.17'),
  UniversalSource('pyftpdlib', '0.7.0'),
  UniversalSource('pyftpdlib', '1.0.0'),
  UniversalSource('pyftpdlib', '1.5.3'),
  Universal('pyopenssl', '17.2.0'),
  Universal('pylint', '1.8.4'),
  Universal('pyparsing', '2.2.0'),
  Universal('pyserial', '3.4'),
  Universal('pytest', '3.5.0'),
  Universal('pytest-cov', '2.5.1'),
  Universal('python-dateutil', '2.7.3'),
  Universal('pytz', '2018.4'),
  Universal('requests', '2.13.0'),
  UniversalSource('requests-unixsocket', '0.1.5'),
  Universal('rsa', '3.4.2'),
  Universal('selenium', '3.4.1'),
  Universal('selenium', '3.14.0'),
  Universal('setuptools', '34.3.2'),
  Universal('singledispatch', '3.4.0.3'),
  Universal('six', '1.10.0'),
  Universal('six', '1.11.0'),
  Universal('smmap2', '2.0.3'),
  Universal('tempora', '1.11'),
  UniversalSource('tlslite', '0.4.9'),
  Universal('typing', '3.6.4'),
  Universal('uritemplate', '3.0.0'),
  Universal('urllib3', '1.22'),
  Universal('yapf', '0.22.0'),
  Universal('yapf', '0.24.0'),

  Infra('infra_libs'),
)}
SPEC_NAMES = sorted(SPECS.keys())
DEFAULT_SPEC_NAMES = sorted([s.spec.tag for s in SPECS.itervalues()
                             if s.spec.default])
