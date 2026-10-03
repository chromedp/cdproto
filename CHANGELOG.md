# Changelog

Each release of `cdproto` is generated from the Chromium and V8 protocol
definitions listed below. The minor version is the Chromium major version. As
the protocol definitions deprecate and remove commands, events, types, and
fields, any release can contain incompatible changes to the generated API.

## v0.157.4 - 2026-10-03

- Chromium: 157.0.8085.1
- V8: 15.7.33
- API changes since v0.157.3: 72 incompatible, 1 compatible

By package, as removed, changed and added names:

  accessibility               0 removed    1 changed    0 added
  animation                   0 removed    1 changed    0 added
  audits                      0 removed    1 changed    0 added
  browser                     0 removed    2 changed    0 added
  cdp                         1 removed    0 changed    0 added
  css                         0 removed    1 changed    0 added
  debugger                    0 removed    2 changed    0 added
  dom                         0 removed    3 changed    0 added
  domdebugger                 0 removed    1 changed    0 added
  emulation                   0 removed   19 changed    0 added
  har                         0 removed    7 changed    0 added
  headlessexperimental        0 removed    1 changed    0 added
  indexeddb                   0 removed    2 changed    0 added
  input                       0 removed    8 changed    0 added
  io                          0 removed    1 changed    0 added
  layertree                   0 removed    1 changed    0 added
  network                     0 removed    5 changed    1 added
  overlay                     0 removed    2 changed    0 added
  page                        0 removed    6 changed    0 added
  runtime                     0 removed    1 changed    0 added
  storage                     0 removed    1 changed    0 added
  target                      0 removed    2 changed    0 added
  webauthn                    0 removed    3 changed    0 added

## v0.157.3 - 2026-10-03

- Chromium: 157.0.8085.1
- V8: 15.7.33
- API changes since v0.157.2: 2237 incompatible, 1079 compatible

By package, as removed, changed and added names:

  accessibility              30 removed   17 changed    8 added
  ads                         4 removed    4 changed    2 added
  animation                  15 removed   20 changed    7 added
  audits                      8 removed    9 changed    3 added
  autofill                    7 removed    8 changed    1 added
  backgroundservice           4 removed    8 changed    2 added
  bluetoothemulation         20 removed   40 changed    6 added
  browser                    45 removed   47 changed   10 added
  cachestorage               14 removed   12 changed    3 added
  cast                        7 removed   12 changed    2 added
  cdp                         4 removed    0 changed    6 added
  crashreportcontext          2 removed    2 changed    1 added
  css                        76 removed   79 changed   35 added
  debugger                   69 removed   71 changed   16 added
  deviceaccess                4 removed    8 changed    1 added
  deviceorientation           2 removed    4 changed    0 added
  digitalcredentials          4 removed    2 changed    0 added
  dom                       136 removed  116 changed   51 added
  domdebugger                13 removed   17 changed    1 added
  domsnapshot                 8 removed   10 changed    1 added
  domstorage                  7 removed   12 changed    5 added
  emulation                 105 removed   98 changed    7 added
  eventbreakpoints            3 removed    6 changed    0 added
  extensions                 13 removed   17 changed    3 added
  fedcm                       9 removed   16 changed    2 added
  fetch                      26 removed   24 changed    4 added
  filesystem                  2 removed    2 changed    1 added
  findinpage                  4 removed    8 changed    0 added
  har                         0 removed    2 changed    1 added
  headlessexperimental        6 removed    3 changed    1 added
  heapprofiler               28 removed   33 changed    9 added
  indexeddb                  36 removed   18 changed    4 added
  input                      65 removed   30 changed    1 added
  inspector                   2 removed    4 changed    4 added
  io                          7 removed    6 changed    2 added
  layertree                  21 removed   20 changed    8 added
  log                         5 removed   10 changed    1 added
  media                       2 removed    4 changed    5 added
  memory                     18 removed   23 changed    5 added
  network                    69 removed   96 changed   54 added
  overlay                    51 removed   59 changed    9 added
  page                      133 removed  132 changed   47 added
  performance                 5 removed    6 changed    2 added
  performancetimeline         1 removed    2 changed    1 added
  preload                     2 removed    4 changed    6 added
  profiler                   16 removed   21 changed    7 added
  pwa                        14 removed   15 changed    3 added
  runtime                    77 removed   71 changed   19 added
  security                    3 removed    6 changed    1 added
  serviceworker              12 removed   24 changed    3 added
  smartcardemulation         14 removed   36 changed   14 added
  storage                    39 removed   54 changed   15 added
  systeminfo                  6 removed    6 changed    3 added
  target                     54 removed   46 changed   16 added
  tethering                   2 removed    4 changed    1 added
  tracing                    20 removed   17 changed    6 added
  webaudio                    4 removed    6 changed   14 added
  webauthn                   25 removed   46 changed    7 added
  webmcp                      5 removed    8 changed    5 added

## v0.157.2 - 2026-10-03

- Chromium: 157.0.8085.1 (was 157.0.8084.3)
- V8: 15.7.33 (was 15.7.23)
- API changes since v0.157.1: 0 incompatible, 0 compatible

## v0.157.1 - 2026-10-03

- Chromium: 157.0.8084.3
- V8: 15.7.23
- API changes since v0.157.0: 267 incompatible, 493 compatible

By package, as removed, changed and added names:

  accessibility               4 removed    0 changed    0 added
  animation                   0 removed    1 changed    4 added
  audits                     29 removed    2 changed    4 added
  autofill                    1 removed    0 changed    0 added
  backgroundservice           1 removed    0 changed    0 added
  bluetoothemulation          5 removed    0 changed    0 added
  browser                     4 removed    3 changed    9 added
  cachestorage                1 removed    0 changed    0 added
  cdp                        11 removed    0 changed    0 added
  css                         2 removed    3 changed   17 added
  debugger                    1 removed   13 changed   49 added
  digitalcredentials          1 removed    0 changed    0 added
  dom                         3 removed    6 changed   10 added
  domdebugger                 2 removed    0 changed    0 added
  emulation                   5 removed   15 changed   37 added
  extensions                  1 removed    0 changed    0 added
  fedcm                       4 removed    0 changed    0 added
  fetch                       1 removed    2 changed    7 added
  headlessexperimental        0 removed    1 changed    4 added
  indexeddb                   0 removed    2 changed    9 added
  input                       2 removed   12 changed   28 added
  layertree                   0 removed    1 changed    4 added
  log                         0 removed    4 changed   29 added
  media                       0 removed    1 changed    5 added
  memory                      1 removed    0 changed    0 added
  network                    27 removed   11 changed   77 added
  overlay                     4 removed    1 changed    3 added
  page                       10 removed   18 changed   44 added
  performance                 0 removed    2 changed    3 added
  preload                     6 removed    0 changed    0 added
  pwa                         1 removed    0 changed    0 added
  runtime                     0 removed    9 changed  142 added
  security                    4 removed    0 changed    0 added
  serviceworker               3 removed    0 changed    0 added
  smartcardemulation          5 removed    0 changed    0 added
  storage                     2 removed    0 changed    0 added
  systeminfo                  2 removed    0 changed    0 added
  target                      1 removed    0 changed    0 added
  tracing                     4 removed    3 changed    8 added
  webaudio                    5 removed    0 changed    0 added
  webauthn                    3 removed    0 changed    0 added
  webmcp                      1 removed    0 changed    0 added

## v0.157.0 - 2026-10-03

- Chromium: 157.0.8084.3
- V8: 15.7.23
- API changes: first tagged release
