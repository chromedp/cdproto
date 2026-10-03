# Changelog

Each release of `cdproto` is generated from the Chromium and V8 protocol
definitions listed below. The minor version is the Chromium major version. As
the protocol definitions deprecate and remove commands, events, types, and
fields, any release can contain incompatible changes to the generated API.

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
