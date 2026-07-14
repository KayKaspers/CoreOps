# CoreOps – Source of Truth and State Authority Model

> Document Status: Implemented, pending Nova review
> Model Status: Foundation source-of-truth and state-authority model
> Implementation Status: Not implemented
> Storage Technology: Not selected
> Reconciliation Technology: Not selected
> Validation Status: Not performed
> Normative Release: Not yet assigned
> Normative Framework: NDF v1.0.0 (Tag `v1.0.0`, Commit `9dcadc1`) — `main` informativ, nicht normativ
> Erzeugt durch `CO-WP-011` (docs-only / data-governance and architecture foundation)

## 1. Status

Foundation, technology-independent model for authoritative data sources, state semantics (desired/observed/effective/derived/cached), and conflict handling. It selects no database, storage, merge, messaging, or synchronisation technology and implements nothing. Companions: [FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md](FIELD_PROVENANCE_AND_DATA_LINEAGE_STANDARD.md) and [OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md).

## 2. Purpose

Make authoritative data ownership, state distinctions, and conflict rules explicit so later data-architecture work can design storage and reconciliation against a shared, honest model — without conflating observed with desired state, derived with authoritative, or a cache with the source of truth.

## 3. Scope

Concepts; source-of-truth vs system-of-record; authority classes; source classes; authoritative modules; desired/observed/effective/derived/cached state; field authority; conflict model; manual overrides; time/sequence boundary; offline; security invariants; threat references.

## 4. Non-Goals

- No database schema, API field definitions, event schema, message-broker, synchronisation protocol, CRDT, merge engine, storage, time-service, or signature selection.
- No runtime code; no full field list of all capabilities; no ADR.
- No identity/module/capability/threat-file change.

## 5. Concepts

`source of truth` · `system of record` · `authoritative source` · `authoritative module` · `field authority` · `source record` · `source reference` · `provenance` · `lineage` · `desired state` · `observed state` · `effective state` · `reported state` · `derived state` · `cached state` · `imported state` · `manual override` · `conflict` · `reconciliation` · `freshness` · `staleness` · `confidence` · `trust status` · `validation status`.

Ground rules:

```text
source of truth ≠ cached copy · system of record ≠ every consuming system ·
observed state ≠ desired state · effective state ≠ automatically compliant state ·
derived state ≠ authoritative source · manual override ≠ silent replacement of provenance ·
latest timestamp ≠ automatically most trustworthy value · successful import ≠ validated provenance ·
field presence ≠ field correctness
```

## 6. Source of Truth versus System of Record

A **source of truth** is the authoritative origin for a field/concept; a **system of record** is a system that holds records but is not automatically authoritative for every field it stores. CoreOps may be the system of record for many fields while another system remains the source of truth for others. A cache/summary/dashboard is never the source of truth.

## 7. Authority Classes

`authoritative` · `operator-declared` · `managed-system-reported` · `externally-reported` · `imported` · `derived` · `advisory` · `cached` · `unknown` · `conflicted`. The authority class describes the relationship to field authority; it is **not** a statement about technical validation, compliance, freshness, correctness, or trustworthiness of all values from the same source. Each authoritative field concept has exactly one authority rule; multiple sources may report a field, but not every source may overwrite the effective field; `unknown` is not positive/trusted; `conflicted` is not silently resolved.

## 8. Source Classes

Producer/source classes: human operator · CoreOps policy/inventory/observation/execution/identity/evidence modules · managed resource · agent/relay · adapter/integration · external provider · offline import package · migration source · derived computation · manual correction. Each has typical authority, permitted field classes, trust assumptions, required owner, provenance/freshness/offline considerations, revocation/invalidation path, and threat references (detailed in the provenance standard). No vendor/protocol dependency selected.

## 9. Authoritative Modules

Using the module boundaries from `CO-WP-008`:

```text
MOD-STA-001 may coordinate state references without silently becoming the authoritative
  owner of every domain field.
MOD-OBS-001 owns observations, not desired policy.
MOD-INV-001 owns registered-resource identity and management status, not every observed
  device property.
MOD-TOP-001 owns derived topology relationships, not authoritative inventory.
MOD-POL-001 owns effective policy/approval; MOD-IAM-001 owns identity/roles.
MOD-EVD-001 records audit/evidence references but gains no control authority.
MOD-OFF-001 owns intake/reconciliation metadata, not the imported source's original authority
  automatically.
```

Per concept: authoritative module · permitted writers · read consumers · derived consumers · conflict owner · override authority · audit requirement · offline behavior (see provenance standard).

## 10. Desired State

The intended/target state established by authorised policy, configuration, or workflow decision (owned via MOD-POL-001/MOD-WFL-001/MOD-STA-001). Observed state must not silently overwrite desired state.

## 11. Observed State

State reported by a managed resource, agent, adapter, or external system (owned by MOD-OBS-001). Observed state is evidence of observation, not ground truth; `stale ≠ current`; `unknown ≠ healthy`.

## 12. Effective State

The state CoreOps currently uses after applying documented authority, conflict, and override rules. Effective state does **not** automatically mean actual physical state, compliance, success, validation, or current freshness.

## 13. Derived and Cached State

**Derived state** is computed/derived from other data; it is not authoritative by default and must not hide the authority/uncertainty of its inputs. **Cached state** is a temporary copy with its own freshness/invalidation boundary; `cached value ≠ live value`. **Reported state** is a source's value before authority/conflict rules apply.

## 14. Field Authority

Field authority is explicit and stable, derived from a canonical field identity — not from UI labels, translated display names, table positions, or temporary adapter fields (see provenance standard §5). Per field concept, a defined authority policy states which source class may set the effective value.

## 15. Conflict Model

Conflict states: `no-conflict` · `potential-conflict` · `active-conflict` · `resolution-pending` · `resolved` · `accepted-exception` · `invalidated`. Ground rules: conflicts are not silently discarded; security-relevant fields do not use automatic last-write-wins; newer timestamps do not automatically win; a higher authority class wins only per documented field rules; resolution is auditable and preserves original provenance; `resolved` does not retroactively make the source trustworthy. Detailed detection/resolution in the [reconciliation policy](../security/OFFLINE_DATA_RECONCILIATION_AND_CONFLICT_POLICY.md).

## 16. Manual Overrides

Overrides are explicit, named-human attributable, reason-/scope-/field-bound, time-bound where appropriate, auditable, reversible/supersedable, and visible to consumers. They must preserve original provenance and must not rewrite audit history or apply automatic global effect. Detailed fields in the provenance standard §13. Threat/invariant: override ≠ silent provenance replacement.

## 17. Time and Sequence Boundary

Event time, observation time, receipt time, and effective time differ; system clocks can be wrong/manipulated; ordering must not follow from an untrusted timestamp alone; replay/duplication must be later-detectable; offline sources may deliver delayed but valid information; clock skew and sequence gaps surface as uncertainty. `received later ≠ happened later`; `timestamp present ≠ timestamp trusted`. No time/sequence/distributed-clock technology selected. Threat refs THR-026/027.

## 18. Offline Considerations

Offline/delayed changes, delayed observations, delayed revocation, and conflicting authority decisions are reconciled explicitly, scope-bound, conflict-aware, auditable, and fail-closed on unclear authority (see reconciliation policy). No claim of implemented reconciliation, automatic conflict-free merge, CRDT, or classified-network suitability.

## 19. Security Invariants

Design requirements (not implemented controls):

```text
Observed state must not silently overwrite desired state.
Derived state must not silently become authoritative.
Unknown state must not be interpreted as healthy.
Stale state must not be interpreted as current.
Latest timestamp must not automatically win.
Imported data must not inherit authority automatically.
Manual override must preserve original provenance.
Revoked sources must not remain authoritative.
Audit and evidence history must not be rewritten by normal reconciliation.
Offline reconciliation must not bypass provenance, integrity or approval.
Field conflicts must remain visible until explicitly resolved.
A summary or dashboard must not become an authoritative source.
```

## 20. Threat References

Relevant existing scenarios: THR-011 (malicious API response), THR-012 (telemetry manipulation), THR-013 (stale telemetry), THR-014 (inventory), THR-015 (topology), THR-016/017 (audit), THR-018 (evidence), THR-024 (offline import), THR-026 (replay), THR-027 (time source), THR-035 (tenant boundary). No parallel threat list; no invented IDs.

## 21. Technology Boundary

No storage, database, merge, CRDT, messaging, synchronisation, time-service, or cryptographic-provenance technology is selected; deferred to later ADR-governed work.

## 22. Compatibility

Additive; consistent with the module architecture (MOD-STA-001/MOD-OBS-001/MOD-INV-001/MOD-TOP-001/MOD-EVD-001/MOD-OFF-001), repository-governance source-of-truth rule, threat model, disclosure policy, and NDF rules. No technology; no ADR; no Capability Matrix change. Breaking-change potential: low.

## 23. Open Questions

- Which field classes need the strictest authority rules (later)?
- How is effective state computed technically (later)?
- Which storage/reconciliation model per deployment (later)?

## 24. Next Decision

Nova review, then Human-Maintainer commit. Storage, merge, and reconciliation technology remain separate later work packages (ADR-governed).
