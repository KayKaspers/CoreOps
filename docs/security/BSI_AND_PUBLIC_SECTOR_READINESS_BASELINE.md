# CoreOps – BSI and Public-Sector Readiness Baseline

> Document Status: Implemented, pending Nova review
> Baseline Status: Foundation baseline
> Implementation Status: Not implemented
> Validation Status: Not assessed
> Certification Status: None claimed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-004C` (docs-only / security-governance baseline)

## 1. Status

Foundation baseline for BSI-oriented and public-sector-adjacent deployment scenarios. This document defines readiness structure, claim boundaries, profiles, responsibility separation, and an evidence model. It implements **no** controls, selects **no** technology, and claims **no** certification. Most technical readiness domains are currently at least `design-required` / `implementation-required`.

## 2. Purpose

Provide a defensible, honest, public-neutral baseline that lets CoreOps be positioned for IT-Grundschutz-oriented and public-sector security processes **without** overclaiming compliance, certification, or approval. The baseline is the structural foundation on which later, version-accurate control-mapping work packages build.

## 3. Scope

- Allowed BSI positioning and its claim boundaries.
- Internal Standard / Hardened / Government readiness profiles.
- Eighteen public-sector readiness domains (PSR-01 … PSR-18).
- Product / operator / shared responsibility separation.
- Evidence model with a strict capability-vs-satisfaction distinction.
- Logging and detection, offline/air-gap, sovereignty/cloud boundaries.
- Registration of an official BSI reference set (as references, not as invented clauses).

## 4. Non-Goals

- No technical implementation.
- No capability-by-capability control mapping (belongs to later work).
- No product certification or certification preparation.
- No Schutzbedarfsfeststellung for a concrete authority.
- No VS-NfD assessment.
- No C5 / C3A audit.
- No legal assessment of a concrete operator.
- No selection of providers, products, or architecture.
- No modification of the Capability Matrix (belongs to `CO-WP-004E`).

## 5. Claims Boundary

### 5.1 Allowed statements

```text
CoreOps is designed with BSI-oriented security and governance principles in mind.

CoreOps provides a foundation for deployments that may need to align with
IT-Grundschutz-oriented or public-sector security processes.

The Government profile is an internal CoreOps readiness profile.

CoreOps can provide evidence-supporting capabilities where those capabilities
are implemented, configured and validated.
```

### 5.2 Prohibited statements

```text
CoreOps is BSI-certified.
CoreOps is fully IT-Grundschutz compliant.
CoreOps is approved for the German federal administration.
CoreOps is automatically suitable for classified information.
CoreOps is VS-NfD approved.
Using CoreOps makes an organisation compliant.
The Government profile is a certification level.
```

### 5.3 Deployment boundary

```text
Product readiness does not equal deployment compliance.

Compliance and approval depend on:
- deployment architecture,
- operator processes,
- configuration,
- connected systems,
- identities and roles,
- operational evidence,
- legal applicability,
- risk treatment,
- and external assessment where required.
```

## 6. Official Reference Set

Registered by Nova from official BSI sources, reference date 2026-07-13. Detailed provenance and version-verification status live in [BSI_REFERENCE_AND_CLAIMS_REGISTER.md](BSI_REFERENCE_AND_CLAIMS_REGISTER.md).

Primary BSI reference set:

```text
BSI-Standard 200-1 — Managementsysteme für Informationssicherheit
BSI-Standard 200-2 — IT-Grundschutz-Methodik
BSI-Standard 200-3 — Risikomanagement
BSI-Standard 200-4 — Business Continuity Management
BSI IT-Grundschutz-Kompendium — edition to be confirmed version-accurately before any later detail mapping
Aktuelle Mindeststandards des BSI für die Bundesverwaltung
Mindeststandard des BSI zur Protokollierung und Detektion von Cyberangriffen, Version 2.1 (2024-11-11)
```

Conditional cloud/sovereignty references (only relevant for cloud-involving deployments):

```text
C5:2026 — Cloud Computing Compliance Criteria Catalogue
C3A — Criteria enabling Cloud Computing Autonomy
```

## 7. Source and Version Boundary

The full official documents were **not** handed to this work package as a source pack. Therefore:

- No individual requirement numbers are invented.
- No verbatim BSI requirements are asserted.
- No full control coverage is claimed.
- No certifiability is derived.
- No clause-to-capability matrix is produced here.
- The sources are registered as an official reference set only.
- Version-accurate source ingestion is marked as a later step (see register §13).

If local copies of official documents exist, they must be path- and version-documented, identity/completeness-checked, used read-only, and must not create additional allowed files.

## 8. CoreOps Profile Model

The profile names are **internal CoreOps product profiles**, not BSI classifications. Full definitions live in [PUBLIC_SECTOR_READINESS_PROFILE.md](../governance/PUBLIC_SECTOR_READINESS_PROFILE.md). Summary:

- **Standard Profile** — secure defaults for general self-hosted and enterprise environments; basic audit and role capability; no special authority claim.
- **Hardened Profile** — extends Standard with restrictive defaults, stronger authentication and role limits, stricter logging/audit, hardening and integrity checks, higher evidence expectations, more controlled change, stronger offline/recovery capability.
- **Government Profile** — extends Hardened as an internal readiness profile with public-sector governance, traceable responsibility assignment, broader audit/evidence, product/operator separation, controlled dependencies, documented offline/air-gap capability, sovereignty and data-control expectations, stronger change/approval/separation-of-duties limits, exportable evidence, and explicit claim boundaries.

```text
Government Profile
≠ BSI certification
≠ government approval
≠ VS-NfD approval
≠ deployment compliance
```

## 9. Readiness Domains

Status vocabulary (no implementation claims without existing evidence):

```text
baseline-defined · reference-verification-pending · design-required ·
implementation-required · evidence-model-required · validation-required · not-applicable
```

Each domain below records: Purpose · Relevant BSI reference family · Product / Operator / Shared responsibility · Standard/Hardened/Government relevance · Expected evidence categories · Current CoreOps status · Known gaps · Follow-up target.

### PSR-01 Governance and ISMS Support
- **Purpose:** Support an operator's ISMS (roles, policies, documentation, review cadence).
- **BSI reference family:** BSI 200-1, 200-2.
- **Product responsibility:** Governance-supporting structures, documentation hooks, role model.
- **Operator responsibility:** Actual ISMS, scope definition, management endorsement.
- **Shared responsibility:** Policy configuration, review cadence.
- **Profiles:** Standard (basic) · Hardened (documented) · Government (comprehensive).
- **Expected evidence:** documentation and training evidence; risk and exception evidence.
- **Current status:** `design-required`.
- **Known gaps:** No ISMS-support structures implemented.
- **Follow-up:** CO-WP-007 / later governance baseline.

### PSR-02 Asset and Configuration Inventory
- **Purpose:** Maintain an authoritative inventory of assets and configuration with provenance.
- **BSI reference family:** IT-Grundschutz Kompendium (asset/configuration families, version-pending).
- **Product responsibility:** Source-of-truth, field provenance, inventory model.
- **Operator responsibility:** Completeness, classification, ownership assignment.
- **Shared responsibility:** Data quality, reconciliation cadence.
- **Profiles:** Standard · Hardened · Government (all relevant, increasing rigor).
- **Expected evidence:** configuration evidence; deployment-topology evidence.
- **Current status:** `design-required`.
- **Known gaps:** Source-of-truth capabilities not implemented (see Capability Matrix, read-only).
- **Follow-up:** CO-WP-011 / CO-WP-012.

### PSR-03 Identity, Authentication and Authorization
- **Purpose:** Human and machine identity, authentication strength, role-based authorization.
- **BSI reference family:** IT-Grundschutz identity/access families (version-pending); Mindeststandards.
- **Product responsibility:** Identity model, RBAC, authentication options.
- **Operator responsibility:** Role assignment, credential policy, joiner/mover/leaver.
- **Shared responsibility:** Authentication configuration, session policy.
- **Profiles:** Standard (RBAC) · Hardened (stronger auth, tighter roles) · Government (SoD, break-glass control).
- **Expected evidence:** identity and authorization evidence.
- **Current status:** `design-required`.
- **Known gaps:** Identity/RBAC not implemented (CO-WP-009).
- **Follow-up:** CO-WP-009.

### PSR-04 Privileged Access and Separation of Duties
- **Purpose:** Constrain privileged access; enforce separation of duties and break-glass control.
- **BSI reference family:** IT-Grundschutz privileged-access families (version-pending).
- **Product responsibility:** Privilege model, approval hooks, break-glass with audit.
- **Operator responsibility:** SoD policy, approver assignment, review of privileged use.
- **Shared responsibility:** Approval workflows, privileged-session logging.
- **Profiles:** Standard (basic) · Hardened (explicit SoD) · Government (strict SoD + break-glass evidence).
- **Expected evidence:** identity and authorization evidence; change and approval evidence; audit-log evidence.
- **Current status:** `design-required`.
- **Known gaps:** Privileged-access and break-glass model not implemented (relates to RISK-04).
- **Follow-up:** CO-WP-009 / CO-WP-013.

### PSR-05 Secure Configuration and Hardening
- **Purpose:** Secure defaults and a hardening baseline per profile.
- **BSI reference family:** IT-Grundschutz hardening/system families (version-pending).
- **Product responsibility:** Secure defaults, hardening options, integrity checks.
- **Operator responsibility:** Applying hardening, environment-specific configuration.
- **Shared responsibility:** Configuration review, drift detection.
- **Profiles:** Standard (secure defaults) · Hardened (restrictive) · Government (restrictive + verified).
- **Expected evidence:** configuration evidence.
- **Current status:** `design-required`.
- **Known gaps:** Hardening baseline not implemented.
- **Follow-up:** CO-WP-007 / CO-WP-004E alignment.

### PSR-06 Vulnerability, Patch and Update Governance
- **Purpose:** Track vulnerabilities; govern patching and updates, including offline update paths.
- **BSI reference family:** IT-Grundschutz patch/update families (version-pending).
- **Product responsibility:** Update mechanism, integrity/provenance verification, version transparency.
- **Operator responsibility:** Patch windows, risk acceptance, testing.
- **Shared responsibility:** Update scheduling, vulnerability triage.
- **Profiles:** Standard · Hardened · Government (documented, controlled).
- **Expected evidence:** patch and vulnerability evidence; dependency and provenance evidence.
- **Current status:** `design-required`.
- **Known gaps:** Update/patch governance not implemented.
- **Follow-up:** CO-WP-022.

### PSR-07 Logging, Audit and Security Event Evidence
- **Purpose:** Produce trustworthy audit logs and security-event evidence.
- **BSI reference family:** Mindeststandard Protokollierung und Detektion v2.1; IT-Grundschutz logging families.
- **Product responsibility:** Audit-log generation, integrity, export, time reference.
- **Operator responsibility:** Retention policy, access control, log review.
- **Shared responsibility:** Logging configuration, event-source coverage.
- **Profiles:** Standard (basic audit) · Hardened (stricter) · Government (comprehensive, tamper-evident).
- **Expected evidence:** audit-log evidence; security-event evidence.
- **Current status:** `evidence-model-required`.
- **Known gaps:** Audit/event model not implemented (relates to RISK-08, RISK-25).
- **Follow-up:** CO-WP-018 / CO-WP-025; see §12 logging and detection.

### PSR-08 Detection and Incident-Response Support
- **Purpose:** Support detection and incident-response processes (interfaces, not a full SIEM).
- **BSI reference family:** Mindeststandard Protokollierung und Detektion v2.1.
- **Product responsibility:** Detection/escalation interfaces, event export/forwarding capability.
- **Operator responsibility:** Detection rules, IR process, SIEM operation.
- **Shared responsibility:** Alert routing, escalation configuration.
- **Profiles:** Standard (interfaces) · Hardened (richer) · Government (documented IR support).
- **Expected evidence:** security-event evidence; documentation and training evidence.
- **Current status:** `design-required`.
- **Known gaps:** Detection/IR interfaces not implemented; full SIEM is a Non-Goal (DEC-N-02).
- **Follow-up:** CO-WP-019 / CO-WP-026.

### PSR-09 Backup, Restore and Recovery Evidence
- **Purpose:** Support backup/restore with verifiable recovery evidence.
- **BSI reference family:** BSI 200-4; IT-Grundschutz backup families.
- **Product responsibility:** Backup/restore hooks, integrity, restore-test evidence.
- **Operator responsibility:** Backup targets, RPO/RTO, restore testing.
- **Shared responsibility:** Restore-test cadence, verification.
- **Profiles:** Standard · Hardened · Government (documented, tested).
- **Expected evidence:** backup and restore evidence; continuity-test evidence.
- **Current status:** `design-required`.
- **Known gaps:** Backup/restore not implemented.
- **Follow-up:** CO-WP-026.

### PSR-10 Business Continuity and Offline Operations
- **Purpose:** Support continuity and degraded/offline operations.
- **BSI reference family:** BSI 200-4.
- **Product responsibility:** Degraded modes, local core function, recovery mode.
- **Operator responsibility:** BCM plan, continuity objectives, exercises.
- **Shared responsibility:** Continuity testing, offline drills.
- **Profiles:** Standard · Hardened (offline-capable) · Government (documented offline/air-gap).
- **Expected evidence:** continuity-test evidence; deployment-topology evidence.
- **Current status:** `design-required`.
- **Known gaps:** Continuity/degraded-mode model not implemented (RISK-11); see §13.
- **Follow-up:** CO-WP-023 / CO-WP-026.

### PSR-11 Network Segmentation and Trust Boundaries
- **Purpose:** Support network segmentation and clear trust boundaries.
- **BSI reference family:** IT-Grundschutz network families (version-pending).
- **Product responsibility:** Documented trust boundaries, segmentation-friendly architecture.
- **Operator responsibility:** Actual network segmentation, firewalling.
- **Shared responsibility:** Boundary configuration, connectivity review.
- **Profiles:** Standard · Hardened (separated networks) · Government (strict boundaries).
- **Expected evidence:** deployment-topology evidence; configuration evidence.
- **Current status:** `design-required`.
- **Known gaps:** Trust-boundary model not yet defined (CO-WP-007).
- **Follow-up:** CO-WP-006 / CO-WP-007.

### PSR-12 Cryptography, Secrets and Key Governance
- **Purpose:** Govern cryptography, secrets storage, and key custody.
- **BSI reference family:** IT-Grundschutz crypto/key families (version-pending).
- **Product responsibility:** Secret storage model, key custody hooks, crypto configuration options.
- **Operator responsibility:** Key management policy, rotation, custody roles.
- **Shared responsibility:** Rotation cadence, crypto configuration.
- **Profiles:** Standard · Hardened · Government (stringent custody + evidence).
- **Expected evidence:** configuration evidence; identity and authorization evidence.
- **Current status:** `design-required`.
- **Known gaps:** Secrets/key custody not implemented (RISK-09, critical).
- **Follow-up:** CO-WP-024.

### PSR-13 Deployment, Change and Automation Control
- **Purpose:** Control deployment, change, and automation with approval and authorization.
- **BSI reference family:** IT-Grundschutz change/deployment families (version-pending).
- **Product responsibility:** Change/approval hooks, authorization model, automation guardrails.
- **Operator responsibility:** Change process, approver roles, maintenance windows.
- **Shared responsibility:** Approval workflow, automation scope.
- **Profiles:** Standard · Hardened · Government (strict change + SoD).
- **Expected evidence:** change and approval evidence; deployment-topology evidence.
- **Current status:** `design-required`.
- **Known gaps:** Policy/approval/execution-authorization not implemented (RISK-04).
- **Follow-up:** CO-WP-013 / CO-WP-021.

### PSR-14 Dependency, Supplier and Software Supply Chain
- **Purpose:** Govern dependencies and software supply chain (provenance, SBOM, revocation).
- **BSI reference family:** IT-Grundschutz supply-chain families (version-pending).
- **Product responsibility:** SBOM support, provenance/integrity verification, revocation handling.
- **Operator responsibility:** Dependency acceptance, supplier assessment, update policy.
- **Shared responsibility:** Dependency updates, provenance review.
- **Profiles:** Standard · Hardened · Government (controlled dependencies).
- **Expected evidence:** dependency and provenance evidence.
- **Current status:** `design-required`.
- **Known gaps:** Supply-chain controls not implemented (RISK-05, RISK-30); governed in part by [SOVEREIGNTY_AND_DEPENDENCY_POLICY.md](../architecture/SOVEREIGNTY_AND_DEPENDENCY_POLICY.md).
- **Follow-up:** CO-WP-022.

### PSR-15 Data Protection, Retention and Deletion Support
- **Purpose:** Support data classification, retention, redaction, and deletion.
- **BSI reference family:** IT-Grundschutz data-protection families (version-pending); privacy law is operator-dependent.
- **Product responsibility:** Data classification, retention/redaction hooks, deletion support.
- **Operator responsibility:** Retention periods, legal applicability, deletion decisions.
- **Shared responsibility:** Retention configuration, redaction rules.
- **Profiles:** Standard · Hardened · Government (documented data control).
- **Expected evidence:** documentation and training evidence; audit-log evidence.
- **Current status:** `design-required`.
- **Known gaps:** Data classification/retention not implemented (RISK-08, RISK-25).
- **Follow-up:** CO-WP-025.

### PSR-16 Documentation, Evidence Export and Auditability
- **Purpose:** Provide exportable, auditable documentation and evidence.
- **BSI reference family:** BSI 200-1/200-2 documentation expectations.
- **Product responsibility:** Evidence export, documentation structure, integrity/time reference.
- **Operator responsibility:** Evidence review, audit preparation, assessor interaction.
- **Shared responsibility:** Export cadence, scope definition.
- **Profiles:** Standard · Hardened · Government (comprehensive, exportable).
- **Expected evidence:** documentation and training evidence; all evidence categories exportable.
- **Current status:** `evidence-model-required`.
- **Known gaps:** Evidence-export capability not implemented; see §11 evidence model.
- **Follow-up:** CO-WP-018 / later evidence-export WP.

### PSR-17 Sovereignty, Data Control and Exit Strategy
- **Purpose:** Support data sovereignty, operator data control, and exit strategy.
- **BSI reference family:** Conditional — C3A for relevant cloud scenarios; otherwise governance.
- **Product responsibility:** No mandatory external control plane; export/exit capability.
- **Operator responsibility:** Data location, jurisdiction, operator-access decisions.
- **Shared responsibility:** Exit planning, dependency review.
- **Profiles:** Standard · Hardened · Government (documented sovereignty + exit).
- **Expected evidence:** deployment-topology evidence; documentation and training evidence.
- **Current status:** `design-required`.
- **Known gaps:** Exit/sovereignty model not fully defined; governed in part by sovereignty policy.
- **Follow-up:** CO-WP-023; see §14.

### PSR-18 Cloud and External-Service Boundary
- **Purpose:** Keep cloud/external services optional and clearly bounded.
- **BSI reference family:** Conditional — C5:2026 for relevant cloud scenarios.
- **Product responsibility:** Core functions without mandatory cloud; clear external-service boundary.
- **Operator responsibility:** Cloud-service selection, provider assessment, applicability decision.
- **Shared responsibility:** Integration configuration, boundary review.
- **Profiles:** Standard · Hardened · Government (documented, minimized external services).
- **Expected evidence:** deployment-topology evidence; dependency and provenance evidence.
- **Current status:** `design-required`.
- **Known gaps:** External-service boundary not implemented; conditional references not yet applicability-assessed.
- **Follow-up:** CO-WP-023; see §14.

## 10. Responsibility Model

No control is satisfied merely because CoreOps offers a technical function.

- **CoreOps Product Responsibility** — secure product defaults; role and permission model; audit functions; configuration options; evidence export; integrity and update functions; documentation.
- **Deployment and Operator Responsibility** — Schutzbedarfsfeststellung; concrete network segmentation; user/role assignment; retention periods; backup targets; operational processes; incident response; legal applicability; local approvals.
- **Shared Responsibility** — logging configuration; patch/maintenance windows; integrations; backup/restore tests; dependency updates; risk acceptance; evidence collection.

## 11. Evidence Model

Evidence categories:

```text
configuration evidence · identity and authorization evidence · change and approval evidence ·
audit-log evidence · security-event evidence · patch and vulnerability evidence ·
backup and restore evidence · continuity-test evidence · dependency and provenance evidence ·
deployment-topology evidence · risk and exception evidence · documentation and training evidence
```

Minimum evidence properties: traceable origin; time reference; scope; responsible party or role; integrity; retention/export capability; no secrets in unsuitable reports; no automatic compliance statement.

Three states that must **never** be conflated:

```text
Evidence capability:   CoreOps can generate or export evidence data.
Evidence available:    A concrete deployment actually holds the evidence.
Requirement satisfied: A responsible body has assessed and accepted the evidence.
```

## 12. Logging and Detection

Registered as a particularly relevant public-sector domain (see PSR-07/PSR-08 and Mindeststandard Protokollierung und Detektion v2.1). The baseline requires at least: a defined logging policy; an event-source inventory; a time and time-synchronisation concept; role and access protection; tamper protection; retention and deletion concept; export and forwarding capability; detection and escalation interfaces; data-protection and secret protection; documented responsibilities.

Not yet fixed (deployment- and requirement-dependent): concrete retention periods; concrete SIEM products; concrete detection rules; concrete data volumes; concrete authority mandates.

## 13. Offline and Air-Gap Readiness

CoreOps shall be able to support public-sector and government scenarios without a permanent internet connection. The baseline documents at least: local core function without mandatory cloud; local identity and role operation; local monitoring and audit; offline documentation; controlled import of updates and artifacts; provenance and integrity verification; local package/artifact sources as a later architecture question; evidence export without permanent external connectivity; backup and restore without cloud dependency; defined handling of unreachable external services; exit and restart strategy.

Not claimed: fully implemented; suitable for every air-gap tier; approved for classified networks.

## 14. Sovereignty and Cloud Boundary

- The CoreOps core must not require a mandatory external control plane.
- Cloud services remain optional integrations.
- Data location, jurisdiction, operator access, and exit must be assessed.
- C5:2026 is a conditional security reference for relevant cloud scenarios.
- C3A is a conditional sovereignty reference for relevant cloud scenarios.
- On-premises or offline CoreOps is not automatically C5- or C3A-relevant.
- Using a C5-assessed service does not automatically make the CoreOps deployment compliant.
- Sovereignty is multi-dimensional: at least technical, operational, legal, and exit aspects.

No provider or product is selected.

## 15. Risk and Exception Handling

New/updated risks are registered in [RISK_REGISTER.md](../../project-system/RISK_REGISTER.md) (RISK-40 … RISK-49): compliance/certification overclaim, outdated BSI reference versions, incomplete source ingestion, ambiguous product/operator responsibility, evidence-capability-vs-satisfaction confusion, government-profile-vs-approval confusion, offline/air-gap overclaim, cloud/sovereignty misclassification, public-sector logging/retention assumptions, unverified capability-to-control mapping. No such risk is closed while its follow-up work is pending. Exceptions to this baseline require a dedicated work package, Nova review, and Human-Maintainer approval.

## 16. Relationship to Capability Matrix

This baseline does **not** modify the Capability Matrix. The security and governance alignment of the Capability Matrix (mapping PSR domains and profiles onto concrete capabilities and their status dimensions) is the subject of `CO-WP-004E`. Until then, the Capability Matrix ([FOUNDATION_CAPABILITY_MATRIX.md](../architecture/FOUNDATION_CAPABILITY_MATRIX.md)) remains authoritative for capability status, read-only here.

## 17. Relationship to Future Work Packages

- `CO-WP-004D` — ITIL and PRINCE2 applicability and tailoring decision (governance vocabulary for change/incident/continual-improvement that intersects several PSR domains).
- `CO-WP-004E` — Capability Matrix security and governance alignment (maps this baseline onto capabilities).
- Version-accurate BSI source ingestion and clause-level control mapping — later dedicated work packages.
- Security baseline / threat model — `CO-WP-007` and dedicated security-baseline WPs.

## 18. Compatibility

This baseline is additive to the accepted product direction ([COREOPS_CONCEPT_V3_1_AMENDMENT.md](../architecture/COREOPS_CONCEPT_V3_1_AMENDMENT.md)) and the existing positioning ([BSI_ALIGNMENT_POSITIONING.md](BSI_ALIGNMENT_POSITIONING.md)). It changes no existing accepted decision, selects no technology, and creates no ADR. It does not claim a normative release; normative release assignment remains pending.

## 19. Open Questions

- Which IT-Grundschutz-Kompendium edition and which requirement families apply per profile (version-accurate, later ingestion)?
- Which PSR domains are mandatory per profile vs. optional?
- How are C5/C3A applicability decisions triggered per deployment?
- How is evidence export structured technically (later capability WP)?
- Which retention and logging parameters are baseline-recommended vs. purely operator-defined?

## 20. Next Decision

Nova review of this baseline, followed by a Human-Maintainer commit. Only afterwards may `CO-WP-004D` begin. Version-accurate BSI source ingestion and the Capability Matrix alignment (`CO-WP-004E`) remain separate, later work.
