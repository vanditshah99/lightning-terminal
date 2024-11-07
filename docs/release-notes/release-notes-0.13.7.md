# Release Notes

- [Lightning Terminal](#lightning-terminal)
  - [Bug Fixes](#bug-fixes)
  - [Functional Changes/Additions](#functional-changesadditions)
  - [Technical and Architectural Updates](#technical-and-architectural-updates)
- [Integrated Binary Updates](#integrated-binary-updates)
  - [LND](#lnd)
  - [Loop](#loop)
  - [Pool](#pool)
  - [Faraday](#faraday)
  - [Taproot Assets](#taproot-assets)
- [Contributors](#contributors-alphabetical-order)

## Lightning Terminal

### Bug Fixes

### Functional Changes/Additions

* Add [custom channel
  functionality](https://github.com/lightninglabs/lightning-terminal/pull/848)
  to `litd`. Custom channels with Taproot Assets can now be created when `litd`
  runs in integrated `lnd` mode (`lnd-mode=integrated`) with the Taproot Assets
  daemon also running in integrated mode (`taproot-assets-mode=integrated`).

* [Add itest](https://github.com/lightninglabs/lightning-terminal/pull/892) for
  the MinRelayFee check added in Taproot Assets. The test ensures that
  transactions with fees below the minimum relay fee are rejected.

### Technical and Architectural Updates

## Integrated Binary Updates

### LND

### Loop

### Pool

### Faraday

### Taproot Assets



# Contributors (Alphabetical Order)

* George Tsagkarelis
* Gijs van Dam
* Jamal James
* Jonathan Harvey-Buschel
* Oliver Gugger
