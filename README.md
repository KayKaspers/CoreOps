# CoreOps

**One Dashboard. Controlled Operations.**

CoreOps is a planned universal, vendor-neutral, self-hosted and offline-capable
operations control plane. **This repository does not yet contain a usable CoreOps
product.** It currently holds the project's architecture, security, governance and
contract foundation, plus the state artifacts that track how that foundation is
allowed to become software.

This README is a public entry point and a low-authority summary. It is **not** a
normative source: it does not override any governance, architecture or release
document, and it grants no authority of its own. Where a detail can drift, this
file links to the authoritative source instead of duplicating it.

## Status

- Foundation 0.1 is **closed**.
- The `Observe` phase has been **entered with an explicit boundary**.
- There is **no functional CoreOps product release**. No usable product is claimed.
- Productive application code is **not authorized** and **not present**.
- Implementation is **not authorized**.
- `Gate A` (productive code / implementation) is **not passed**.
- `Gate B` (target access / execution) is **not passed**.
- The publication artifacts `README.md` and `LICENSE` now exist. Their creation was a
  separately authorized repository action and authorizes nothing further.

For the authoritative, maintained gate and transition state, see
[Productive-Code Transition Prerequisites](docs/governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md).

## What CoreOps Is

The accepted product vision describes CoreOps as a modular operations control plane
that brings source of truth, monitoring, topology, desired state, policy, trusted
automation and auditable change together in one interface — self-hosted, without a
cloud requirement, and usable in restricted, offline and air-gapped environments.

Its intended users are homelabs, self-hosters, associations, small and medium-sized
organizations, educational institutions, developers and IT departments, including
distributed sites and offline environments.

The product principles it is being designed against include: documentation first,
architecture first, security first, privacy by design, fail closed, read-only before
write, preview before execute, verification after every change, auditability, offline
first, and *AI creates, humans approve*.

All of the above describes **intent and accepted vision**, not delivered behaviour.
Nothing in this section is implemented.

## What This Repository Contains

Currently, this repository consists primarily of documentation and project state:

- **Architecture documentation** — the product concept, module and plane models, state
  and drift models, telemetry, topology, deployment and artifact models.
- **Governance documentation** — scope lock, release taxonomy, repository governance,
  language standard, data-handling and retention policy, transition prerequisites.
- **Security documentation** — threat model, trust boundaries, policy, identity,
  secrets, audit and disclosure models. This is documentation, not an assessment,
  audit result or certification.
- **Testing and contract documents** — observation contracts and test envelopes that
  define intended behaviour. No test is implemented and none has been executed.
- **Project-system state and registers** — work package queue, decision index, risk
  register, project profile and phase state.
- **Project-brain and context artifacts** — the maintained working memory and context
  packs used to carry project state between sessions.
- **Publication artifacts** — this `README.md` and the root [`LICENSE`](LICENSE).

There is no productive application code in this repository. There is no Go source
tree and no Go module metadata.

## What This Repository Does Not Yet Provide

This repository provides **no usable CoreOps product**. Specifically, there is:

- no installation or setup procedure, and nothing to install;
- no build, package, container image or distributable artifact;
- no deployed, running or operational service;
- no supported runtime, platform or integration;
- no production readiness, and no claim of it;
- no service-level agreement, uptime, maintenance or compatibility commitment;
- no release schedule and no completion date.

## Current Development Phase

Foundation 0.1 is closed. The `Observe` phase has been entered with an explicit
boundary. The first selected Observe value slice is
**Local Linux Host Identity & Basic System Observation**.

Selection is not delivery. The following distinctions are binding:

```text
Observe entered        != target access authorized
slice selected         != productive source authorized
slice selected         != slice implemented
contract documented    != behaviour implemented
```

The slice is defined as a read-only observation contract. It is **not** implemented,
no target access is authorized, and no real observation has been performed.

## Governance and Architecture Sources

These documents are authoritative; this README is not.

- [Project Brief](docs/architecture/PROJECT_BRIEF.md) — project identity, scope, goals and non-goals.
- [Foundation Scope Lock](docs/governance/FOUNDATION_SCOPE_LOCK.md) — what the Foundation phase was permitted to produce.
- [Release and Version Taxonomy](docs/governance/RELEASE_TAXONOMY.md) — milestone, version and tag semantics.
- [Productive-Code Transition Prerequisites](docs/governance/PRODUCTIVE_CODE_TRANSITION_PREREQUISITES.md) — the authoritative current gate state.
- [Work Package Queue](project-system/WORK_PACKAGE_QUEUE.md) — the active, binding work package queue.
- [Repository Governance Standard](docs/governance/REPOSITORY_GOVERNANCE_STANDARD.md) — repository rules and the source-of-truth hierarchy.
- [Observe — Local Linux Host Observation Contract](docs/architecture/OBSERVE_LOCAL_LINUX_HOST_OBSERVATION_CONTRACT.md) — the first Observe slice contract.

## License

CoreOps is licensed under the Apache License 2.0. The full, unmodified license text
is in the root [`LICENSE`](LICENSE) file.

The license governs copyright permissions only. It is not a statement about the
project's maturity, fitness or support:

```text
licensed                != warranted
open source             != supported service
source visible          != supported
repository contains docs != product released
```

## Contributions

**A public contribution program is not active.**

This repository being visible does not mean contributions are being solicited,
reviewed, accepted or merged. There is no `CONTRIBUTING.md`, no contributor licence
agreement and no developer certificate of origin, and no contribution workflow is
defined. Please do not expect review, response, merge or contributor support for
unsolicited changes at this time.

```text
repository visible      != contribution program active
pull request possible   != contribution accepted
contribution received   != contribution reviewed or merged
```

## Project Name and Branding

The Apache License 2.0 applies to copyrighted material. It does not grant rights to
use the CoreOps project name, wordmark, logo or branding in a way that implies
endorsement of, or affiliation with, the project.

A factual reference to CoreOps does not by itself imply endorsement or affiliation.
No registered trademark is claimed here, and a separate trademark policy is deferred
and does not exist.

## Support, Security and Compliance Claims

None are made. This project makes **no** claim of production readiness, safety,
security, compliance or any other certification; **no** support commitment, service
level or uptime guarantee; **no** maintenance or compatibility guarantee; and **no**
installation readiness or release schedule.

The repository contains documented security models, threat analysis and policy
material. That material is design documentation about intended behaviour. It is not
an audit, an assessment, a certification, or evidence about any running system —
because no system is running.

There is no `SECURITY.md` and no published vulnerability reporting process at this
time.

## Language

This README is written in English. English is the canonical language for
machine-facing artifacts across the project; much of the current governance and
architecture documentation is written in German. CoreOps is intended to be a German
and English project for product-facing content.

**No claim of full German/English translation parity is made.** Parity is claimed only
where it has actually been reviewed. See the
[Language Standard](docs/governance/COREOPS_LANGUAGE_STANDARD.md) for the binding rules.
